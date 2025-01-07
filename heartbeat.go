package heartbeat

import (
	"log"
	"os"
	"sync/atomic"
	"time"
)

const defaultFilePath = "heartbeat.timestamp"

var (
	filePath      string
	lastRefreshTS = new(atomic.Int64)
)

// init initializes the global file path using the default file path or environment variable.
func init() {
	filePath = os.Getenv("HEARTBEAT_FILE")
	if filePath == "" {
		filePath = defaultFilePath
	}
}

// Refresh creates or updates the timestamp file with the current time.
// It skips calls if refreshed within the last second.
func Refresh() {
	now := time.Now().Unix()
	if lastRefreshTS.Swap(now) == now {
		return // Skip if already refreshed this second
	}

	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("[WARNING] Heartbeat Refresh failed: could not create or open file '%s': %v", filePath, err)

		return
	}
	defer file.Close()

	_, err = file.WriteString(time.Now().Format(time.RFC3339))
	if err != nil {
		log.Printf("[WARNING] Heartbeat Refresh failed: could not write to file '%s': %v", filePath, err)
	}
}
