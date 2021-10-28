#pragma once

#include <common.h>
#include <matrix.h>

typedef struct {
    args_t *mult_args;
    size_t tid;
    size_t size;
    size_t cnt_threads;
} pthread_args_t;

typedef void *(*routine_t)(void *);

uint8_t start_threading(args_t *args, const size_t cnt_threads, routine_t routine, uint64_t *ticks);
