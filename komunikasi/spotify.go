package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func pegangMusik(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Client connected: %s\n", conn.RemoteAddr().String())

	file, err := os.Open("Musik.txt")
	if err != nil {
		fmt.Println("Error opening Musik file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()

		_, err := conn.Write([]byte(data + "\n"))
		if err != nil {
			fmt.Println("Error writing to client:", conn.RemoteAddr().String(), err)
			return
		}

		time.Sleep(500 * time.Millisecond)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading Musik file:", err)
	}

	fmt.Println("Finish Him to streaming the music %s\n", conn.RemoteAddr().String())
}

func main() {
	budeg, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		return
	}
	defer budeg.Close()
	fmt.Println("Spotify-Indo streaming server started on port 9000")

	for {
		conn, err := budeg.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go pegangMusik(conn)
	}
}
