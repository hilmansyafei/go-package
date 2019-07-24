package modules

import (
	"log"
	"time"

	"github.com/hilmansyafei/go-package/status"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

// LogProvider : provider logger
type LogProvider interface {
	GenLog(c status.Log, dataRequest interface{}, resp interface{}, info string) error
	GenErrLog(logData map[string]interface{}, msg string) error
	GenReqLog(logData map[string]interface{}, msg string) error
}

// Log : handler
type Log struct {
	Logger *logrus.Logger
}

// NewLogger : init Log
func NewLogger(logPath string) (*Log, error) {
	l := logrus.New()
	logf, err := rotatelogs.New(
		logPath+".%Y%m%d",

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

// GenErrLog : general logger
func (log *Log) GenErrLog(logData map[string]interface{}, msg string) error {
	log.Logger.WithFields(logData).Error(msg)
	return nil
}

// GenLog for general log
func (log *Log) GenLog(c status.Log, dataRequest interface{}, resp interface{}, info string) error {
	// Create log
	log.Logger.WithFields(logrus.Fields{
		"remote_ip": c.IP,
		"protocol":  c.Protocol,
		"host":      c.Host,
		"uri":       c.URI,
		"headers":   c.Headers,
		"request":   dataRequest,
		"response":  resp,
	}).Info(info)

	return nil
}

// GenReqLog for request log
func (log *Log) GenReqLog(logData map[string]interface{}, msg string) error {
	log.Logger.WithFields(logData).Info(msg)
	return nil
}
