/* WORKED EXAMPLE: A deliberately invalid pointer example for diagnosis; do not submit it. */
gcc -std=c17 -Wall -Wextra -Wpedantic -fsanitize=address -fno-omit-frame-pointer -g bad.c -o bad
ASAN_OPTIONS=detect_stack_use_after_return=1 ./bad
