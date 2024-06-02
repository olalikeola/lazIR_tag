#include <wiringPi.h>
#include <stdio.h>
#include <assert.h>
#include <stdint.h>

// Struct to store the readings
struct readings {
    unsigned usec, v;
};

// Function to spin in a loop while the pin value is equal to 'v'
static int read_while_eq(int pin, int v, unsigned timeout) {
    unsigned start = micros();
    while (1) {
        if (digitalRead(pin) != v)
            return micros() - start + 1;
        if ((micros() - start) >= timeout)
            return 0;
    }
}

// Function to get readings from the IR sensor
static int get_readings(int pin, struct readings *r, unsigned N) {
    unsigned val = 0;
    unsigned idx = 0;
    while (idx < N) {
        uint32_t t = read_while_eq(pin, val, 40000);
        if (!t)
            break;
        r[idx++] = (struct readings){.usec = t, .v = 1 - val};
        val = 1 - val;
    }
    return idx;
}

// Function to check if the signal is a header
int is_header(struct readings *r, unsigned n) {
    if (n < 2)
        return 0;
    if (r[0].usec < 4300 || r[0].usec > 4800)
        return 0;
    return 1;
}

// Function to convert readings to a signal value
unsigned convert(struct readings *r, unsigned n) {
    unsigned converted_result = 0;
    unsigned idx_in_output = 0;

    for (unsigned idx = 0; idx < n; idx++) {
        if (idx == 0 || idx == 1) {
            //assert(is_header(r, n));
            continue;
        }
        if ((idx % 2) == 0) {
            continue;
        }
        if ((idx > 1) && (r[idx].usec > 4000)) {
            idx_in_output = 0;
            continue;
        }
        if (r[idx].usec < 1000) {
            converted_result |= (0 << idx_in_output);
            idx_in_output++;
        }
        if (r[idx].usec > 1000) {
            converted_result |= (1 << idx_in_output);
            idx_in_output++;
        }
    }
    return converted_result;
}

// Function to print the readings
static void print_readings(struct readings *r, int n) {
    assert(n);
    printf("-------------------------------------------------------\n");
    for (int i = 0; i < n; i++) {
        printf("\t%d: %d = %d usec\n", i, r[i].v, r[i].usec);
    }
    printf("readings=%d\n", n);
    if (!is_header(r, n))
        printf("NOISE\n");
    else
        printf("convert=%x\n", convert(r, n));
}

const char *key_to_str(unsigned x) {
    switch (x) {
        case 0x8ff30:
            return "pi1";
        case 0x8af70:
            return "pi2";
        case 0x74fb0:
            return "pi3";
        case 0x73ff0:
            return "pi4";
        default:
            return NULL;
    }
}


const char * main(void) {
    if (wiringPiSetup() == -1) {
        fprintf(stderr, "WiringPi initialization failed\n");
        return "WiringPi initialization failed\n";
    }
    int in = 29; // WiringPi pin 29 corresponds to BCM GPIO pin 21
    pinMode(in, INPUT);
    pullUpDnControl(in, PUD_UP); // Set pull-up resistor

    //printf("about to start reading\n");

    // Define the maximum number of readings
    #define N 256
    static struct readings r[N];

    while (1) {
        while (digitalRead(in));

        int n = get_readings(in, r, N);
        // printf("done getting readings\n");

        unsigned x = convert(r, n);
        //printf("converted to %x\n", x);

        // Print the key based on the converted value
        const char *key = key_to_str(x); // Replace with key_to_str(x) if implemented
        if (key) { //return key;
              //printf("%s\n", key);
              return key;
	}
	//else
        //    print_readings(r, n);

        // Add a delay to debounce and avoid reading the same signal multiple times
        delay(500); // Adjust delay as needed
    }
}
