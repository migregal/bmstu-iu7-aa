#include <io.h>

#include <stdio.h>

#include <matrix_io.h>

#define ANSI_COLOR_YELLOW "\x1b[33m"
#define ANSI_COLOR_CYAN "\x1b[36m"
#define ANSI_COLOR_RESET "\x1b[0m"

inline void print_main_prompt() {
    printf(ANSI_COLOR_CYAN "\nMATRIX MULTIPLICATION\n\n" ANSI_COLOR_RESET);
}

inline uint8_t get_from_stdin(matrix_t *const a, matrix_t *const b) {
    uint8_t rc = OK;

    printf(ANSI_COLOR_YELLOW "Please, input first matrix:\n" ANSI_COLOR_RESET);
    if ((rc = read_matrix(stdin, a, 1))) return rc;

    printf(ANSI_COLOR_YELLOW "Please, input second matrix:\n" ANSI_COLOR_RESET);
    if ((rc = read_matrix(stdin, b, 1))) return rc;

    return (a->m != b->n) ? INCORRECT_MATR_SIZES : rc;
}

inline uint8_t get_from_file(const char *const f, matrix_t *const a, matrix_t *const b) {
    FILE *fin = fopen(f, "r");
    if (!fin) return FILE_OPEN_ERROR;

    uint8_t rc = OK;
    if ((rc = read_matrix(fin, a, 0))) {
        fclose(fin);
        return rc;
    }

    if ((rc = read_matrix(fin, b, 0))) {
        fclose(fin);
        return rc;
    }

    fclose(fin);
    return (a->m != b->n) ? INCORRECT_MATR_SIZES : rc;
}

inline void print_multiplication_results(matrix_t *res, const uint64_t ticks, int8_t printable_res) {
    printf(ANSI_COLOR_YELLOW "Execution time: %" PRIu64 " (cpu ticks)\n" ANSI_COLOR_RESET, ticks);

    if (printable_res) print_matrix(res);
}
