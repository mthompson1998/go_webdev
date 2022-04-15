package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
	go serve(conn)
	}
}

	func serve(c net.Conn) {
		defer c.Close()
		scanner := bufio.NewScanner(c)
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
			if ln == "" {
				fmt.Println("THIS IS THE END OF THE SOMETHING")
				break
			}
	}
	io.WriteString(c, "here we WRITE to the response")
}