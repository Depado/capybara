package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/spf13/viper"
)

type LogConf struct {
	Level  string `mapstructure:"level"`
	Type   string `mapstructure:"type"`
	Caller bool   `mapstructure:"caller"`
}

type ServerConf struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func (s ServerConf) ListenAddr() string {
	l := s.Host
	if s.Port > 0 {
		l += ":" + strconv.Itoa(s.Port)
	}
	return l
}

type DatabaseConf struct {
	Path                string        `mapstructure:"path"`
	MaxBucketsRecursion int           `mapstructure:"max_buckets_recursion"`
	DefaultLockTTL      time.Duration `mapstructure:"default_lock_ttl"`
}

type Conf struct {
	Log      LogConf      `mapstructure:"log"`
	Server   ServerConf   `mapstructure:"server"`
	Database DatabaseConf `mapstructure:"database"`
}

// NewLogger will return a new logger
func NewLogger(c *Conf) zerolog.Logger {
	// Level parsing
	warns := []string{}
	lvl, err := zerolog.ParseLevel(c.Log.Level)
	if err != nil {
		warns = append(warns, fmt.Sprintf("unrecognized log level '%s', fallback to 'info'", c.Log.Level))
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		zerolog.SetGlobalLevel(lvl)
	}

	// Type parsing
	switch c.Log.Type {
	case "console":
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	case "json":
		break
	default:
		warns = append(warns, fmt.Sprintf("unrecognized log type '%s', fallback to 'json'", c.Log.Type))
	}

	// Caller
	if c.Log.Caller {
		log.Logger = log.With().Caller().Logger()
	}

	// Log messages with the newly created logger
	for _, w := range warns {
		log.Warn().Msg(w)
	}

	return log.Logger
}

// NewConf will parse and return the configuration
func NewConf() (*Conf, error) {
	// Environment variables
	viper.AutomaticEnv()
	viper.SetEnvPrefix("vuemonit")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Configuration file
	if viper.GetString("conf") != "" {
		viper.SetConfigFile(viper.GetString("conf"))
	} else {
		viper.SetConfigName("conf")
		viper.AddConfigPath(".")
		viper.AddConfigPath("/config/")
	}

	viper.ReadInConfig() // nolint: errcheck
	conf := &Conf{}
	if err := viper.Unmarshal(conf); err != nil {
		return conf, fmt.Errorf("unable to unmarshal conf: %w", err)
	}

	p := os.Getenv("PORT")
	if p != "" {
		port, err := strconv.Atoi(p)
		if err != nil {
			return conf, fmt.Errorf("given port isn't an integer: %w", err)
		}
		conf.Server.Port = port
	}

	return conf, nil
}
