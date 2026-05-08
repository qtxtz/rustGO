package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	pidFlag := flag.Int("pid", 0, "host process PID for lifecycle binding")
	flag.Parse()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	client := NewCGOClient()

	if !client.HealthCheck() {
		fmt.Println("rustGO SO not loaded")
		if *pidFlag > 0 {
			proc, _ := os.FindProcess(*pidFlag)
			if proc != nil {
				proc.Signal(syscall.SIGTERM)
			}
		}
		os.Exit(1)
	}
	defer client.Close()

	fmt.Println("=== rustGO Script ===")

	go func() {
		<-exit
		fmt.Println("Script stopped by signal")
		client.Close()
		os.Exit(0)
	}()

	time.Sleep(1 * time.Second)
	runDemoCGO(client)
}
