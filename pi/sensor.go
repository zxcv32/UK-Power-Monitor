package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/influxdata/influxdb-client-go/v2"
	"go.bug.st/serial"
	"log"
	"os"
)

// Data structure to parse sensor serial response
type Response struct {
	VSensor struct {
		InputVoltage float64 `json:"inputVoltage"`
		SensorOutput float64 `json:"sensorOutput"`
		PowerStatus  string  `json:"powerStatus"`
	} `json:"vSensor"`
	HtSensor struct {
		HumidityPercent    int     `json:"humidityPercent"`
		TemperatureCelsius float64 `json:"temperatureCelsius"`
	} `json:"htSensor"`
}

/**
Read sensor serial response and push to InfluxDB
*/
func main() {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	portToUse := ""
	envPort := os.Getenv("SENSOR_SERIAL_PORT")
	if len(ports) > 1 {
		log.Print("More than one serial ports found: ", ports)
		if len(envPort) == 0 {
			log.Fatalln("Set SENSOR_SERIAL_PORT environment variable to specify which one to use")
		}
	}
	if len(envPort) > 0 {
		portToUse = envPort // Override
	} else {
		portToUse = ports[0] // assuming the only serial device connected is our sensor module
	}

	log.Print("Using port: ", portToUse)
	mode := &serial.Mode{
		BaudRate: 115200,
	}
	port, err := serial.Open(portToUse, mode)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(port)
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		var response Response
		json.Unmarshal([]byte(line), &response)
		log.Print("Data collected | ", response)
		writeDb(response)
	}
}

/**
Write the stats to InfluxDB
*/
func writeDb(data Response) {
	client := influxdb2.NewClient(os.Getenv("INFLUXDB_HOST"), os.Getenv("INFLUXDB_TOKEN"))
	writeAPI := client.WriteAPI("zxcv32", "lab")
	writeAPI.WriteRecord(fmt.Sprintf("lab,unit=celsius temperature=%f", data.HtSensor.TemperatureCelsius))
	writeAPI.WriteRecord(fmt.Sprintf("lab,unit=percentage humidity=%d", data.HtSensor.HumidityPercent))
	writeAPI.WriteRecord(fmt.Sprintf("lab,unit=voltage inputVoltage=%f,sensor=%f,status=\"%s\"", data.VSensor.InputVoltage, data.VSensor.SensorOutput,
		data.VSensor.PowerStatus))
	writeAPI.Flush()
	defer client.Close()
}
