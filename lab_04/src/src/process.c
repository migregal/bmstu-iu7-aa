#include <process.h>

#include <stdlib.h>

#include <io.h>
#include <matrix.h>
#include <multiplication.h>
#include <parallel_multiplication.h>
#include <threads.h>

#define THREAD_COUNT 1

uint8_t process_stdin() {
    setbuf(stdout, NULL);

    print_main_prompt();

    matrix_t *a, *b;
    INIT_MATR_PTR(a);
    INIT_MATR_PTR(b);

    uint8_t rc = get_from_stdin(a, b);
    if (rc) return rc;

    matrix_t *c;
    INIT_MATR_PTR(c);
    args_t args = {.a = a, .b = b, .res = c};

    uint64_t ticks = 0;
    if ((rc = base_multiplication(&args, &ticks))) return rc;
    print_multiplication_results(c, ticks, 1);

    if ((rc = start_threading(&args, THREAD_COUNT, parallel_multiplication_by_rows, &ticks))) return rc;
    print_multiplication_results(c, ticks, 1);

    if ((rc = start_threading(&args, THREAD_COUNT, parallel_multiplication_by_cols, &ticks))) return rc;
    print_multiplication_results(c, ticks, 1);

    free_matrixes(3, a, b, c);

    return rc;
}
