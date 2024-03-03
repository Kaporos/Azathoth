#!/bin/bash


# Function to compile and run the server
run_server() {
    echo "Compiling and running the server..."
    go build -o /tmp/neo-mud . && /tmp/neo-mud 4000 &
    SERVER_PID=$!
}

# Initial run of the server
run_server


# Function to stop the server
stop_server() {
    echo "Stopping the server..."
    kill $SERVER_PID
    exit 0
}

# Trap Ctrl+C and stop the server
trap stop_server SIGINT

# Monitor changes to Go files and restart the server
while inotifywait --exclude '\.git' -e modify,move,create,delete -r .; do
    echo "Changes detected. Restarting the server..."
    # Kill the previous server process
    killall neo-mud
    sleep 1
    clear
    # Compile and run the server again
    run_server
done
