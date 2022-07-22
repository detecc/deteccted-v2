package logging

import (
	"github.com/detecc/deteccted-v2/internal/models/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	lSyslog "github.com/sirupsen/logrus/hooks/syslog"
	"log/syslog"
	"time"
)

const (
	RemoteLogging  = LogType("remote")
	FileLogging    = LogType("file")
	ConsoleLogging = LogType("console")

	Syslog = LogFormat("syslog")
	Json   = LogFormat("json")

	logFilePath = "/var/log/deteccted/client.log"
)

type (
	LogType   string
	LogFormat string
)

// Setup set up all logs
func Setup(logger *log.Logger, loggingConfig config.Logging, isDebug bool) {
	var (
		// Default (production) logging settings
		logLevel                = log.WarnLevel
		formatter log.Formatter = &log.JSONFormatter{}
		logFormat               = LogFormat(loggingConfig.Format)
	)

	if isDebug {
		logLevel = log.DebugLevel
	}

	logger.SetFormatter(formatter)
	logger.SetLevel(logLevel)

	for _, logType := range loggingConfig.Type {
		switch LogType(logType) {
		case FileLogging:
			fileLogging(logger, isDebug, logFilePath)
			break
		case RemoteLogging:
			remoteLogging(logger, loggingConfig.Address, logFormat)
			break
		case ConsoleLogging:
			break
		}
	}
}

// remoteLogging sends logs remotely to Graylog or any Syslog receiver.
func remoteLogging(logger *log.Logger, address string, format LogFormat) {

	switch format {
	case Json:
		break
	case Syslog:
		hook, err := lSyslog.NewSyslogHook(
			"tcp",
			address,
			syslog.LOG_WARNING,
			"deteccted",
		)
		if err == nil {
			logger.AddHook(hook)
		}
		break
	}
}

// fileLogging sets up the logging to file.
func fileLogging(logger *log.Logger, isDebug bool, path string) {
	writer, err := rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(604800)*time.Second),
	)
	if err != nil {
		return
	}

	writerMap := make(lfshook.WriterMap)
	writerMap[log.InfoLevel] = writer
	writerMap[log.ErrorLevel] = writer

	if isDebug {
		writerMap[log.DebugLevel] = writer
	}

	hook := lfshook.NewHook(
		writerMap,
		&log.JSONFormatter{},
	)

	logger.AddHook(hook)
}
