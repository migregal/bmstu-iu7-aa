#include <matrix.h>

#include <stdarg.h>
#include <stdlib.h>
#include <string.h>

#include <inttypes.h>

uint8_t create_matrix(matrix_t *matrix, size_t n, size_t m) {
    if (!matrix) return NULL_PTR_ERROR;

    matrix_t *temp = (matrix_t *) malloc(sizeof(matrix_t));
    if (!temp) return ALLOCATION_ERROR;

    temp->n = temp->m = 0;
    if (!(temp->data = (int64_t **) malloc(n * sizeof(int64_t *)))) return ALLOCATION_ERROR;

    for (size_t i = 0; i < n; ++i) {
        if ((temp->data[i] = (int64_t *) malloc(m * sizeof(int64_t)))) continue;

        free_matrix(temp);
        return ALLOCATION_ERROR;
    }

    temp->n = n;
    temp->m = m;

    if (matrix->data) free_matrix(matrix);
    memcpy(matrix, temp, sizeof(matrix_t));

    return OK;
}

void free_matrix(matrix_t *matrix) {
    if (!matrix || !(matrix->data)) return;

    for (size_t i = 0; i < matrix->n; ++i)
        if (matrix->data[i])
            free(matrix->data[i]);

    free(matrix->data);

    matrix->data = NULL;
    matrix->n = matrix->m = 0;
}

void free_matrixes(size_t count, ...) {
    va_list ap;
    va_start(ap, count);
    for (size_t j = 0; j < count; j++)
        free_matrix(va_arg(ap, matrix_t *));
    va_end(ap);
}
