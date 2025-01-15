package heartbeat

import (
	"log"
	"os"
	"sync/atomic"
	"time"
)

const defaultFilePath = "/tmp/healthy"

var (
	filePath  string
	timestamp = new(atomic.Int64)
) // second.

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
	current := time.Now().Local()
	if timestamp.Load() == current.Unix() {
		return // skip.
	}

	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		var file *os.File
		file, err = os.Create(filePath)
		if err != nil {
			log.Printf("heartbeat: [ERROR] Heartbeat Refresh failed: could not create or open file '%s': %v", filePath, err)
		}
		defer file.Close()
	case err != nil:
		log.Printf("heartbeat: [ERROR] Heartbeat Refresh failed: could not create or open file '%s': %v", filePath, err)
	default:
		err = os.Chtimes(filePath, current, current)
		if err != nil {
			log.Printf("heartbeat: [ERROR] Heartbeat Refresh failed: could not modification times '%s': %v", filePath, err)
		}
		timestamp.Store(current.Unix())
	}
}
