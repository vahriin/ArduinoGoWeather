#include "FC37.h"

FC37::FC37(int8_t pin)
{
    _pin = pin;
    _lastValue = 0;
}

void FC37::begin(void)
{
    pinMode(_pin, INPUT);
    _lastValue = analogRead(_pin);
}

uint16_t FC37::readRain(void)
{
     return (uint16_t)analogRead(_pin);
}
