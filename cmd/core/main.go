package main

import (
	"fmt"

	"wume-composer/internal/app/core"
)

func main() {
	err := core.StartApp()
	if err != nil {
		core.StopApp()
		fmt.Println(err)
	}
}
