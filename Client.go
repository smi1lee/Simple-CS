package main

import (
	"fmt"
	"net"
	"log"
	"bufio"
	"os"
	"strings"
)

func connectServer() {
	
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("link defeat")
		log.Fatal("error is",err)
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your ID：")
	name, _ := inputReader.ReadString('\n')
	trimName := strings.Trim(name, "\r\n")
	conn.Write([]byte(trimName))
	// nameInfo := make([]byte, 512)
	// _, err := conn.Write(nameInfo)
	// if err != nil {
	// 	fmt.Println("Write defeat")
	// }
	for {
		// buf := make([]byte, 512)
		// _, err := conn.Write(buf)
		input, _ := inputReader.ReadString('\n')
		trimInput := strings.Trim(input, "\r\n")
		if (trimInput != "" && trimInput != "quit") {
			_, err = conn.Write([]byte(trimName + ": " + trimInput))
		} else if (trimInput == "") {
			fmt.Println("the message cant's null！")
		} else if (trimInput == "quit") {
			_, err = conn.Write([]byte(trimName + " left"))
			return
		}
	}

}

func main() {
	connectServer()
}