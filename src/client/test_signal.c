#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include <unistd.h>
#include <test_signal.h>

// Function to return a random string from the set {"pi1", "pi2", "pi3", "pi4"}
char* get_random_string() {
    const char* strings[] = {"pi1", "pi2", "pi3", "pi4"};
    int index = rand() % 4;  // Random index between 0 and 3
    return (char*)strings[index];
}

int loop_function(char* out) {
    srand(time(NULL));  // Seed the random number generator
    int wait_time = rand() % 10;  // Random wait time between 0 and 9 seconds

    time_t start_time = time(NULL);
    while (difftime(time(NULL), start_time) < 15) {
        if (difftime(time(NULL), start_time) == wait_time) {
            strcpy(out, get_random_string());
            return strlen(out);
        }
        sleep(1);  // Sleep for 1 second
    }
    strcpy(out, "15 seconds have passed!");
    return strlen(out);
}

// int main() {
//     char *out = malloc(30);
//     loop_function(out);
//     printf("%s\n", out);
//     free(out);
//     return 0;
// }
