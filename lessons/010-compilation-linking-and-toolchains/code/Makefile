CC ?= gcc
CFLAGS ?= -std=c17 -Wall -Wextra -Wpedantic -g

.PHONY: all exercise test clean

all: exercise

exercise: exercise_main.o exercise_math.o
	$(CC) $(CFLAGS) $^ -o $@

exercise_main.o: exercise_main.c exercise_math.h
	$(CC) $(CFLAGS) -c $< -o $@

exercise_math.o: exercise_math.c exercise_math.h
	$(CC) $(CFLAGS) -c $< -o $@

test: exercise_math.o exercise_test.o
	$(CC) $(CFLAGS) $^ -o $@
	./test

exercise_test.o: exercise_test.c exercise_math.h
	$(CC) $(CFLAGS) -c $< -o $@

clean:
	rm -f exercise test *.o
