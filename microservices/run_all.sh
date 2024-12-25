#!/bin/bash

# Daftar folder proyek
projects=(
  "company"
  "division"
  "leaveDetails"
  "leaves"
  "position"
  "users"
)

pids=()

cleanup() {
  echo ""
  echo "Detected Ctrl+C. Stopping all running services..."
  for pid in "${pids[@]}"; do
    if kill -0 "$pid" 2>/dev/null; then
      echo "Terminating process $pid"
      kill "$pid"
    fi
  done
  echo "All services stopped."
  exit 0
}

trap cleanup SIGINT

# Loop untuk menjalankan semua main.go
for project in "${projects[@]}"; do
  echo "Starting project: $project..."

  if cd "$project"; then
    go run main.go &
    pid=$!
    if kill -0 "$pid" 2>/dev/null; then
      pids+=("$pid")
      echo "[$project] Running with PID $pid."
    else
      echo "[$project] Failed to start."
    fi
    cd ..
  else
    echo "Failed to open folder: $project"
  fi
  echo "-----------------------------------"
done

# Menunggu semua proses berjalan
if [ ${#pids[@]} -eq 0 ]; then
  echo "No services are running."
else
  echo "All services are running. Press Ctrl+C to stop."
  wait
fi
