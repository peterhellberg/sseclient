package main

import (
	"fmt"
	"os"

	"github.com/peterhellberg/sseclient"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: sseclient <URL>")
		os.Exit(0)
	}

	events, err := sseclient.OpenURL(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	for event := range events {
		fmt.Println(event.Name, event.Data)
	}
}
