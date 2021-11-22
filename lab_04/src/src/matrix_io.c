#include <matrix_io.h>

#include <stdlib.h>
#include <string.h>

#include <inttypes.h>

#define INPUT_ROWS_PROMPT "Please, input rows count: "
#define INPUT_COLS_PROMPT "Please, input cols count: "
#define INPUT_DATA_PROMPT "Please, input matrix data:\n"

inline static uint8_t input_matrix_dim(FILE *fin, const char *msg, size_t *value) {
    if (msg) printf("%s", msg);
    return (1 != fscanf(fin, "%zu", value)) ? INPUT_ERROR : OK;
}

inline static uint8_t input_matrix_data(FILE *fin, const char *msg, matrix_t *matr) {
    if (msg) printf("%s", msg);
    for (size_t i = 0; i < matr->n; ++i)
        for (size_t j = 0; j < matr->m; ++j)
            if (1 != fscanf(fin, "%" PRId64, &(matr->data[i][j])))
                return INPUT_ERROR;
    return OK;
}

uint8_t read_matrix(FILE *fin, matrix_t *matrix, int8_t prompt) {
    if (!matrix) return NULL_PTR_ERROR;

    uint8_t rc = OK;

    size_t n, m;
    if ((rc = input_matrix_dim(fin, (prompt) ? INPUT_ROWS_PROMPT : NULL, &n))) return rc;
    if ((rc = input_matrix_dim(fin, (prompt) ? INPUT_COLS_PROMPT : NULL, &m))) return rc;

    matrix_t *temp;
    INIT_MATR_PTR(temp);

    if ((rc = create_matrix(temp, n, m))) return rc;
    if ((rc = input_matrix_data(fin, (prompt) ? INPUT_DATA_PROMPT : NULL, temp))) return rc;

    free_matrix(matrix);
    memcpy(matrix, temp, sizeof(matrix_t));

    return OK;
}

void print_matrix(matrix_t *matrix) {
    for (size_t i = 0; i < matrix->n; ++i) {
        for (size_t j = 0; j < matrix->m; ++j)
            printf("%" PRId64 " ", matrix->data[i][j]);
        printf("\n");
    }
    printf("\n");
}
