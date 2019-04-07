package common

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/getsentry/raven-go"
	"github.com/spf13/viper"
)

func SetConfig() error {
	viper.SetConfigName("example") // name of config file (without extension)
	viper.AddConfigPath("conf")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("EXAMPLE")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Println("Fatal error config file:", err)
		raven.CaptureError(err, map[string]string{"type": "config"})
		return err
	}

	return nil
}

func WatchConfig() error {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})
	return nil
}

func DefaultConfig() error {
	// basic default values
	viper.SetDefault("basic.debug", true)
	viper.SetDefault("basic.maintainance", false)
	viper.SetDefault("basic.port", "8080")
	// storage default values
	viper.SetDefault("storage.mysql.user", "root")
	viper.SetDefault("storage.mysql.password", "")
	viper.SetDefault("storage.mysql.host", "localhost")
	viper.SetDefault("storage.mysql.port", "3306")
	viper.SetDefault("sotrage.mysql.database", "example")
	viper.SetDefault("storage.mysql.timezone", "Asia%2FShanghai")
	viper.SetDefault("storage.mysql.retry_interval", 20)
	viper.SetDefault("storage.mysql.max_idle_conns", 30)
	viper.SetDefault("storage.mysql.max_open_conns", 100)
	// sentry default values
	viper.SetDefault("sentry.dsn", "")
	viper.SetDefault("sentry.default_logger_name", "example")
	viper.SetDefault("sentry.sample_rate", 1)
	// log default values
	viper.SetDefault("log.file", "log/example.log")

	return nil
}
