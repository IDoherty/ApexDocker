package metricFuncs

//"fmt"
//"encoding/binary"
//"encoding/hex"
//"bytes"

func Truncate(some float32) float32 {
	return float32(int(some*100)) / 100
}

func metricDecoder(metricData *MetricStruct, gpsData *GPSStruct, decodedData *DecodedStruct) {

	// Declare Constant Values
	var DistConv float32 = 0.0711111
	var SpeedConv float32 = 0.00277778

	// Calculate Distance
	tempDist := metricData.sessDist
	var floatDist float32 = float32(tempDist) * DistConv
	decodedData.TotalDistance = Truncate(floatDist)

	tempDist = metricData.sessDistZ5
	floatDist = float32(tempDist) * DistConv
	decodedData.TotalDistanceZ5 = Truncate(floatDist)

	tempDist = metricData.sessDistZ6
	floatDist = float32(tempDist) * DistConv
	decodedData.TotalDistanceZ6 = Truncate(floatDist)

	// Calculate Max Speed
	tempMaxSpeed := metricData.maxSpeed10s
	var floatMaxSpeed float32 = float32(tempMaxSpeed) * SpeedConv
	decodedData.MaxSpeed = Truncate(floatMaxSpeed)
	decodedData.SpeedIntens = metricData.speedIntens

	decodedData.TotalAccel = metricData.totalNrAccel
	decodedData.TotalDecel = metricData.totalNrDecel

	// Calculate Step Balance + impacts and DSL

	decodedData.StepBalLeft = Truncate(40 + ((float32(metricData.stepBal) * 20) / 255))
	decodedData.StepBalRight = 100 - decodedData.StepBalLeft

	decodedData.Impacts = metricData.impacts
	decodedData.Dsl = metricData.dsl

	// Calculate Average Speed
	tempAvSpeed := metricData.speed
	var floatAvSpeed float32 = float32(tempAvSpeed) * SpeedConv
	decodedData.AverageSpeed = Truncate(floatAvSpeed)

	// Calculate Timestamp
	decodedData.Timestamp = float32(gpsData.gpsTime)

}
