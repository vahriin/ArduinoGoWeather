#ifndef FC37_H
#define FC37_H

#if (ARDUINO >= 100)
 #include "Arduino.h"
#else
 #include "WProgram.h"
#endif

class FC37 {
private:
    int8_t _pin;
    int16_t _lastValue;

public:
    FC37(int8_t port);
    void begin(void);
    uint16_t readRain(void);
};

#endif
