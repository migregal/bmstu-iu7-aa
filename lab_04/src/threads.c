#include <threads.h>

#include <pthread.h>
#include <stdlib.h>

#include <timer.h>

uint8_t start_threading(args_t *args, const size_t cnt_threads, routine_t routine, uint64_t *ticks) {
    uint8_t rc = OK;
    if ((rc = create_matrix(args->res, args->a->n, args->b->m))) return rc;

    pthread_t *threads = (pthread_t *) malloc(cnt_threads * sizeof(pthread_t));
    if (!threads) return ALLOCATION_ERROR;

    pthread_args_t *args_array = malloc(sizeof(pthread_args_t) * cnt_threads);
    if (!args_array) {
        free(threads);
        return ALLOCATION_ERROR;
    }

    for (size_t i = 0; i < cnt_threads; i++) {
        args_array[i].mult_args = args;
        args_array[i].tid = i;
        args_array[i].size = args->a->n;
        args_array[i].cnt_threads = cnt_threads;
    }

    *ticks = 0;
#ifdef __TICKS_COUNT__
    uint64_t total_ticks = 0;
    for (int c = 0; c < REPEATS_COUNT; ++c) {
        uint64_t start = tick();
#endif
        for (size_t i = 0; i < cnt_threads; i++)
            pthread_create(&threads[i], NULL, routine, &args_array[i]);

        for (size_t i = 0; i < cnt_threads; i++)
            pthread_join(threads[i], NULL);
#ifdef __TICKS_COUNT__
        uint64_t end = tick();
        total_ticks += end - start;
    }
    *ticks = total_ticks / REPEATS_COUNT;
#endif

    free(args_array);
    free(threads);

    return OK;
}
