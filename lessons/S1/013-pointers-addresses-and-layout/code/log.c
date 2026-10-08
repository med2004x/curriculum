/* WORKED EXAMPLE: A complete parser example for study; implement the assignment in WORK.md. */
gcc -std=c17 -Wall -Wextra -Wpedantic -g main.c log.c -o log-indexer
gcc -std=c17 -Wall -Wextra -Wpedantic -g test_log.c log.c -o test-log
./test-log
