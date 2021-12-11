package logutil

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func NewLogger(logName, level string) *logrus.Logger {
	writer, err := rotatelogs.New(
		logName+".%Y%m%d%H%M%S",
		// 设置日志分割的时间，设置为24小时分割一次
		rotatelogs.WithRotationTime(time.Hour*24),
		// 为最新的日志建立软连接，以方便随着找到当前日志文件
		rotatelogs.WithLinkName(logName),
		// 设置文件清理前最多保存的个数
		rotatelogs.WithRotationCount(20),
	)
	if err != nil {
		logrus.Fatalf("init log error: %v", err)
	}

	Log = logrus.New()

	lel, err := logrus.ParseLevel(level)

	if err != nil {
		Log.SetLevel(logrus.InfoLevel)
	} else {
		Log.SetLevel(lel)
	}
	Log.Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.DebugLevel: writer,
			logrus.InfoLevel:  writer,
			logrus.WarnLevel:  writer,
			logrus.ErrorLevel: writer,
			logrus.FatalLevel: writer,
			logrus.PanicLevel: writer,
		}, &logrus.JSONFormatter{},
	))
	return Log
}

type Vlog struct {
	LogName string
	Level   string
}
