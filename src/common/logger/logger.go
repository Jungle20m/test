package logger

import (
	"time"

	zrlog "github.com/rs/zerolog"

	configuration "golang-structure/src/configs"

	"tawesoft.co.uk/go/log"
	"tawesoft.co.uk/go/log/zerolog"
)

var ProcessLog *zrlog.Logger

func InitLogger() {
	cfg := log.Config{
		Syslog: log.ConfigSyslog{
			Enabled:  false,
			Network:  "",
			Address:  "",
			Priority: log.LOG_ERR | log.LOG_DAEMON,
			Tag:      "example",
		},
		File: log.ConfigFile{
			Enabled:          configuration.Config.GetBool("logger.enable"),
			Mode:             0600,
			Path:             configuration.Config.GetString("logger.path"),
			Rotate:           configuration.Config.GetBool("logger.rotate.enable"),
			RotateCompress:   configuration.Config.GetBool("logger.rotate.compress"),
			RotateMaxSize:    configuration.Config.GetInt("logger.rotate.max_size_mb") * 1024 * 1024,
			RotateKeepAge:    10 * 24 * time.Hour,
			RotateKeepNumber: 32, // 32 * 64 MB = 2 GB max storage (before compression)
		},
		Stderr: log.ConfigStderr{
			Enabled: false,
			Color:   true,
		},
	}

	logger, closer, err := zerolog.New(cfg)
	if err != nil {
		panic(err)
	}
	defer closer()

	ProcessLog = &logger
}
