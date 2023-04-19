package main

import (
    "fmt"
    "net"
    "os"
	"time"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s host\n", os.Args[0])
        os.Exit(1)
    }

    host := os.Args[1]
    addr, err := net.ResolveIPAddr("ip", host)
    if err != nil {
        fmt.Println("Error:", err.Error())
        os.Exit(1)
    }

    conn, err := net.DialIP("ip4:icmp", nil, addr)
    if err != nil {
        fmt.Println("Error:", err.Error())
        os.Exit(1)
    }

    defer conn.Close()

    fmt.Printf("PING %s (%s):\n", addr.String(), host)
    for i := 0; i < 3; i++ {
        msg := []byte("Hello!")
        _, err := conn.Write(msg)
        if err != nil {
            fmt.Println("Error:", err.Error())
            os.Exit(1)
        }

        buf := make([]byte, 1024)
        conn.SetReadDeadline(time.Now().Add(time.Second))
        _, err = conn.Read(buf)
        if err != nil {
            fmt.Println("Error:", err.Error())
            os.Exit(1)
        }

        fmt.Printf("Received from %s: %s\n", addr.String(), string(buf))
    }
}