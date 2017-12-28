package main

import (
	"fmt"
	"os"
	"log"
	"time"
	"net"
	// "strings"
)

func startServer() {

	fileName := (time.Now().Format("20060102") + "_Debug.log")
	logFile, err := os.Create(fileName)
	logger := log.New(logFile,"[Debug]",log.Llongfile)
	defer logFile.Close()

	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("link defeat", err.Error())
		logger.Println("link defeat！")
		return
	}

	defer listener.Close()

	fmt.Println("The ChatRoom build success!")

	for {
		conn, err := listener.Accept()
		defer conn.Close()
		if err != nil {
			logger.Println("accept defeat")
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	nameInfo := make([]byte, 512)
	_, err := conn.Read(nameInfo)

	// fmt.Println(string(nameInfo))
	fmt.Printf("%sjoin in\n", nameInfo)

	if err != nil {
		fmt.Println("read defeat,43h")
	}

	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println(string(nameInfo) + "left")
			break
		}
		fmt.Printf("%s\n", buf)
		// fmt.Println(string(buf))
	}
}



func main() {
	
	startServer()

}