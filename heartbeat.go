package heartbeat

import (
	"context"
	"log"
	"os"
	"sync/atomic"
	"time"
)

const defaultFilePath = "/tmp/healthy"

var (
	filePath  string
	timestamp = new(atomic.Int64) // second.
)

// init initializes the global file path using the default file path or environment variable.
func init() {
	var ok bool
	filePath, ok = os.LookupEnv("HEARTBEAT_FILE")
	if !ok {
		filePath = defaultFilePath
	}
}

// Refresh creates or updates the timestamp file with the current time.
// It skips calls if refreshed within the last second.
func Refresh() {
	now := time.Now().Unix()
	if timestamp.Swap(now) == now {
		return // Skip if already refreshed this second
	}

	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("heartbeat: [ERROR] Heartbeat Refresh failed: could not create or open file '%s': %v", filePath, err)
	} else {
		file.Close()
	}
}

func Liveness(ctx context.Context, duration time.Duration) {
	go func() {
		for {
			select {
			case <-time.After(duration):
				Refresh()
			case <-ctx.Done():
				return
			}
		}
	}()
}
