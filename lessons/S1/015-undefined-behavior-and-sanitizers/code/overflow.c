gcc -O2 -Wall -Wextra overflow.c -o overflow
gcc -O2 -fsanitize=undefined -g overflow.c -o overflow_ubsan
