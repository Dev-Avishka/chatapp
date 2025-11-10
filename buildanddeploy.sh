#!/usr/bin/env bash
set -e

# Termux-friendly build & deploy script for the chatapp Go project
# Usage: ./buildanddeploy.sh build|start|stop|restart|status|run

BINARY_NAME="chatapp"
SRC_DIR="$(cd "$(dirname "$0")" && pwd)"
LOG_FILE="$SRC_DIR/$BINARY_NAME.log"
PID_FILE="$SRC_DIR/$BINARY_NAME.pid"

# Ensure go is available
if ! command -v go >/dev/null 2>&1; then
  echo "go not found. Install Go in Termux: pkg install golang" >&2
  exit 1
fi

cmd="$1"
case "$cmd" in
  build)
    echo "Downloading modules..."
    (cd "$SRC_DIR" && go mod download)
    echo "Building..."
    (cd "$SRC_DIR" && go build -o "$BINARY_NAME" .)
    echo "Built $SRC_DIR/$BINARY_NAME"
    ;;
  start|run)
    if [ -f "$PID_FILE" ] && kill -0 "$(cat "$PID_FILE")" 2>/dev/null; then
      echo "$BINARY_NAME is already running (pid $(cat $PID_FILE))"
      exit 0
    fi
    if [ ! -x "$SRC_DIR/$BINARY_NAME" ]; then
      echo "Binary not found or not executable, building first..."
      (cd "$SRC_DIR" && go build -o "$BINARY_NAME" .)
    fi
    echo "Starting $BINARY_NAME (listens on port defined in main.go)..."
    nohup "$SRC_DIR/$BINARY_NAME" > "$LOG_FILE" 2>&1 &
    echo $! > "$PID_FILE"
    echo "Started with pid $(cat $PID_FILE). Logs: $LOG_FILE"
    ;;
  stop)
    if [ -f "$PID_FILE" ]; then
      pid=$(cat "$PID_FILE")
      if kill -0 "$pid" 2>/dev/null; then
        echo "Stopping $BINARY_NAME (pid $pid)..."
        kill "$pid"
        sleep 1
        rm -f "$PID_FILE"
        echo "Stopped."
      else
        echo "Process $pid not running. Removing stale PID."
        rm -f "$PID_FILE"
      fi
    else
      echo "No pid file, attempting to find process by name..."
      pkill -f "$BINARY_NAME" || echo "No process found"
    fi
    ;;
  restart)
    "$0" stop
    sleep 1
    "$0" start
    ;;
  status)
    if [ -f "$PID_FILE" ] && kill -0 "$(cat "$PID_FILE")" 2>/dev/null; then
      echo "$BINARY_NAME running, pid $(cat $PID_FILE)"
    else
      echo "$BINARY_NAME not running"
    fi
    ;;
  *)
    echo "Usage: $0 {build|start|stop|restart|status|run}"
    ;;
esac
