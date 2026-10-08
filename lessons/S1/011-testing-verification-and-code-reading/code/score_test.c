/* TEST HARNESS: A supplied check; keep it unchanged while implementing the task. */
gcc -std=c17 -Wall -Wextra -Wpedantic -g score_test.c -o score_test
./score_test
echo $?
