#pragma once

#include <stddef.h>

#define OK 0

#define INPUT_ERROR 1
#define FILE_OPEN_ERROR 2
#define INCORRECT_MATR_SIZES 3

#define NULL_PTR_ERROR 10
#define ALLOCATION_ERROR 11
#define MEM_ALREADY_ALLOCATED 12

#define REPEATS_COUNT 100

static const char *const files[] = {
        "./data/32x32.txt",
        "./data/64x64.txt",
        "./data/128x128.txt",
        "./data/256x256.txt",
        "./data/512x512.txt",
        "./data/1024x1024.txt",
        NULL};
