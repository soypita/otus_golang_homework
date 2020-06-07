package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger(logOut io.Writer, logLevel string) (logrus.FieldLogger, error) {
	log := logrus.New()
	Formatter := new(logrus.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, fmt.Errorf("error while configure logger : %w", err)
	}
	log.SetLevel(level)
	mw := io.MultiWriter(os.Stdout, logOut)
	log.SetOutput(mw)
	return log, nil
}
