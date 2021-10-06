#include <parallel_process.h>

#include <stdlib.h>

#include <io.h>
#include <matrix.h>
#include <multiplication.h>
#include <parallel_multiplication.h>
#include <threads.h>

#define THREAD_COUNT 16

static routine_t routines[] = {
        parallel_multiplication_by_rows,
        parallel_multiplication_by_cols,
        NULL};

static inline uint8_t process(const char *filename, matrix_t *const a, matrix_t *const b, matrix_t *c) {
    uint8_t rc = OK;
    if ((rc = get_from_file(filename, a, b))) return rc;

    args_t args = {.a = a, .b = b, .res = c};

    uint64_t ticks = 0;
    if ((rc = base_multiplication(&args, &ticks))) return rc;
    print_multiplication_results(c, ticks, 0);

    for (size_t j = 1; j <= THREAD_COUNT; j *= 2) {
        printf("%zu threads:\n", j);
        for (size_t k = 0; routines[k]; ++k) {
            ticks = 0;
            if ((rc = start_threading(&args, j, routines[k], &ticks))) break;
            print_multiplication_results(c, ticks, 0);
        }
        if (rc) break;
    }

    free_matrix(a);
    free_matrix(b);
    free_matrix(c);

    return rc;
}

uint8_t process_file() {
    setbuf(stdout, NULL);
    print_main_prompt();

    matrix_t *a, *b, *c;
    INIT_MATR_PTR(a);
    INIT_MATR_PTR(b);
    INIT_MATR_PTR(c);

    uint8_t rc = OK;
    for (size_t i = 0; files[i]; ++i) {
        printf("\n%s\n", files[i]);
        process(files[i], a, b, c);
    }

    free(c);
    free(b);
    free(a);

    return rc;
}
