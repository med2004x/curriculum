CC = gcc
CFLAGS = -std=c11 -Wall -Wextra -Wpedantic
PROGRAMS = task_01 hello greet report arith exercise

ifeq ($(OS),Windows_NT)
RM = cmd /c del /q
else
RM = rm -f
endif

.PHONY: all check clean $(PROGRAMS)

# `make` builds every lesson program. Each program has its own main function.
all: $(PROGRAMS)

# Beginner-friendly targets: `make hello` creates hello.exe.
$(PROGRAMS): %: %.exe

%.exe: %.c
	$(CC) $(CFLAGS) $< -o $@

# Compile and run the exact-output checks for all six assignments.
check:
	python ../../../tools/check_lesson_001.py

clean:
	$(RM) $(addsuffix .exe,$(PROGRAMS))
