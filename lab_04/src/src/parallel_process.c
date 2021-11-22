#include <parallel_process.h>

#include <stdlib.h>

#include <io.h>
#include <matrix.h>
#include <multiplication.h>
#include <parallel_multiplication.h>
#include <threads.h>

#define THREAD_COUNT 2

static routine_t routines[] = {
        parallel_multiplication_by_rows,
        parallel_multiplication_by_cols,
        NULL};

static inline uint8_t process(const char *filename, args_t *args) {
    if (!filename || !args) return NULL_PTR_ERROR;

    uint8_t rc = OK;
    if ((rc = get_from_file(filename, args->a, args->b))) return rc;

    uint64_t ticks = 0;
    if ((rc = base_multiplication(args, &ticks))) return rc;
    print_multiplication_results(args->res, ticks, 0);

    for (size_t j = 1; j <= THREAD_COUNT; j *= 2) {
        printf("%zu threads:\n", j);
        for (size_t k = 0; routines[k]; ++k) {
            if ((rc = start_threading(args, j, routines[k], &ticks))) return rc;
            print_multiplication_results(args->res, ticks, 0);
        }
    }

    return rc;
}

uint8_t process_file() {
    setbuf(stdout, NULL);
    print_main_prompt();

    matrix_t *a, *b, *c;
    INIT_MATR_PTR(a);
    INIT_MATR_PTR(b);
    INIT_MATR_PTR(c);

    args_t args = {.a = a, .b = b, .res = c};
    for (size_t i = 0, rc = 0; files[i]; ++i) {
        printf("\n%s\n", files[i]);
        if ((rc = process(files[i], &args))) return rc;
    }
    free_matrixes(3, a, b, c);

    return OK;
}
