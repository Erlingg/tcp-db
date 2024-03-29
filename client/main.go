package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	//_, _ = conn.Write([]byte("SET A 111\n"))
	//_, _ = conn.Write([]byte("SET B 222\n"))
	//_, _ = conn.Write([]byte("GET A\n"))
	//_, _ = conn.Write([]byte("GET B\n"))
	//_, _ = conn.Write([]byte("KEYS\n"))
	//_, _ = conn.Write([]byte("DEL A\n"))

	//_, _ = conn.Write([]byte("привет1\n"))
	//_, _ = conn.Write([]byte("привет2\n"))
	//_, _ = conn.Write([]byte("привет3\n"))
	//_, _ = conn.Write([]byte("привет4\n"))

	r := bufio.NewReader(conn)
	for {
		data, _, err := r.ReadLine()
		if err != nil {
			if err != io.EOF {
				log.Println(err)
				return
			}
			break
		}
		fmt.Printf("read: %v \n", string(data))
	}
	time.Sleep(time.Second * 2)
	conn.Close()
}
