/* WORKED EXAMPLE: A deliberately undefined operation for sanitizer observation; do not use it as a solution. */
gcc -O2 -Wall -Wextra overflow.c -o overflow
gcc -O2 -fsanitize=undefined -g overflow.c -o overflow_ubsan
