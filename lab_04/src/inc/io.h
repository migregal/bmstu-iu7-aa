#pragma once

#include <inttypes.h>

#include <matrix.h>

void print_main_prompt();

uint8_t get_from_stdin(matrix_t *const a, matrix_t *const b);

uint8_t get_from_file(const char *const f, matrix_t *const a, matrix_t *const b);

void print_multiplication_results(matrix_t *res, const uint64_t ticks, int8_t printable_res);
