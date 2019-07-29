package main

import (
	"fmt"
	"os"

	"wume-composer/internal/app/core"
)

func main() {
	params := core.Params{Port: os.Getenv("PORT")}
	if params.Port == "" {
		params.Port = "6002"
	}

	err := core.StartApp(params)
	if err != nil {
		core.StopApp()
		fmt.Println(err)
	}
}
