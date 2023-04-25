package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Welcome to Mobile Music, your home for quick and easy network-queued music!")

	fmt.Print("Hosted at ")

	fmt.Print(GetOutboundIP())
	fmt.Println(":3621")
	fmt.Println()
	fmt.Println("Starting Server...")

	initialize()

	fmt.Println("Done.")

}

func initialize() {

}

// Gets the device's local address, which is returned to the user.
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
