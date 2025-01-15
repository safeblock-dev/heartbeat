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
	current := time.Now()
	unix := current.Unix()
	if timestamp.Swap(unix) == unix {
		return // Skip if already refreshed this second
	}

	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		// Create the file if it doesn't exist
		var file *os.File
		file, err = os.Create(filePath)
		if err == nil {
			log.Printf("heartbeat: [ERROR] Unable to create file '%s': %v", filePath, err)

			return
		}
		_ = file.Close()
	case err != nil:
		// Log unexpected errors while accessing the file
		log.Printf("heartbeat: [ERROR] Unable to access file '%s': %v", filePath, err)
	default:
		// Update file modification times
		err = os.Chtimes(filePath, current, current)
		if err != nil {
			log.Printf("heartbeat: [ERROR] Failed to update modification times for '%s': %v", filePath, err)
		}
	}
}
