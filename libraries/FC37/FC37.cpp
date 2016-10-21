#include "FC37.h"

FC37::FC37(int pin)
{
    _pin = pin;
    _lastValue = 0;
}

void FC37::begin(void)
{
    pinMode(_pin, INPUT_PULLUP);
    _lastValue = analogRead(_pin);
}

bool FC37::getError(void)
{
    return ((_lastValue - analogRead(_pin)) > 200);
}

int FC37::getRain(void)
{
    return analogRead(_pin);
}
