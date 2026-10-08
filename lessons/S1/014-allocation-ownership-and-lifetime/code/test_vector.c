gcc -std=c17 -Wall -Wextra -Wpedantic -fsanitize=address,undefined -g test_vector.c vector.c -o test_vector
./test_vector
echo $?
