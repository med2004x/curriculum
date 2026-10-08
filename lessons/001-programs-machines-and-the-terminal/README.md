# Lesson 001 — Programs, Machines, and the Terminal

## Threads

**Primary:** C

**Supporting:** Computing model; shell; study method

**Math companion:** number systems

**Time:** 3–4 hours. **Prerequisites:** none. This is the first lesson; there is nothing to retrieve yet.

## What you will be able to do

1. Say what a program *is* and what happens between typing code and seeing output.
2. Move around a Linux filesystem from the terminal and read a path without guessing.
3. Explain why a file's size in bytes can differ from its number of characters.
4. Compile and run a C program, and read the exit status the shell reports.

## Why this lesson exists

Every later lesson assumes you can open a terminal, find a file, build something, and know what "it ran" actually means. Learners who skip this spend months confused by problems that are really "I was in the wrong directory" or "I never saved the file". This lesson removes that whole category of failure.

## 1. A computer is a machine that changes state

A computer holds **state** (bits in memory and on disk) and follows **instructions** that change that state. Nothing else is happening. A "program" is a file of bytes, some of which the CPU interprets as instructions.

```text
source text  ->  compiler  ->  executable bytes  ->  process  ->  reads/writes state  ->  exit
  hello.c                        ./hello            running        stdout, files         status code
```

Three words that beginners blur together:

| Word | What it is | Where it lives |
|---|---|---|
| **Source code** | Text a human writes | Disk (`hello.c`) |
| **Executable** | Bytes a CPU can run | Disk (`hello`) |
| **Process** | An executable that is *currently running*, with its own memory | RAM, while alive |

When the process exits, its memory is gone. Only what it wrote to disk survives.

## 2. Bits, bytes, and text

A **bit** is 0 or 1. A **byte** is 8 bits, so it holds 256 distinct values (0–255). Text is stored as bytes through an encoding.

- **ASCII** maps 128 characters to bytes: `A` is 65, `a` is 97, `0` is 48.
- **UTF-8** is a variable-width encoding: ASCII characters take 1 byte, others take 2, 3, or 4.

A worked example, verified on a real system. The file below contains the four characters `a é € 😀`:

| Character | Unicode code point | UTF-8 bytes (hex) | Byte count |
|---|---|---|---|
| a | U+0061 | `61` | 1 |
| é | U+00E9 | `c3 a9` | 2 |
| € | U+20AC | `e2 82 ac` | 3 |
| 😀 | U+1F600 | `f0 9f 98 80` | 4 |

Four characters, **ten bytes**. `wc -m` (characters) reports 4 and `wc -c` (bytes) reports 10, *provided your locale is UTF-8*. Under the plain `C`/`POSIX` locale, `wc -m` counts bytes and wrongly reports 10 as well. Run `LC_ALL=C.UTF-8 wc -m file` if you see that. This is a real environment trap, not a typo.

**Hexadecimal** is base 16 (digits `0–9`, `a–f`). One hex digit is exactly 4 bits, so two hex digits are exactly one byte. That is the only reason it is used: it is a compact way to write bytes. `0xff` is `1111 1111` is 255.

## 3. The terminal

The shell reads a line, runs a program, shows its output, and records how it ended. Learn these ten commands and nothing more today:

```text
pwd                 print working directory (where am I?)
ls -l               list files with sizes and permissions
cd DIR              change directory       (cd .. goes up; cd alone goes home)
mkdir -p a/b        make directories, including parents
touch f             create empty file or update its timestamp
cat f               print a file
cp a b   mv a b     copy / move-or-rename
rm f                delete (no trash can, no undo)
man cmd             read the manual
echo $?             print the exit status of the last command
```

**Paths.** An *absolute* path starts at `/` and means the same thing everywhere. A *relative* path is resolved from your current directory, so it means different things in different places. `.` is the current directory and `..` is its parent.

**Exit status.** Every process ends with a number. `0` means success; anything else means failure. The shell stores it in `$?`. This is how scripts and later your own tools know whether a step worked.

```text
$ ls /nonexistent
ls: cannot access '/nonexistent': No such file or directory
$ echo $?
2
```

## 4. Your first C program

```c
#include <stdio.h>

int main(void) {
    printf("hello\n");
    return 0;
}
```

Read every token:

- `#include <stdio.h>` tells the preprocessor to paste in declarations for `printf`.
- `int main(void)` is the entry point. The operating system calls it. `int` is the exit status it hands back.
- `"hello\n"` is a string; `\n` is one newline byte (`0x0a`).
- `return 0;` ends the process with status 0.

Build and run:

```text
gcc -Wall -Wextra -std=c11 hello.c -o hello
./hello
echo $?
```

Always build with `-Wall -Wextra`. The compiler will warn about mistakes that are legal C but almost never intended. Treat a warning as a bug report.

The `./` matters. The shell searches directories listed in `PATH` for a bare name, and the current directory is deliberately not in it. `./hello` says "the file named `hello` in this directory".

## 5. Predict, then observe

Before running each item, write your predicted output on paper. Then run. A wrong prediction is the most useful thing that can happen to you today.

1. `printf("a\tb\n");`: how many bytes does the tab occupy in the output?
2. Remove `\n` from `"hello\n"`. Where does your shell prompt appear afterwards, and why?
3. Change `return 0;` to `return 7;`. What does `echo $?` print? (Verified: `7`.)
4. A file has 5 characters, all ASCII. How many bytes? Now replace one with `é`. How many now?

## 6. Build from blank

Create a directory `lesson001/` and, without copying any code above, write three programs:

1. `greet.c`: prints a one-line greeting.
2. `report.c`: prints a two-line report (a label and a value on each line).
3. `arith.c`: prints the result of a fixed arithmetic expression such as `(17 + 5) * 3`, using `%d`.

For each, state the expected output *before* running it, then compare.

## 7. Failure drill

Break each program on purpose, one defect at a time, and **classify the message before fixing it**:

| Defect | What to notice (verified with gcc 13) |
|---|---|
| Delete a `;` | gcc reports `expected ';'` on the line that *lacks* it, with the caret just past the statement. Other compilers may blame the next line, so learn to look at both. |
| Misspell `printf` as `printff` | Two messages. A compile *warning* (`implicit declaration`, often with "did you mean `printf`?") then a **link error**: `undefined reference to 'printff'`. No executable is produced. Which stage failed? |
| Delete `#include <stdio.h>` | Only *warnings* (`implicit declaration of function 'printf'`), and gcc 13 **still builds and runs it**. That is why `-Wall` matters: nothing stops you except the warning. Modern C standards forbid this, and newer compilers are turning it into a hard error. |
| Delete `return 0;` in `main` | Legal in C99 and later: `main` implicitly returns 0. Confirm with `echo $?`. This special rule applies to `main` only. |

The skill being trained is reading diagnostics, not memorising them.

## 8. Transfer

Express your `report.c` output as a sequence of state changes on paper, one line per step, showing what is in "the terminal" after each `printf`. Then change one requirement (three lines, or values aligned in columns) and update the trace *before* touching code.

## 9. Retrieval questions (answer without looking)

1. What are the four stages from source text to exit status?
2. Why is `./hello` needed and not just `hello`?
3. What does an exit status of `0` mean, and where do you read it?
4. Why can a 4-character file be 10 bytes?
5. Which is destroyed when a process exits: its memory, or a file it wrote?
6. Convert `0x2f` to binary and decimal.

## 10. Evidence to record

Create `lesson001/EVIDENCE.md` containing:

- the exact build command and the output of `echo $?` after running each program;
- one compiler message you triggered on purpose, copied verbatim, with your classification;
- your answers to the six retrieval questions, written from memory first, checked second.

## 11. Common mistakes

- Editing a file but forgetting to save, then recompiling the old version.
- Being in the wrong directory. Run `pwd` when anything "cannot be found".
- Forgetting that `rm` cannot be undone. Never run `rm -rf` on a path you have not printed first with `ls`.
- Trusting output you did not predict. If you cannot say why it printed that, you have not learned it yet.

## 12. Revisit schedule

Redo the blank-page `report.c` (no notes) after **1 day**, then **1 week**. Lesson 002 begins by retrieving the byte/text material above.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You may advance when, with no notes, you can write and build a C program from a blank file, predict its output and exit status correctly, and explain why a UTF-8 file's byte count differs from its character count.

**Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
