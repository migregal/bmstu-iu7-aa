#pragma once

#include <stddef.h>
#include <stdint.h>

#include <stdio.h>

#include <common.h>

typedef struct {
    int64_t **data;
    size_t n, m;
} matrix_t;

#define INIT_MATR_PTR(__ptr__)                                                           \
    do {                                                                                 \
        if (!(__ptr__ = (matrix_t *) malloc(sizeof(matrix_t)))) return ALLOCATION_ERROR; \
        __ptr__->data = NULL;                                                            \
        __ptr__->n = __ptr__->m = 0;                                                     \
    } while (0)

typedef struct {
    matrix_t *a;
    matrix_t *b;
    matrix_t *res;
} args_t;

uint8_t create_matrix(matrix_t *matrix, size_t n, size_t m);

void free_matrix(matrix_t *matrix);

void free_matrixes(size_t count, ...);
