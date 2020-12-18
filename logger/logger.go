package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"path"
	"runtime"
)

// WriterHook is a hook that writes logs of specified LogLevels to specified Writer
type WriterHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

// Fire will be called when some logging function is called with current hook
// It will format log entry to string and write it to appropriate writer
func (hook *WriterHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.Writer {
		_, err = w.Write([]byte(line))
	}
	return err
}

// Levels define on which log levels this hook would trigger
func (hook *WriterHook) Levels() []logrus.Level {
	return hook.LogLevels
}

var L *logrus.Logger

func init() {
	L = logrus.New()
	L.SetReportCaller(true)
	L.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s:%d", filename, f.Line), fmt.Sprintf("%s()", f.Function)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	err := os.MkdirAll("logs", 0755)

	if err != nil || os.IsExist(err) {
		logrus.Error("can't create log dir. no configured logging to files")
	} else {
		allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
		if err != nil {
			panic(fmt.Sprintf("[Error]: %s", err))
		}

		L.SetOutput(ioutil.Discard) // Send all logs to nowhere by default

		L.AddHook(&WriterHook{
			Writer:    []io.Writer{allFile, os.Stdout},
			LogLevels: logrus.AllLevels,
		})
	}

	// Only logrus the warning severity or above.
	L.SetLevel(logrus.TraceLevel)
}
