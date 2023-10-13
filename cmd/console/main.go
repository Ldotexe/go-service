package main

import (
	"fmt"
	"os"

	console "homework-4/internal/console/core"
)

func main() {
	err := console.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
