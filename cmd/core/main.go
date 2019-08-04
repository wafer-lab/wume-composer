package main

import (
	"wume-composer/internal/app/core"
	"wume-composer/internal/pkg/config"
	"wume-composer/internal/pkg/logger"
)

func main() {
	err := core.StartApp(core.Params{
		Port:   config.Core.Port,
		Prefix: config.Core.Prefix,
	})
	if err != nil {
		core.StopApp()
		logger.Error(err.Error())
	}
}
