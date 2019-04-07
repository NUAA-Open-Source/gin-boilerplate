package common

import (
	"log"
	"os"

	"github.com/getsentry/raven-go"
	"github.com/spf13/viper"
)

var (
	logFile *os.File
	err     error
)

func InitLogger() {

	os.Mkdir("log", os.ModePerm|os.ModeDir)
	logFile, err = os.OpenFile(viper.GetString("log.file"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		raven.CaptureError(err, map[string]string{"type": "log"})
		log.Fatal(err)
	}
	if !viper.GetBool("basic.debug") {
		log.SetOutput(logFile)
	}
}

func GetLogFile() *os.File {
	return logFile
}
