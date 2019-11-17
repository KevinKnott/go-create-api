package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer listen.Close()

	for {
		connection, err := listen.Accept()
		if err != nil {
			log.Panic(err)
		}

		fmt.Println("Connection accepted")

		io.WriteString(connection, "\nHello from new backend api\n")
		fmt.Fprintln(connection, "How is your day?")
		fmt.Fprintf(connection, "%v", "Will, I hope!")

		go handle(connection)

	}

}

func handle(connection net.Conn) {
	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer connection.Close()
	fmt.Println("Able to close out handle connection")

}
