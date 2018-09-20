package metricFuncs

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type MetricPack0 struct {
	heartRate    byte
	speed        uint16
	maxSpeed10s  uint16
	mp           uint16
	speedIntens  uint16
	hrExert      uint16
	avgHeartrate byte
}

type MetricPack1 struct {
	sessDist   uint32
	sessDistZ5 uint32
	sessDistZ6 uint32
	sessHMLD   uint32
	sessEMD    uint32
}

type MetricPack2 struct {
	stepBal      byte
	totalNrAccel uint16
	totalNrDecel uint16
	totalTimRedZ uint16
	amp          uint16
	impacts      uint16
	dsl          uint16
}

type MetricAccel struct {
	lByteNrAccel    byte
	durationAccel   uint16
	startTimAccel   uint16
	startSpeedAccel uint16
	endSpeedAccel   uint16
	maxAccel        uint16
	distanceAccel   byte
}

type MetricDecel struct {
	lByteNrDecel    byte
	durationDecel   uint16
	startTimDecel   uint16
	startSpeedDecel uint16
	endSpeedDecel   uint16
	maxDecel        uint16
	distanceDecel   byte
}

type MetricStruct struct { // Contains All Fields Used in Various Metric Packets

	MetricPack0
	MetricPack1
	MetricPack2
	MetricAccel
	MetricDecel
}

type GPSStruct struct {
	devID      uint16
	gpsTime    uint32
	gpsDate    uint32
	codedLat   uint32
	codedLong  uint32
	codedAlt   uint16
	codedSpeed uint16
	magX       uint16
	magY       uint16
	magZ       uint16
	Flags      uint16 //GPS Lock -> 0x8000, HR -> 0x4000, Low Bat -> 0x2000, Ext Flag -> 0x1000
}

type DecodedStruct struct {
	TotalDistanceZ5 float32
	TotalDistanceZ6 float32
	TotalDistance   float32
	MaxSpeed        float32
	SpeedIntens     uint16
	AverageSpeed    float32
	TotalAccel      uint16
	TotalDecel      uint16
	Timestamp       float32
	StepBalLeft     float32
	StepBalRight    float32
	Impacts         uint16
	Dsl             uint16
}

// Temporary Struct for generation of Metrics at Fenway.
// Data sored as Name:Value `JSON Designator for Name` which outputs "Designator":Value when converted to JSON
type FenwayStruct struct {
	TotalDistance float32 `json:"Distance Travelled"`
	MaxSpeed      float32 `json:"Maximum Speed"`
}

func MetricFunc(metricChan <-chan string) {

	// Set number of Fragments in each Packet
	nrPkts := 3 //Number of individual Metric Packets in each Datagram
	pktTypes := make([]int, 3)

	// Set Arrays for Passing Data
	var packetIn = make([]byte, 80)
	var gpsIn = make([]byte, 30)
	var metricPack = make([]byte, 48)

	// Declare GPS Data Struct
	var gpsData GPSStruct

	// Declare Raw Metrics Struct
	var metricRaw MetricStruct

	// Declare Decoded Metrics Struct
	var decodedMetrics DecodedStruct

	// Declare Fenway Metrics Struct
	var fenwayMetrics FenwayStruct

	for {
		// Read in String from channel and convert to []byte
		readPacketIn := <-metricChan

		decodedHex, err := hex.DecodeString(readPacketIn)
		if err != nil {
			panic(err)
		}
		//fmt.Printf("%s", hex.Dump(decodedHex))
		//fmt.Println()

		headerSlice := decodedHex[4:]
		packetIn = headerSlice[:80]

		// Packet Identifiers
		SeqNo := packetIn[0:1]
		SlotID := packetIn[1:2]

		// Slice GPS Data - Slice all 28 bytes (+ 2 devID Bytes) & pass to function
		gpsIn = packetIn[2:32]
		devID := gpsSlicer(gpsIn, &gpsData)

		//Error Clear Statement for testing. Remove for final version
		//devID += 2

		// Slice Metric Packs
		for x := 0; x < nrPkts; x++ {
			start := (x * 16) + 32
			end := start + 17
			metricPack = packetIn[start:end]
			//fmt.Println("metric Pack ", x+1, " - ", metricPack)

			tempCnt := metricSlicer(metricPack, &metricRaw)
			pktTypes[x] = tempCnt

			//fmt.Println(tempVar)
			//fmt.Println()
		}

		metricDecoder(&metricRaw, &gpsData, &decodedMetrics)

		fenwayMetrics.MaxSpeed = decodedMetrics.MaxSpeed
		fenwayMetrics.TotalDistance = decodedMetrics.TotalDistance

		jsonOut, err := json.Marshal(fenwayMetrics)
		if err != nil {
			fmt.Println("error:", err)
		}

		//*/
		// Output tests
		fmt.Println("SeqNo 		- ", SeqNo)
		fmt.Println("SlotID 		- ", SlotID)
		fmt.Println("devID 		- ", devID)
		fmt.Println()
		//*/

		/*/
		// GPS Data
		fmt.Println("devID		- ", gpsData.devID)
		fmt.Println("gpsTime		- ", gpsData.gpsTime)
		fmt.Println("gpsDate		- ", gpsData.gpsDate)
		fmt.Println("codedLat	- ", gpsData.codedLat)
		fmt.Println("codedLong	- ", gpsData.codedLong)
		fmt.Println("codedAlt	- ", gpsData.codedAlt)
		fmt.Println("codedSpeed	- ", gpsData.codedSpeed)
		//fmt.Println("magX		- ", gpsData.magX)
		//fmt.Println("magY		- ", gpsData.magY)
		//fmt.Println("magZ		- ", gpsData.magZ)
		fmt.Println("Flags		- ", gpsData.Flags)
		fmt.Println()
		//*/

		/*/
		// Metric Pack 0
		//fmt.Println("heartRate	- ", metricRaw.heartRate)
		//fmt.Println("speed		- ", metricRaw.speed)
		fmt.Println("maxSpeed10s	- ", metricRaw.maxSpeed10s)
		//fmt.Println("mp		- ", metricRaw.mp)
		fmt.Println("speedIntens	- ", metricRaw.speedIntens)
		//fmt.Println("hrExert		- ", metricRaw.hrExert)
		//fmt.Println("avgHeartrate	- ", metricRaw.avgHeartrate)
		fmt.Println()
		//*/

		/*/
		// Metric Pack 1
		fmt.Println("sessDist	- ", metricRaw.sessDist)
		fmt.Println("sessDistZ5	- ", metricRaw.sessDistZ5)
		fmt.Println("sessDistZ6	- ", metricRaw.sessDistZ6)
		fmt.Println("sessHMLD	- ", metricRaw.sessHMLD)
		fmt.Println("sessEMD	- ", metricRaw.sessEMD)
		fmt.Println()
		//*/

		/*/
		// Metric Pack 2
		fmt.Println("stepBal 	- ", metricRaw.stepBal)
		fmt.Println("totalNrAccel 	- ", metricRaw.totalNrAccel)
		fmt.Println("totalNrDecel 	- ", metricRaw.totalNrDecel)
		//fmt.Println("totalTimRedZ 	- ", metricRaw.totalTimRedZ)
		//fmt.Println("amp 		- ", metricRaw.amp)
		fmt.Println("impacts 	- ", metricRaw.impacts)
		fmt.Println("dsl 		- ", metricRaw.dsl)
		fmt.Println()
		//*/

		/*/
		// Accel Pack
		fmt.Println("lByteNrAccel 		- ", metricRaw.lByteNrAccel)
		//fmt.Println("durationAccel 		- ", metricRaw.durationAccel)
		//fmt.Println("startTimAccel 		- ", metricRaw.startTimAccel)
		//fmt.Println("startSpeedAccel 	- ", metricRaw.startSpeedAccel)
		//fmt.Println("endSpeedAccel 	- ", metricRaw.endSpeedAccel)
		//fmt.Println("maxAccel 		- ", metricRaw.maxAccel)
		fmt.Println("distanceAccel 		- ", metricRaw.distanceAccel)
		fmt.Println()
		//*/

		/*/
		// Decel Pack
		fmt.Println("lByteNrDecel 	- ", metricRaw.lByteNrDecel)
		//fmt.Println("durationDecel 	- ", metricRaw.durationDecel)
		//fmt.Println("startTimDecel 	- ", metricRaw.startTimDecel)
		//fmt.Println("startSpeedDecel - ", metricRaw.startSpeedDecel)
		//fmt.Println("endSpeedDecel 	- ", metricRaw.endSpeedDecel)
		//fmt.Println("maxDecel	- ", metricRaw.maxDecel)
		fmt.Println("distanceDecel 	- ", metricRaw.distanceDecel)
		fmt.Println()
		//*/

		/*/
		// Decode Metrics from Packets
		fmt.Printf("Distance - %.2fm \r\n", decodedMetrics.totalDistance)
		fmt.Printf("Distance - %.2fKm \r\n", decodedMetrics.totalDistance/1000)
		fmt.Println()

		fmt.Printf("Zone 5 Distance - %.2fm \r\n", decodedMetrics.totalDistanceZ5)
		fmt.Printf("Zone 5 Distance - %.2fKm \r\n", decodedMetrics.totalDistanceZ5/1000)
		fmt.Println()

		fmt.Printf("Zone 6 Distance - %.2fm \r\n", decodedMetrics.totalDistanceZ6)
		fmt.Printf("Zone 6 Distance - %.2fKm \r\n", decodedMetrics.totalDistanceZ6/1000)
		fmt.Println()

		fmt.Printf("Max Speed (10s) - %.2fm/s \r\n", decodedMetrics.maxSpeed)
		fmt.Printf("Speed Intensity - %d \r\n", decodedMetrics.speedIntens)
		fmt.Println()

		fmt.Printf("Accelerations - %d \r\n", decodedMetrics.totalAccel)
		fmt.Printf("Decelerations - %d \r\n", decodedMetrics.totalDecel)
		fmt.Println()

		fmt.Printf("Step Balance Left - %.2f \r\n", decodedMetrics.stepBalLeft)
		fmt.Printf("Step Balance Left - %.2f \r\n", decodedMetrics.stepBalRight)
		fmt.Println()

		fmt.Printf("impacts - %d \r\n", decodedMetrics.impacts)
		fmt.Printf("dsl - %d \r\n", decodedMetrics.dsl)
		fmt.Println()
		//*/

		//*/
		fmt.Println(string(jsonOut))
		fmt.Println()
		//*/

		fmt.Println()
	}
}
