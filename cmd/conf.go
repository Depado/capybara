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

// LogConf represents the logger configuration.
type LogConf struct {
	Level  string `mapstructure:"level"`
	Type   string `mapstructure:"type"`
	Caller bool   `mapstructure:"caller"`
}

// ServerConf represents the server configuration.
type ServerConf struct {
	Host string    `mapstructure:"host"`
	Port int       `mapstructure:"port"`
	TLS  TLSConfig `mapstructure:"tls"`
}

// TLSConfig represents the TLS configuration of the service.
type TLSConfig struct {
	CertPath string `mapstructure:"cert_path"`
	KeyPath  string `mapstructure:"key_path"`
	Type     string `mapstructure:"tls"`
}

// ListenAddr returns a formatted string to listen on.
func (s ServerConf) ListenAddr() string {
	l := s.Host
	if s.Port > 0 {
		l += ":" + strconv.Itoa(s.Port)
	}

	return l
}

// DatabaseConf stores the database configuration.
type DatabaseConf struct {
	Path                string        `mapstructure:"path"`
	MaxBucketsRecursion int           `mapstructure:"max_buckets_recursion"`
	DefaultLockTTL      time.Duration `mapstructure:"default_lock_ttl"`
}

// Conf holds the various configuration structures and is used to parse the
// config file if any.
type Conf struct {
	Log      LogConf      `mapstructure:"log"`
	Server   ServerConf   `mapstructure:"server"`
	Database DatabaseConf `mapstructure:"database"`
}

// NewLogger will return a new logger.
func NewLogger(conf *Conf) zerolog.Logger {
	// Level parsing
	warns := []string{}

	lvl, err := zerolog.ParseLevel(conf.Log.Level)
	if err != nil {
		warns = append(warns, fmt.Sprintf("unrecognized log level '%s', fallback to 'info'", conf.Log.Level))

		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		zerolog.SetGlobalLevel(lvl)
	}

	// Type parsing
	switch conf.Log.Type {
	case "console":
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	case "json":
		break
	default:
		warns = append(warns, fmt.Sprintf("unrecognized log type '%s', fallback to 'json'", conf.Log.Type))
	}

	// Caller
	if conf.Log.Caller {
		log.Logger = log.With().Caller().Logger()
	}

	// Log messages with the newly created logger
	for _, w := range warns {
		log.Warn().Msg(w)
	}

	return log.Logger
}

// NewConf will parse and return the configuration.
func NewConf() (*Conf, error) {
	// Environment variables
	viper.AutomaticEnv()
	viper.SetEnvPrefix("capybara")
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
