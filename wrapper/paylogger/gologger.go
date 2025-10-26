package paylogger

import "github.com/donnie4w/go-logger/logger"

type LoggerOption struct {
	Level     string `yaml:"level"`
	Output    string `yaml:"output"`
	Path      string `yaml:"path"`
	TimeMode  string `yaml:"time-mode"`
	MaxSizeMB int    `yaml:"max-size-mb"`
	MaxBackup int    `yaml:"max-backup"`
	Compress  bool   `yaml:"compress"`
}

func NewGoLogger(option LoggerOption) *logger.Logging {
	goption := &logger.Option{
		Format:  logger.FORMAT_LEVELFLAG | logger.FORMAT_DATE | logger.FORMAT_TIME,
		Console: true,
	}
	if option.Output == "file" {
		fileOption := &logger.FileMixedMode{
			Filename:   option.Path,
			Maxsize:    int64(option.MaxSizeMB) * int64(logger.MB),
			Maxbuckup:  option.MaxBackup,
			IsCompress: option.Compress,
		}

		switch option.TimeMode {
		case "day":
			fileOption.Timemode = logger.MODE_DAY
		case "hour":
			fileOption.Timemode = logger.MODE_HOUR
		case "month":
			fileOption.Timemode = logger.MODE_MONTH
		default:
			fileOption.Timemode = logger.MODE_MONTH
		}

		goption.FileOption = fileOption
		goption.Console = false
	}

	return logger.NewLogger().SetOption(goption)
}
