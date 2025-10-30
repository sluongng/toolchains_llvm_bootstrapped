#include "greet.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// Implementation of the greeting function
const char* greet(const char* name) {
    static char greeting[256];
    snprintf(greeting, sizeof(greeting), "Hello from C, %s!", name);
    return greeting;
}

// Implementation of the add function
int add_numbers(int a, int b) {
    return a + b;
}
