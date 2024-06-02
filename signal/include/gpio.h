#ifndef GPIO_H
#define GPIO_H

//int tsop_init(int input);
int gpio_read(int pin);
void gpio_set_pullup(int pin);

#endif // GPIO_H
