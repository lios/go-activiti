package logger

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
)

const (
	FieldTime         = "time"
	FieldLevel        = "level"
	FieldFileInfo     = "file"
	FieldLineNumber   = "lineNumber"
	FieldFunctionName = "functionName"
	FieldLogName      = "logName"
	FieldMessage      = "msg"
)

type Formatter struct {
	TimestampFormat string
	MaxLogItemSize  int
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {

	data := make(logrus.Fields)
	for k, v := range entry.Data {
		data[k] = v
	}

	fixedKeys := []string{FieldTime, FieldLevel, FieldLogName,
		FieldFileInfo, FieldLineNumber, FieldFunctionName}

	for k := range data {
		fixedKeys = append(fixedKeys, k)
	}

	fixedKeys = append(fixedKeys, FieldMessage)

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	for _, key := range fixedKeys {
		var value interface{}
		switch {
		case key == FieldTime:
			value = entry.Time.Format(f.TimestampFormat)
		case key == FieldLevel:
			value = entry.Level.String()
		case key == FieldMessage:
			value = entry.Message
		default:
			value = data[key]
		}
		f.appendKeyValue(b, key, value)
	}

	if len(b.Bytes()) >= f.MaxLogItemSize {
		truncatedBuf := b.Bytes()[:f.MaxLogItemSize]
		truncatedBuf = append(truncatedBuf, '\n')
		return truncatedBuf, nil
	}
	b.WriteByte('\n')
	return b.Bytes(), nil
}

func (f *Formatter) appendKeyValue(b *bytes.Buffer, key string, value interface{}) {
	b.WriteByte('|')
	f.appendValue(b, value)
}

func (f *Formatter) appendValue(b *bytes.Buffer, value interface{}) {
	stringVal, ok := value.(string)
	if !ok {
		stringVal = fmt.Sprint(value)
	}

	b.WriteString(stringVal)
}
