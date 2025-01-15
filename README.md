# Heartbeat Package

The `heartbeat` package provides a simple mechanism for creating or updating a timestamp file that can be used to monitor the health or activity of an application.

## Features

- Automatically initializes with a default file path or a path specified via an environment variable.
- Provides a `Refresh` function to create or update the file with the current timestamp.
- Logs warnings in case of errors during file operations.

## Installation

```bash
go get github.com/safeblock-dev/heartbeat
```

## Usage

### Basic Example

```go
package main

import (
	"fmt"
	"time"

	"github.com/safeblock-dev/heartbeat"
)

func main() {
	fmt.Println("Starting heartbeat...")

	// Update the heartbeat file with the current timestamp
	heartbeat.Refresh()
	fmt.Println("Heartbeat file updated.")

	// Simulate a delay
	time.Sleep(5 * time.Second)

	// Update the heartbeat file again
	heartbeat.Refresh()
	fmt.Println("Heartbeat file updated again.")
}
```

### Custom File Path

You can specify a custom file path for the heartbeat file by setting the `HEARTBEAT_FILE` environment variable:

```bash
export HEARTBEAT_FILE=/path/to/custom/heartbeat.timestamp
```

The package will use the specified file path instead of the default (`heartbeat.timestamp`).

### File Check Script

A [script](check_liveness.sh) can be used alongside this package to check the heartbeat file. This script verifies the file's existence and age, returning a status of 0 if the file has been updated within a specified timeout (default: 30 seconds) or 1 otherwise.

#### Example Integration

You can integrate this script into your application for health checks. For example, in Kubernetes, it can be used as a liveness probe:

```yaml
livenessProbe:
  exec:
    command:
    - "/path/to/check_liveness.sh"
  initialDelaySeconds: 5
  periodSeconds: 10
```

## Logging

The package logs warnings in case of errors during file operations, such as:

- Failure to create or open the file.
- Failure to write to the file.
- Failure to close the file.

All warnings are logged using the standard logger.