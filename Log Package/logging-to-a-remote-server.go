package main

import (
	"log"
	"net"
)

func main() {
	// connect to a remote log server
	conn, err := net.Dial("udp", "log-server.example.com:514")
	if err != nil {
		log.Fatal("Failed to connect to a log server", err)
	}
	defer conn.Close()

	// create a logger that writes to the remote log server
	logger := log.New(conn, "", log.Ldate|log.Ltime)

	// log a message
	logger.Println("Hello to the log server.")
}

// this program demonstrates how to log to a remote server. It uses the net.Dial() function to connect to a remote log server over UDP. If the connection fails, the program logs a fatal error and exits.
// the program then creates a logger that writes to the remote log server using the conn connection. It sets the logger's flags to include the date and time.
// finally, the program uses the custom logger to log a message with logger.Println().
