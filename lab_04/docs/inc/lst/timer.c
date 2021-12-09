#include <timer.h>

#include <stdio.h>
#include <x86intrin.h>

inline uint64_t tick(void) {
    return __rdtsc();
}
