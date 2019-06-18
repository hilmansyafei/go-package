package modules

import (
	"log"
	"time"

	"github.com/labstack/echo"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

type (
	Log struct {
		Logger *logrus.Logger
	}
)

func NewLogger() (*Log, error) {
	l := logrus.New()
	logf, err := rotatelogs.New(
		"storages/logs/access_log.%Y%m%d",

		// symlink current log to this file
		//rotatelogs.WithLinkName("/tmp/app_access.log"),

		// max : 7 days to keep
		rotatelogs.WithMaxAge(24*7*time.Hour),

		// rotate every day
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		log.Printf("failed to create rotatelogs: %s", err)
		return nil, err
	}

	l.Formatter = &logrus.JSONFormatter{}
	l.Out = logf
	l.Level = logrus.DebugLevel

	return &Log{
		Logger: l,
	}, nil
}

// GenLog for general log
func GenLog(c echo.Context, dataRequest interface{}, resp interface{}, info string) {
	log, errLog := NewLogger()
	if errLog != nil {
		panic(errLog)
	}
	// Create log
	log.Logger.WithFields(logrus.Fields{
		"remote_ip": c.RealIP(),
		"protocol":  c.Request().Proto,
		"host":      c.Request().Host,
		"uri":       c.Request().RequestURI,
		"headers":   c.Request().Header,
		"request":   dataRequest,
		"response":  resp,
	}).Info(info)
}
