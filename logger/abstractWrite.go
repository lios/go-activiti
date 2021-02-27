package logger

import (
	"io"
)

type WriteSyncer interface {
	io.WriteCloser
	Setup(config *Config) error
	Sync() error
}
