package logger

import (
	"context"
	"fmt"
	"io"
)

type AsyncWrapper struct {
	bufferChan chan string
	ctx        context.Context
	cancel     context.CancelFunc
	Writer     WriteSyncer
	ErrOutput  io.Writer
}

func (asfw *AsyncWrapper) bufferPump() {

	defer func() {
		if err := recover(); err != nil {
			asfw.ErrOutput.Write([]byte(fmt.Sprintf("receive log msg fail: %+v", err)))
		}
	}()

	for {
		buffer, ok := <-asfw.bufferChan
		if !ok {
			asfw.cancel()
			break
		}

		n, err := asfw.Writer.Write([]byte(buffer))

		if err != nil || n < len(buffer) {
			asfw.ErrOutput.Write([]byte(fmt.Sprintf("write log msg fail: %+v origin len: %d current len: %d", err, len(buffer), n)))
		}
	}
}

func (asfw *AsyncWrapper) Setup(config *Config) error {

	asfw.bufferChan = make(chan string, config.AsyncBufferSize)
	asfw.ctx, asfw.cancel = context.WithCancel(context.Background())

	if err := asfw.Writer.Setup(config); err != nil {
		return err
	}

	go asfw.bufferPump()

	return nil
}

func (asfw *AsyncWrapper) Write(bs []byte) (int, error) {

	defer func() {
		if err := recover(); err != nil {
			asfw.ErrOutput.Write([]byte(fmt.Sprintf("send log msg fail: %+v", err)))
		}
	}()

	return 0, nil
}

func (asfw *AsyncWrapper) Sync() error {

	close(asfw.bufferChan)
	<-asfw.ctx.Done()

	if err := asfw.Writer.Sync(); err != nil {
		return err
	}

	return nil
}

func (asfw *AsyncWrapper) Close() error {
	if err := asfw.Writer.Close(); err != nil {
		return err
	}
	return nil
}
