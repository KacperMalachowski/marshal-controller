package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	LogDir string
	file   *os.File
}

type Options = func(*Logger)

func New(options ...Options) (*Logger, error) {

	logger := &Logger{
		LogDir: "logs",
	}
	for _, option := range options {
		option(logger)
	}

	ts := time.Now().Format(time.RFC3339)
	filename := fmt.Sprintf("log-%s.log", strings.Replace(strings.Replace(ts, ":", "", -1), "-", "", -1))

	var path string
	if filepath.IsAbs(logger.LogDir) {
		path = filepath.Join(logger.LogDir, filename)
	} else {
		ex, err := os.Executable()
		if err != nil {
			return nil, errors.WithMessage(err, "cannot find executable path")
		}
		path = filepath.Join(filepath.Dir(ex), logger.LogDir, filename)
	}

	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return nil, errors.WithMessage(err, "cannot make directory")
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, errors.WithMessage(err, "cannot open log file")
	}

	logger.file = file
	logrus.SetOutput(file)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.TraceLevel)

	return logger, nil
}

func (l *Logger) Print(message string) {
	logrus.Print(message)
}

func (l *Logger) Trace(message string) {
	logrus.Trace(message)
}

func (l *Logger) Debug(message string) {
	logrus.Debug(message)
}

func (l *Logger) Info(message string) {
	logrus.Info(message)
}

func (l *Logger) Warning(message string) {
	logrus.Warning(message)
}

func (l *Logger) Error(message string) {
	logrus.Error(message)
}

func (l *Logger) Fatal(message string) {
	logrus.Fatal(message)
}
