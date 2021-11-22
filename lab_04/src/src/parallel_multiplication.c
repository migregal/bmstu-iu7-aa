#include <parallel_multiplication.h>

#include <inttypes.h>

#include <threads.h>
#include <timer.h>

void *parallel_multiplication_by_rows(void *args) {
    pthread_args_t *argsp = (pthread_args_t *) args;

    int row_start = argsp->tid * (argsp->size / argsp->cnt_threads);
    int row_end = (argsp->tid + 1) * (argsp->size / argsp->cnt_threads);

    matrix_t *a = argsp->mult_args->a, *b = argsp->mult_args->b, *res = argsp->mult_args->res;
    for (int i = row_start; i < row_end; i++) {
        for (int j = 0; j < b->m; j++) {
            res->data[i][j] = 0;
            for (int k = 0; k < a->m; k++) {
                res->data[i][j] += a->data[i][k] * b->data[k][j];
            }
        }
    }

    return NULL;
}

void *parallel_multiplication_by_cols(void *args) {
    pthread_args_t *argsp = (pthread_args_t *) args;

    int col_start = argsp->tid * (argsp->size / argsp->cnt_threads);
    int col_end = (argsp->tid + 1) * (argsp->size / argsp->cnt_threads);

    matrix_t *a = argsp->mult_args->a, *b = argsp->mult_args->b, *res = argsp->mult_args->res;
    for (int i = 0; i < a->n; i++) {
        for (int j = col_start; j < col_end; j++) {
            res->data[i][j] = 0;
            for (int k = 0; k < a->m; k++) {
                res->data[i][j] += a->data[i][k] * b->data[k][j];
            }
        }
    }

    return NULL;
}
