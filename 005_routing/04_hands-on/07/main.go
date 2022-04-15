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
		serve(conn)
	}
}

func serve(c net.Conn) {
	defer c.Close()
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		ln :=scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			fmt.Println("THIS IS THE END OF THE HEADERS")
			break
		}
	}
	body := "CHECK OUT THE RESPONSE FROM BODY PAYLOAD"
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "content-length: %d\r\n", len(body))
	fmt.Fprintf(c, "content-type: text/plain\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}