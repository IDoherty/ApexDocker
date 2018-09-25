package aggFuncs

import (
	//"fmt"
	"log"

	//"time"
	"encoding/binary"
	"encoding/hex"
)

func procPackets(inUDPChan <-chan string, outUDPChan chan<- string, metricChan chan<- string, keepAlive []byte) {

	// Build Variables
	var arrayLastPackets [50]LastPacket
	//valPkt:= 0;

	keepAlivePacketID := make([]byte, 2)

	keepAlivePacketID[0] = 0x55
	keepAlivePacketID[1] = 0xdd

	packetTest := binary.LittleEndian.Uint16(keepAlivePacketID)

	//fmt.Println(packetTest)
	//fmt.Println()

	for {
		returnedData := <-inUDPChan

		// Revert Data to []byte
		destringifiedData, err := hex.DecodeString(returnedData)
		if err != nil {
			log.Fatal(err)
		}

		// Filter out KA responses and test Validity for incoming packets
		if binary.LittleEndian.Uint16(destringifiedData[0:2]) == packetTest {
			//fmt.Println("Keep Alive Packet: Processing")
			//fmt.Printf("%s", hex.Dump(destringifiedData))
			//fmt.Println()

		} else {
			seqNo := destringifiedData[4:5]
			slotNo := destringifiedData[5:6]
			gpsTime := binary.BigEndian.Uint32(destringifiedData[8:12])

			//fmt.Printf("%s", hex.Dump(destringifiedData))
			//fmt.Println()

			//fmt.Println(slotNo[0])
			//fmt.Println(seqNo[0])
			//fmt.Println()
			//fmt.Println()

			valid, pktType := TestValidity(&arrayLastPackets[slotNo[0]], seqNo[0], gpsTime)

			//countChan <- pktType
			pktType++

			if valid {
				//valPkt++
				outUDPChan <- returnedData
				metricChan <- returnedData
			}

		}

	}

}
