#include <FC37.h>
#include <DHT.h>
#include <BMP280.h>

DHT *dhtSensor = nullptr;
BMP280 *bmpSensor = nullptr;
FC37 *rainSensor = nullptr;

bool dhtSensorExist = false;
bool bmpSensorExist = false;
bool rainSensorExist = false;

void setup() {
  Serial.begin(115200);
  delay(1000);
  while (!Serial.available()){}
  setSensorConfig();
  Serial.println("GO!");
}

/*input format:
 * "bool exist, uint8_t type, uint8_t port; bool exist, uint8_t type, uint8_t port; bool exist, uint8_t port;"
 *  DHT                                     BMP                                     RAIN
 */

void setSensorConfig() {
  while (!Serial.available()){} //for correct read.
  dhtSensorExist = uint8_t(Serial.read()); //if use bool, it doesn't work
  if(dhtSensorExist){ //read exist
    while (Serial.available() < 2){}
    //Serial.println("Read DHT");
    setDHT((uint8_t)Serial.read(), (uint8_t)Serial.read()); //read type and port
  }
  //setDHT(22, 6); //for debugging

  //Serial.println(Serial.read());
  while (!Serial.available()){}
  bmpSensorExist = uint8_t(Serial.read());
  if(bmpSensorExist){
    while (Serial.available() < 2){}
    //Serial.println("Read BMP");
    setBMP((uint8_t)Serial.read(), (uint8_t)Serial.read()); //set a Chip Select (CS) / SDA port
  }
  setBMP(28, 10); //debug

  while (Serial.available() < 2){} //End of line is "0".
  rainSensorExist = uint8_t(Serial.read());
  if(rainSensorExist){
    while (Serial.available() < 2){}
    //Serial.println("Read Rain");
    setRAIN((uint8_t)Serial.read());
  }
  //setRAIN(A0); //debug

  Serial.print(dhtSensorExist);
  Serial.print(bmpSensorExist);
  Serial.println(rainSensorExist);
}

void setDHT(uint8_t type, uint8_t pin) {
  Serial.print(pin);
  Serial.print("r");
  dhtSensor = new DHT(pin, type);
  dhtSensor->begin();
}

void setBMP(uint8_t type, uint8_t pin){ //type not use. TODO: add a BMP180 support. type: 28 - BMP280, 18 - BMP180
  bmpSensor = new BMP280(pin, 11, 12, 13);
  bmpSensor->begin();
}

void setRAIN(uint8_t pin){
  rainSensor = new FC37(pin);
  rainSensor->begin();
}

/*output format:
 * bool, float, float, bool, float, bool, float
 * DHT                 BMP          RAIN
 */
void putSerial(int16_t t, uint16_t h, uint16_t p, uint16_t r){  
  if (dhtSensorExist || bmpSensorExist){
    Serial.print(t);
  }else{
    Serial.print("#"); //sensor absent
  }

  Serial.print("&"); //separator

  if (dhtSensorExist){
    Serial.print(h);
  }else{
    Serial.print("#"); //sensor absent
  }

  Serial.print("&"); //separator
  
  if (bmpSensorExist){
    Serial.print(p);
  }else{
    Serial.print("#"); //sensor absent
  }

  Serial.print("&"); //separator

  if (rainSensorExist){
    Serial.print(r);
  }else{
    Serial.print("#"); //sensor absent
  }  

  Serial.print("@"); //end of line
  Serial.println(); //for debug
}

//void(* reset) (void) = 0;

void loop() {
  delay(2000);
  int16_t temp = -273; 
  int16_t hum = -1;
  uint16_t pres = 0;
  uint16_t rain = 1023;
  
  if (dhtSensorExist){
    temp = (int16_t)(dhtSensor->readTemperature()*10);
    if (temp == 0 && bmpSensorExist) {
      temp = (int16_t)(bmpSensor->readTemperature()*10); //read reserve sensor
    }
    hum = (int16_t)(dhtSensor->readHumidity()*10);
  }else if (bmpSensorExist){
    temp = (int16_t)(bmpSensor->readTemperature()*10);
  }
  
  if (bmpSensorExist){
    pres = (uint16_t)(bmpSensor->readPressure()/10);
  }

  if (rainSensorExist){
    rain = (uint16_t)(rainSensor->readRain());
  }
  
  putSerial(temp, hum, pres, rain);
}
