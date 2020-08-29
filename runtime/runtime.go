package runtime

import (
	"bytes"
	"runtime"
)

func GoroutineId() string {
	buf := make([]byte, 30)
	buf = buf[:runtime.Stack(buf, false)]
	return string(bytes.Split(buf, []byte(" "))[1])
}
