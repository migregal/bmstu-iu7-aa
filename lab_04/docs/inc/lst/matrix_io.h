#pragma once

#include <matrix.h>

uint8_t read_matrix_from_file(matrix_t *matrix, FILE *fin);

uint8_t read_matrix(FILE *fin, matrix_t *matrix, int8_t prompt);

void print_matrix(matrix_t *matrix);
