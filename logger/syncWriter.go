package logger

import "gopkg.in/natefinch/lumberjack.v2"

type SyncFileWriter struct {
	lumberjack.Logger
}

func (sfw *SyncFileWriter) Sync() error {
	return nil
}

func (sfw *SyncFileWriter) Setup(config *Config) error {

	sfw.Filename = config.FileName
	sfw.MaxSize = config.MaxSize
	sfw.MaxAge = config.MaxDays
	sfw.MaxBackups = config.MaxBackups
	sfw.LocalTime = true

	return nil
}
