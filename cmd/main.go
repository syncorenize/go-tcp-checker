package main

import (
	"flag"
	"fmt"
	"go-tcp-checker/internal/checker"
	"go-tcp-checker/internal/config"
	"os"
)

func main() {
	configPath := flag.String("config", "", "Path to the configuration file.")
	flag.Parse()

	// Check if the config flag was provided
	if *configPath == "" {
		fmt.Println("No configuration file specified. Please use the -config flag to specify the path to the config.yaml file.")
		os.Exit(1)
	}

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		fmt.Printf("Failed to load configuration from '%s': %v\n", *configPath, err)
		os.Exit(1)
	}

	tcpChecker := &checker.TCPServiceChecker{}
	fmt.Println("Checking TCP services concurrently...")
	checker.CheckServicesConcurrentlyWithContext(tcpChecker, cfg)
}
