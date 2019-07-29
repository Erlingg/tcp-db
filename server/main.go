package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"tcp-db/server/domain"
	"tcp-db/server/handler"
	"tcp-db/server/service"
	"tcp-db/server/store"
)

func main() {
	dbStore := store.NewDBStore()
	dbService := service.NewDBService(dbStore)
	dbHandler := handler.NewDBHandlers(dbService)
	tcp, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer tcp.Close()

	fmt.Println("started ...")

	for {
		conn, err := tcp.Accept()
		if err != nil {
			log.Println(err)
		}

		go newConnect(conn, dbHandler)
	}
}

func newConnect(c net.Conn, h domain.DBHandler) {
	fmt.Println("conn:\t", c.RemoteAddr())

	r := bufio.NewReader(c)
	for {
		var res string
		data, _, err := r.ReadLine()
		if err != nil {
			if err != io.EOF {
				log.Println(err)
				return
			}
			break
		}

		fmt.Printf("read: %v \n", string(data))
		res = h.Query(string(data))
		fmt.Printf("res: %v \n", string(res))
		c.Write([]byte(res))
	}

	fmt.Println("closed:\t", c.RemoteAddr())

	c.Close()
}
