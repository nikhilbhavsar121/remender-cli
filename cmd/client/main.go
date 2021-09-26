package main

import (
	"Projects/reminders-cli/client"
	"flag"
	"fmt"
	"os"
)

var (
	backendURIFlag = flag.String("backend", "https://localhost:5081", "Backend API URI")
	helpFlag       = flag.Bool("help", false, "Display a helpful message")
)

func main() {
	flag.Parse()
	s := client.NewSwitch(*backendURIFlag)

	if *helpFlag || len(os.Args) == 1 {
		s.Help()
		return
	}

	err := s.Switch()
	if err != nil {
		fmt.Printf("cmd switch error: %v\n", err)
		os.Exit(2)
	}
}
