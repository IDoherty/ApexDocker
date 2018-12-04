package aggFuncs

import (
	"fmt"
	"log"
	"net"
)

func udpConnect(beaconAddr []Beacon, inUDPChan chan<- string, keepAlive []byte) {

	for i := 0; i < len(beaconAddr); i++ {
		RemoteAddr, err := net.ResolveUDPAddr("udp", beaconAddr[i].Address)
		if err != nil {
			log.Fatal(err)
		}

		LocalAddr, err := net.ResolveUDPAddr("udp", "192.168.187.131:0")
		if err != nil {
			log.Fatal(err)
		}

		conn, err := net.DialUDP("udp", LocalAddr, RemoteAddr)
		if err != nil {
			log.Fatal(err)
		}
		// note : you can use net.ResolveUDPAddr for LocalAddr as well

		log.Printf("Remote UDP address : %s \n", conn.RemoteAddr().String())
		log.Printf("Local UDP client address : %s \n", conn.LocalAddr().String())

		//Start Read and KA threads
		go ReadIn(conn, inUDPChan)
		go KeepAlive(conn, keepAlive)
	}
	fmt.Println()
}
