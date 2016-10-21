#ifndef FC37_H
#define FC37_H

#if (ARDUINO >= 100)
 #include "Arduino.h"
#else
 #include "WProgram.h"
#endif

class FC37 {
private:
    int _pin;
    int _lastValue;

public:
    FC37(int port);
    void begin(void);
    int getRain(void);
    bool getError(void);
};

#endif
