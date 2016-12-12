#include <FC37.h>
#include <DHT.h>
#include <BMP280.h>

DHT dhtSensor(6, 22);
BMP280 bmpSensor(10, 11, 12, 13);
FC37 rainSensor(A0);

void setup() {
  Serial.begin(115200);
  bmpSensor.begin();
}


/*output format:
 * bool, float, float, bool, float, bool, float
 * DHT                 BMP          RAIN
 */
void putSerial(int16_t t, uint16_t h, uint16_t p, uint16_t r){  
  Serial.print(t);
  Serial.print("&"); //separator
  
  Serial.print(h);
  Serial.print("&"); //separator
  
  Serial.print(p);
  Serial.print("&"); //separator
  
  Serial.print(r);
  Serial.print("@"); //end of line
  //Serial.println();
}

//void(* reset) (void) = 0;

void loop() {
  delay(2000);
  int16_t temp = -273; 
  int16_t hum = -1;
  uint16_t pres = 0;
  uint16_t rain = 1023;
  
  temp = (int16_t)(dhtSensor.readTemperature()*10);
  hum = (int16_t)(dhtSensor.readHumidity()*10);
  pres = (uint16_t)(bmpSensor.readPressure()/10);
  rain = (uint16_t)(rainSensor.readRain());
  
  putSerial(temp, hum, pres, rain);
}
