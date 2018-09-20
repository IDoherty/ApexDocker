package udpServer

import(
	"fmt"
)

type udpClientStruct struct{
	clientAddr	
	index		int
	// Channel Elements
	deadChan 	chan string	// Channel that passes back index if channel times out
	writeChan 	chan string // Clone of outChan containing outgoing packets
}

type UDPServer struct {
	addr   	string
	server 	*net.UDPConn
}

// Declare Indexing Variable
var clientIndex int = 0

// Declare Server Struct
var udp UDPServer

// Declare Client Address Slice
var UdpClient []udpClientStruct

// Prototype Function for Adding Clients to Address Slice	
func addUdpClient(udpClient *udpClientStruct, clientAddr /*type*/){
	// Declare empty Struct to Append
	var	tempIndex int
	var newClient udpClientStruct
	
	// Add new Struct to slice
	udpClient = append(udpClient, newClient)
	tempIndex = len(udpClient)
	
	// Fill new Struct's fields
	udpClient[tempIndex].clientAddr = clientAddr
	udpClient[tempIndex].index = clientIndex
	clientIndex++
}
	
// Prototype Function for Removing Clients from Address Slice	
func removeUdpClient(udpClient *udpClientStruct, deadIndex int){
	// Locate Desired Element
	for i:=range udpClient{
		if udpClient.index == deadIndex{
			clearIndex = i
		}
	}
	udpClient = append(udpclient[:clearIndex] + udpClient[clearIndex:]...)
}

// Main Server Function
func udpServer(){
	
	// Declare Server Address
	hostName := "192.168.187.131:8808"
	//portNum 	:= ":8808"
	udp.addr = hostName //+  portNum
	
	// Resolve Local Address
	laddr, err := net.ResolveUDPAddr("udp", udp.addr)
	if err != nil {
		log.Fatal(err)
	}
	
	
	
	// setup listener for incoming UDP connection
	udp.server, err = net.ListenUDP("udp", laddr)
	if err != nil {
		log.Fatal(err)
	}
	
	
	
	
	
	

	

	

	//go KeepAliveResponse(decodedHex)

	defer udp.server.Close()
	
	
	
}