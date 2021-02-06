package logger

import (
	"fmt"
	. "github.com/Unknwon/goconfig"
	"github.com/lios/go-activiti/common"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
)

const (
	ConfLoggerLevel           = "logger.level"
	ConfLoggerOutput          = "logger.output"
	ConfLoggerTimestampFormat = "logger.timestampFormat"
	ConfLoggerMaxItemSize     = "logger.maxLogItemSize"
	ConfLoggerFilename        = "logger.fileName"
	ConfLoggerMaxSize         = "logger.maxFileSize"
	ConfLoggerMaxDays         = "logger.maxFileDays"
	ConfLoggerMaxBackups      = "logger.maxFileBackups"
	ConfLoggerAsyncLog        = "logger.asyncLog"
	ConfLoggerAsyncBufferSize = "logger.asyncBufferSize"
)

var (
	logger         *aLogger
	callerInitOnce sync.Once
	logFile        = "/conf/logger.properties"
)

type aLogger struct {
	logger          *logrus.Logger
	loggerCloseOnce sync.Once
}

func init() {
	configFile := common.ReadConfig(logFile)
	if configFile == nil {
		fmt.Errorf("read log file:%s err", logFile)
		initLoggerWithConfig(GetDefaultConfig())
	} else {
		initLogConfig(configFile)
	}
}
func LoggerClose() {
	logger.logger.Exit(0)
}

func InitLoggerWithConfig(logFile string) {
	configFile := common.ReadConfig(logFile)
	if configFile == nil {
		fmt.Errorf("read log file:%s err", logFile)
	}
}
func initLogConfig(configFile *ConfigFile) {
	loggerConfig, err := configFile.GetSection("logger")
	if err != nil {
		Error("get logger conf error:", err)
		return
	}
	config := &Config{}
	config.Level = loggerConfig[ConfLoggerLevel]
	config.Output = loggerConfig[ConfLoggerOutput]
	config.FileName = loggerConfig[ConfLoggerFilename]
	config.TimestampFormat = loggerConfig[ConfLoggerTimestampFormat]
	config.AsyncBufferSize = viper.GetInt(loggerConfig[ConfLoggerAsyncBufferSize])
	config.AsyncLog = viper.GetBool(loggerConfig[ConfLoggerAsyncLog])
	config.MaxBackups = viper.GetInt(loggerConfig[ConfLoggerMaxBackups])
	config.MaxDays = viper.GetInt(loggerConfig[ConfLoggerMaxDays])
	config.MaxLogItemSize = viper.GetInt(loggerConfig[ConfLoggerMaxItemSize])
	config.MaxSize = viper.GetInt(loggerConfig[ConfLoggerMaxSize])

	initLoggerWithConfig(config)
}

func initLoggerWithConfig(config *Config) {
	logger = &aLogger{}
	logger.loggerInitWithConfig(config)
}

func getLogFields() logrus.Fields {
	fields := make(logrus.Fields)

	file, line, funcName := getCaller()

	fields[FieldLogName] = "logger"
	fields[FieldFileInfo] = file
	fields[FieldLineNumber] = line
	fields[FieldFunctionName] = funcName

	return fields
}

func (q *aLogger) loggerInitWithConfig(c *Config) {

	var out io.Writer
	var exitFunc func(int)

	out = os.Stdout
	exitFunc = os.Exit
	f := &Formatter{
		TimestampFormat: c.TimestampFormat,
		MaxLogItemSize:  c.MaxLogItemSize,
	}

	q.logger = &logrus.Logger{
		Out:          out,
		Formatter:    f,
		Level:        logrus.InfoLevel,
		ReportCaller: false,
		ExitFunc:     exitFunc,
	}

	//level
	level, err := logrus.ParseLevel(c.Level)
	if err != nil {
		return
	}
	q.logger.Level = level

	//output
	if c.Output == "file" {
		var w WriteSyncer
		w = &SyncFileWriter{}

		if c.AsyncLog {
			w = &AsyncWrapper{
				Writer:    w,
				ErrOutput: os.Stderr,
			}
		}

		if err := w.Setup(c); err != nil {
			return
		}

		out = w
		exitFunc = func(c int) {
			q.loggerCloseOnce.Do(func() {
				w.Sync()
				w.Close()
			})
			if c == 1 {
				os.Exit(c)
			}
		}

		q.logger.Out = out
		q.logger.ExitFunc = exitFunc
	}
}
func getCaller() (file string, line int, funcName string) {
	var loggerPackage string
	callerInitOnce.Do(func() {
		pcs := make([]uintptr, 2)
		_ = runtime.Callers(0, pcs)
		loggerPackage = getPackageName(runtime.FuncForPC(pcs[1]).Name())
	})

	pcs := make([]uintptr, 10)
	depth := runtime.Callers(1, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	var caller *runtime.Frame

	for f, again := frames.Next(); again; f, again = frames.Next() {
		pkg := getPackageName(f.Function)
		if pkg != loggerPackage {
			caller = &f
			break
		}
	}

	if caller == nil {
		return "unknow", 0, "unknow"
	}

	return getFileName(caller.File), caller.Line, caller.Function
}

func getPackageName(f string) string {
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, "/")
		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}

	return f
}

func getFileName(f string) string {
	lastSlash := strings.LastIndex(f, "/")

	f = f[lastSlash+1:]
	return f
}

func TraceF(fstr string, args ...interface{}) {
	logger.TraceF(fstr, args...)
}

func Trace(args ...interface{}) {
	logger.Trace(args...)
}

func DebugF(fstr string, args ...interface{}) {
	logger.DebugF(fstr, args...)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func InfoF(fstr string, args ...interface{}) {
	logger.InfoF(fstr, args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func WarnF(fstr string, args ...interface{}) {
	logger.WarnF(fstr, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func ErrorF(fstr string, args ...interface{}) {
	logger.ErrorF(fstr, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func FatalF(fstr string, args ...interface{}) {
	logger.FatalF(fstr, args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func (q *aLogger) TraceF(fstr string, args ...interface{}) {
	entry := q.logger.WithFields(getLogFields())
	entry.Tracef(fstr, args...)
}

func (q *aLogger) Trace(args ...interface{}) {
	entry := q.logger.WithFields(getLogFields())
	entry.Trace(args...)
}

func (q *aLogger) DebugF(fstr string, args ...interface{}) {
	entry := q.logger.WithFields(getLogFields())
	entry.Debugf(fstr, args...)
}

func (q *aLogger) Debug(args ...interface{}) {
	entry := q.logger.WithFields(getLogFields())
	entry.Debug(args...)
}

func (q *aLogger) InfoF(fstr string, args ...interface{}) {
	entry := q.logger.WithFields(getLogFields())
	entry.Infof(fstr, args...)
}

func (q *aLogger) Info(args ...interface{}) {
	entry := q.logger.WithFields(getLogFields())
	entry.Info(args...)
}

func (q *aLogger) WarnF(fstr string, args ...interface{}) {
	entry := q.logger.WithFields(getLogFields())
	entry.Warningf(fstr, args...)
}

func (q *aLogger) Warn(args ...interface{}) {
	entry := q.logger.WithFields(getLogFields())
	entry.Warning(args...)
}

func (q *aLogger) ErrorF(fstr string, args ...interface{}) {
	entry := q.logger.WithFields(getLogFields())
	entry.Errorf(fstr, args...)
}

func (q *aLogger) Error(args ...interface{}) {
	entry := q.logger.WithFields(getLogFields())
	entry.Error(args...)
}

func (q *aLogger) FatalF(fstr string, args ...interface{}) {
	entry := q.logger.WithFields(getLogFields())
	entry.Fatalf(fstr, args...)
}

func (q *aLogger) Fatal(args ...interface{}) {
	entry := q.logger.WithFields(getLogFields())
	entry.Fatal(args...)
}
