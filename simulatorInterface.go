package main

// export GOPATH=$(pwd)/
// export GOPATH=~/

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	. "simulatorDef"
	"strconv"
)

func listenForUserInput(userInput chan string) {
	for {
		var command int
		fmt.Scanf("%c", &command)
		if (command >= 97) && (command <= 122) {
			fmt.Println("Sending: ", string(command))
			userInput <- string(command)
		} else if command != 10 {
			fmt.Println("Rejected: ", string(command))
		}
	}
}

func UDPTransmitServer(lconn *net.UDPConn, sendChannel chan string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("ERROR in UDPTransmitServer.\nClosing connection.")
			lconn.Close()
			os.Exit(1)
		}
	}()
	for {
		msg := <-sendChannel
		networkPack, err := json.Marshal(msg)
		if err != nil {
			log.Println(err)
		}
		_, err = lconn.Write(networkPack)
		if err != nil {
			fmt.Printf("UDPTransmitServer-Simulator:\tError: Sending\n")
			panic(err)
		}
	}
}

func main() {
	fmt.Println("Starting Simulator interface")

	//Generating recive adress
	raddr, err := net.ResolveUDPAddr("udp4", "localhost:"+strconv.Itoa(PortFromInterface))
	if err != nil {
		fmt.Println("Can not resolve this adress")
		log.Println(err)
		panic(err)
	} else {
		fmt.Printf("Sending to:\t %s\n", raddr)
	}

	//Creating local connection
	localTransmitConn, err := net.DialUDP("udp4", nil, raddr)
	if err != nil {
		fmt.Println("Can not create UDP soccet on this port")
		log.Println(err)
		panic(err)
	} else {
		fmt.Println("From:\t\t", localTransmitConn.LocalAddr().String())
	}

	//Making channels
	sendChannel := make(chan string, 10)
	userInput := make(chan string)

	//Spawning threads
	go UDPTransmitServer(localTransmitConn, sendChannel)
	go listenForUserInput(userInput)

	for {
		select {
		case char := <-userInput:
			sendChannel <- char
		}
	}
}
