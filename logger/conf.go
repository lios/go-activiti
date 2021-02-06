package logger

const (
	FileOutput             = "file"
	DefaultLevel           = "debug"
	DefaultOutput          = "console"
	DefaultTimestampFormat = "2020-01-02 15:04:05.000000,123"
	DefaultMaxLogItemSize  = 4 * 1024
	DefaultFileName        = "./log.log"
	DefaultMaxSize         = 300
	DefaultMaxDays         = 30
	DefaultMaxBackups      = 0
	DefaultAsyncLog        = false
	DefaultAsyncBufferSize = 2000
)

type Config struct {
	// Log level.
	Level string

	//file or console
	Output string

	//time format
	TimestampFormat string

	//max log item size
	MaxLogItemSize int

	// Log filename, leave empty to disable file log.
	FileName string

	// Max size for a single file, in MB.
	MaxSize int

	// Max log keep days, default is never deleting.
	MaxDays int

	// Maximum number of old log files to retain.
	MaxBackups int

	// Whether to write logs asynchronously.
	AsyncLog bool

	//async buffer size
	AsyncBufferSize int
}

func GetDefaultConfig() *Config {
	config := &Config{}

	config.Level = DefaultLevel
	config.Output = DefaultOutput
	config.TimestampFormat = DefaultTimestampFormat
	config.MaxLogItemSize = DefaultMaxLogItemSize
	config.FileName = DefaultFileName
	config.MaxSize = DefaultMaxSize
	config.MaxDays = DefaultMaxDays
	config.MaxBackups = DefaultMaxBackups
	config.AsyncLog = DefaultAsyncLog
	config.AsyncBufferSize = DefaultAsyncBufferSize

	return config
}
