package common

import (
	"github.com/getsentry/raven-go"
	"github.com/spf13/viper"
)

func InitSentry() {
	raven.SetDSN(viper.GetString("sentry.dsn"))
	raven.SetDefaultLoggerName(viper.GetString("sentry.default_logger_name"))
	raven.SetDebug(viper.GetBool("basic.debug"))
	raven.SetRelease("behavior@" + VERSION)
	raven.SetSampleRate((float32)(viper.GetFloat64("sentry.sample_rate")))
}
