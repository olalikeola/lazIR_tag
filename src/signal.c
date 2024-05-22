#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <unistd.h>

// Function to return a random string from the set {"pi1", "pi2", "pi3", "pi4"}
char* get_random_string() {
    const char* strings[] = {"pi1", "pi2", "pi3", "pi4"};
    int index = rand() % 4;  // Random index between 0 and 3
    return (char*)strings[index];
}

char* loop_function() {
    srand(time(NULL));  // Seed the random number generator
    int wait_time = rand() % 10;  // Random wait time between 0 and 9 seconds

    time_t start_time = time(NULL);
    while (difftime(time(NULL), start_time) < 15) {
        if (difftime(time(NULL), start_time) == wait_time) {
            return get_random_string();
        }
        sleep(1);  // Sleep for 1 second
    }
    return "15 seconds have passed!";
}

int main() {
    char *result = loop_function();
    printf("%s\n", result);
    return 0;
}
