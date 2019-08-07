package logger

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime/debug"

	"wume-composer/internal/pkg/common/config"
)

const (
	debugColor   = "\033[1;36m%s\033[0m"
	infoColor    = "\033[1;34m%s\033[0m"
	warningColor = "\033[1;33m%s\033[0m"
	errorColor   = "\033[1;31m%s\033[0m"
	fatalColor   = "\033[0;35m%s\033[0m"
)

var (
	logger = newLogger()
	files  []*os.File
)

func getWriter(conf *config.LoggerOut, prefix string, color string) logWriter {
	writer := ioutil.Discard
	format := "%s"

	switch conf.Mode {
	case "file":
		if conf.Filename == "" {
			conf.Filename = "default.log"
		}
		file, err := os.OpenFile(conf.Filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic("error opening log file: " + err.Error())
		}
		files = append(files, file)
		writer = file
	case "console":
		writer = os.Stdout
		format = color
	}

	return logWriter{log.New(writer, prefix, log.LstdFlags), format}
}

type logWriter struct {
	logger *log.Logger
	format string
}

type logMessage struct {
	writer  logWriter
	message string
}

type chanLogger struct {
	logChan   chan logMessage
	fatalChan chan interface{}

	DebugLogger   logWriter
	InfoLogger    logWriter
	WarningLogger logWriter
	ErrorLogger   logWriter
	FatalLogger   logWriter
}

func newLogger() *chanLogger {
	l := &chanLogger{
		logChan:   make(chan logMessage, 256),
		fatalChan: make(chan interface{}, 256),
	}

	l.DebugLogger = getWriter(&config.Logger.Debug, "[DEBUG]: ", debugColor)
	l.InfoLogger = getWriter(&config.Logger.Info, "[INFO]: ", infoColor)
	l.WarningLogger = getWriter(&config.Logger.Warning, "[WARNING]: ", warningColor)
	l.ErrorLogger = getWriter(&config.Logger.Error, "[ERROR]: ", errorColor)
	l.FatalLogger = getWriter(&config.Logger.Fatal, "[FATAL]: ", fatalColor)

	l.Run()
	return l
}

func (l *chanLogger) Log(logger logWriter, formatMessage string, values ...interface{}) {
	l.logChan <- logMessage{logger, fmt.Sprintf(formatMessage, values...)}
}

func (l *chanLogger) LogFatal(formatMessage string, values ...interface{}) {
	l.fatalChan <- fmt.Sprintf(formatMessage+"\n"+string(debug.Stack())+"\n", values...)
}

func (l *chanLogger) Run() {
	go func() {
		defer func() {
			for _, f := range files {
				err := f.Close()
				if err != nil {
					panic(err)
				}
			}
		}()

		for {
			select {
			case message := <-l.logChan:
				message.writer.logger.Printf(message.writer.format, message.message)
			case message := <-l.fatalChan:
				l.FatalLogger.logger.Fatalf(l.FatalLogger.format, message)
			}
		}
	}()
}

func Debug(formatMessage string, values ...interface{}) {
	logger.Log(logger.DebugLogger, formatMessage, values...)
}

func Info(formatMessage string, values ...interface{}) {
	logger.Log(logger.InfoLogger, formatMessage, values...)
}

func Warn(formatMessage string, values ...interface{}) {
	logger.Log(logger.WarningLogger, formatMessage, values...)
}

func Error(formatMessage string, values ...interface{}) {
	logger.Log(logger.ErrorLogger, formatMessage, values...)
}

func Fatal(formatMessage string, values ...interface{}) {
	logger.LogFatal(formatMessage, values...)
}
