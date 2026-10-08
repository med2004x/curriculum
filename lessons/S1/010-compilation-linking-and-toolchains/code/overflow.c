/* WORKED EXAMPLE: A deliberately unsafe compiler demonstration; do not use it as the assignment submission. */
gcc -Wall -Wextra -std=c11 -fsanitize=address,undefined -g overflow.c -o overflow
