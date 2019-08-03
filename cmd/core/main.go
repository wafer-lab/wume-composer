package main

import (
	"fmt"

	"wume-composer/internal/app/core"
	"wume-composer/internal/pkg/config"
)

func main() {
	err := core.StartApp(core.Params{
		Port:   config.Core.Port,
		Prefix: config.Core.Prefix,
	})
	if err != nil {
		core.StopApp()
		fmt.Println(err)
	}
}
