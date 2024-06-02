#include <wiringPi.h>
#include <pigpio.h>
#include "timer.h"

unsigned timer_get_usec_raw(void) {
    return micros();
}
