#include "DHT.h"
#include <ArduinoJson.h>

// DHT11
#define dhtSensorIn 3
#define dhtType DHT11
DHT dht(dhtSensorIn, dhtType);

// Voltage Sensor
const int voltageSensorInput = A0;
const int voltageSensorOutput = 2;

float vOut = 0.0;
float vIn = 0.0;
float R1 = 30000.0;
float R2 = 7500.0;
int vSensor = 0;


void setup()
{
  Serial.begin(9600);  
  pinMode(LED_BUILTIN, OUTPUT);
  pinMode(voltageSensorOutput, OUTPUT);
  dht.begin();
}

void loop()
{
  digitalWrite(LED_BUILTIN, HIGH);
  
  // Voltage Sensor
  vSensor = analogRead(voltageSensorInput);
  vOut = (vSensor* 5.0) / 1024.0;
  vIn = vOut / (R2/(R1+R2));
  StaticJsonDocument<48> vSensor;
  vSensor["inputVoltage"] = vIn;
  vSensor["sensorOutput"] = vOut;
  if (vIn >= 5){
    digitalWrite(voltageSensorOutput, HIGH); 
    vSensor["powerStatus"] = "live";
  } else {
    digitalWrite(voltageSensorOutput, LOW); 
    vSensor["powerStatus"] = "down";
  }

  // Temperature and Humidity Sensor
  StaticJsonDocument<16> htSensor;
  float humidity = dht.readHumidity(); 
  float temperature = dht.readTemperature(); 
  if (isnan(humidity) || isnan(temperature)) {
    htSensor["error"] = "Failed to read from DHT sensor!";  
  } else{
    htSensor["humidityPercent"] = humidity;
    htSensor["temperatureCelsius"] = temperature;  
  }
  StaticJsonDocument<96> output;
  output["vSensor"] = vSensor;
  output["htSensor"] = htSensor;
  serializeJson(output, Serial);
  Serial.println();
  digitalWrite(LED_BUILTIN, LOW); 
  delay(50);
}
