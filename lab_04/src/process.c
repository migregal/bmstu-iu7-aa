#include <process.h>

#include <stdlib.h>

#include <io.h>
#include <matrix.h>
#include <multiplication.h>
#include <parallel_multiplication.h>
#include <threads.h>


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

    ticks = 0;
    if ((rc = start_threading(&args, 8, parallel_multiplication_by_rows, &ticks))) return rc;
    print_multiplication_results(c, ticks, 1);

    free_matrix(a);
    free_matrix(b);
    free_matrix(c);

    return rc;
}
