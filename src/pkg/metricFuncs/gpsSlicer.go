package metricFuncs

import (
	"encoding/binary"
)

func gpsSlicer(dataPacket []byte, gpsData *GPSStruct) (devID uint16) {

	gpsData.devID = binary.LittleEndian.Uint16(dataPacket[0:2])

	gpsData.gpsTime = binary.BigEndian.Uint32(dataPacket[2:6])
	gpsData.gpsDate = binary.BigEndian.Uint32(dataPacket[6:10])

	gpsData.codedLat = binary.LittleEndian.Uint32(dataPacket[10:14])
	gpsData.codedLong = binary.LittleEndian.Uint32(dataPacket[14:18])
	gpsData.codedAlt = binary.LittleEndian.Uint16(dataPacket[18:20])
	gpsData.codedSpeed = binary.LittleEndian.Uint16(dataPacket[20:22])

	gpsData.magX = binary.LittleEndian.Uint16(dataPacket[22:24])
	gpsData.magY = binary.LittleEndian.Uint16(dataPacket[24:26])
	gpsData.magZ = binary.LittleEndian.Uint16(dataPacket[26:28])

	gpsData.Flags = binary.LittleEndian.Uint16(dataPacket[28:30])

	return gpsData.devID
}
