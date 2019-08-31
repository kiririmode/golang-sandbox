package main

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	log := &logrus.Logger{
		Out: os.Stderr,
		Formatter: &logrus.JSONFormatter{
			TimestampFormat: time.RFC3339Nano,
		},
		Level: logrus.InfoLevel,
	}

	log.WithField("user", "kiririmode").Infoln("hello world")
}
