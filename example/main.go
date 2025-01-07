package main

import (
	"fmt"
	"time"

	"github.com/safeblock-dev/heartbeat"
)

const delay = 5 * time.Second

func main() {
	fmt.Println("Starting heartbeat...")

	// Refresh the timestamp file with default or environment-specified path
	heartbeat.Refresh()
	fmt.Println("Heartbeat file updated.")

	// Simulate a delay
	time.Sleep(delay)

	// Refresh again
	heartbeat.Refresh()
	fmt.Println("Heartbeat file updated again.")
}
