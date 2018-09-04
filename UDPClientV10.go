package main

import (
        //"log"
		//"fmt"
		"time"
		//"encoding/hex"
		//"encoding/binary"		
		"pkg/aggFuncs"
)

// Start of Main Body of Code

func main(){
	
	// Build/Read Variables
	// Get Beacon Addresses & Details
	fileLocation := "NetworkInfo.txt"
	beaconAddr := aggFuncs.GetCSV(fileLocation)
	
	// Setup Communication Channels
	// Packet Counter Channel
	//countChan 	:= make(chan int)
	
	// Input Channel - From ReadIn routines  
	inUDPChan 	:= make(chan string)
	
	// Output Channel - From Validated Data
	outUDPChan	:= make(chan string)
	
	
	// Start UDP Connection Thread for each beacon connected to the system. Addresses taken from CSV file.
	// Add Activity check for each? Use Keep Alive response packets?
	aggFuncs.UdpConnect(beaconAddr, inUDPChan)

	// Start Packet Counter
	//go aggFuncs.PacketCounter(countChan)
	
	// Start Packet Processing Thread	
	go aggFuncs.ProcPackets(inUDPChan, outUDPChan)
	
	// Start UDP Distribution Thread
	//go aggFuncs.UdpTransmit(outUDPChan)
	
	// Run Indefinitely until Break. Add Monitoring for Functions and Routines?	Break Function to Kill code?
	for{
		time.Sleep(time.Second * 15)	
	}	
}