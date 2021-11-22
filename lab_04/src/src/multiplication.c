#include <multiplication.h>

#include <timer.h>

static inline uint8_t check_args(args_t *args) {
    return (args && args->a && args->b && args->res) ? OK : NULL_PTR_ERROR;
}

uint8_t base_multiplication(args_t *args, uint64_t *ticks) {
    uint8_t rc = check_args(args);
    if (rc) return rc;

    matrix_t *a = args->a, *b = args->b, *res = args->res;
    if ((rc = create_matrix(res, a->n, b->m))) return rc;

    *ticks = 0;
#ifdef __TICKS_COUNT__
    uint64_t total_ticks = 0;
    for (int c = 0; c < REPEATS_COUNT; ++c) {
        uint64_t start = tick();
#endif
        for (int i = 0; i < a->n; i++) {
            for (int j = 0; j < b->m; j++) {
                res->data[i][j] = 0;
                for (int k = 0; k < a->m; k++) {
                    res->data[i][j] += a->data[i][k] * b->data[k][j];
                }
            }
        }
#ifdef __TICKS_COUNT__
        uint64_t end = tick();
        total_ticks += end - start;
    }
    *ticks = total_ticks / REPEATS_COUNT;
#endif

    return rc;
}
