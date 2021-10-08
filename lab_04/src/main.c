#include <parallel_process.h>
#include <process.h>

int main(void) {
#ifdef __FILE_READING__
    return process_file();
#else
    return process_stdin();
#endif
}
