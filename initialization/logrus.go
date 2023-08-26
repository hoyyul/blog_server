package initialization

import (
	"blog_server/global"
	"bytes"
	"fmt"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

// color
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}

// implement Formatter(entry *logrus.Entry) ([]byte, error) interface
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//diy color
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	logger := global.Config.Logger
	//diy date
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//path
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//diy format
		fmt.Fprintf(b, "[%s][%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", logger.Prefix, timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "[%s][%s] \x1b[%dm[%s]\x1b[0m %s\n", logger.Prefix, timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

func InitLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)                           //stdout
	logger.SetReportCaller(global.Config.Logger.ShowLine) //show line
	logger.SetFormatter(&LogFormatter{})
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)
	return logger
}
