#include <wiringPi.h>
#include "gpio.h"

//int tsop_init(int input) {
//    pinMode(input, INPUT);
//    pullUpDnControl(input, PUD_UP); // Adjust based on your pull-up/down requirements
    // gpioSetPullUpDown(input, PUD_UP); // Adjust based on your pull-up/down requirements
//    return 1;
//}

int gpio_read(int pin) {
    return digitalRead(pin);
}
