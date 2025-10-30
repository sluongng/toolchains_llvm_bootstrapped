#ifndef GREET_H
#define GREET_H

#ifdef __cplusplus
extern "C" {
#endif

// Simple greeting function that can be called from Go via cgo
const char* greet(const char* name);

// Math function to demonstrate numeric operations
int add_numbers(int a, int b);

#ifdef __cplusplus
}
#endif

#endif // GREET_H
