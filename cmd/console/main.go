package main

import (
	"fmt"
	"os"

	console "homework-6/internal/console/core"
)

func main() {
	err := console.Run(os.Args[1:])
	if err != nil {
		fmt.Println(err)
	}
}
