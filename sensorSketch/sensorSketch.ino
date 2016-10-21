#include <DHT.h>
#include <BMP280.h>
#include <FC37.h>

DHT *dhtSensor = nullptr;
BMP280 *bmpSensor = nullptr;
FC37 *rainSensor = nullptr;

void setup() {
  Serial.begin(9600);
  delay(1000);
  setSensorConfig();
  Serial.println("GO!");
}

/*input format:
 * "bool exist, uint8_t type, uint8_t port; bool exist, uint8_t type, uint8_t port; bool exist, uint8_t port;"
 *  DHT                                     BMP                                     RAIN
 */

void setSensorConfig() {
    if(Serial.read()){ //read exist
      setDHT(Serial.read(), Serial.read()); //read type and port
    }
    setDHT(22, 6); //for debugging

    if(Serial.read()){
      setBMP(Serial.read(), Serial.read()); //set a Chip Select (CS) / SDA port
    }
    setBMP(280, 10); //debug

    if(Serial.read()){
      setRAIN(Serial.read());
    }
    setRAIN(A0); //debug
}

void setDHT(uint8_t type, uint8_t pin) { //type: 11 - DHT11, 22 - DHT22, 21 - DHT21, AM2301
  dhtSensor = new DHT(pin, type);
  dhtSensor->begin();
}

void setBMP(uint8_t type, uint8_t pin){ //type not use. TODO: add a BMP180 support
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
void putSerial(float t, float h, double p, float r){  
  if (!isnan(t)) {
    Serial.print(t);
  }else{
    Serial.print("#"); //error
  }
  Serial.print("&"); //separator
  
  if (!isnan(h)) {
    Serial.print(h);
  }else{
    Serial.print("#") ;
  }
  Serial.print("&");

  if (!isnan(p)) {
    Serial.print(p);
  }else{
    Serial.print("#") ;
  }
  Serial.print("&");
  
  if (!isnan(r)) {
    Serial.print(r);
  }else{
    Serial.print("#");
  }
  Serial.print("@"); //end of line
  Serial.println();
}

//void(* reset) (void) = 0;

void loop() {
  delay(2000);
  float temp = dhtSensor->readTemperature();
  float hum = dhtSensor->readHumidity();
  double pres = bmpSensor->readPressure();
  float rain = rainSensor->getRain();
  
  putSerial(temp, hum, pres, rain);
}
