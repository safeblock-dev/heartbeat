#!/bin/bash

# Set default value for HEARTBEAT_FILE if not provided by the environment
HEARTBEAT_FILE="${HEARTBEAT_FILE:-/tmp/healthy}"

# Set default value for LIVENESS_TIMEOUT if not provided by the environment
LIVENESS_TIMEOUT="${LIVENESS_TIMEOUT:-30}"

# Check if the file exists
if [[ ! -f "$HEARTBEAT_FILE" ]]; then
  echo "File '$HEARTBEAT_FILE' does not exist."
  exit 1
fi

# Get the current time in seconds since epoch
current_time=$(date +%s)

# Get the last modification time of the file in seconds since epoch
# Use `stat` with different options based on the platform (macOS or Linux)
if stat --version &>/dev/null; then
  # GNU stat (Linux)
  file_mod_time=$(stat -c %Y "$HEARTBEAT_FILE")
else
  # BSD stat (macOS)
  file_mod_time=$(stat -f %m "$HEARTBEAT_FILE")
fi

# Calculate the time difference between now and the file's last modification
time_diff=$((current_time - file_mod_time))

# Check if the file's last modification time is within the liveness timeout
if [[ $time_diff -le $LIVENESS_TIMEOUT ]]; then
  # File is recent enough, exit with status 0 (success)
  exit 0
else
  # File is too old, exit with status 1 (failure)
  echo "File '$HEARTBEAT_FILE' is older than $LIVENESS_TIMEOUT seconds."
  exit 1
fi
