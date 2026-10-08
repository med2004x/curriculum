# Lesson 001 — Programs, Machines, and the Terminal

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

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

The folder includes a `Makefile`: it remembers the compiler options so you can focus on the program. From this folder, build and run the example with:

```powershell
make hello.exe
.\hello.exe
$LASTEXITCODE
```

The Makefile enables `-Wall -Wextra` for you. The compiler will warn about mistakes that are legal C but almost never intended. Treat a warning as a bug report. To build every program in this lesson, run `make`.

The `./` (or PowerShell's `.` prefix) matters. The shell searches directories listed in `PATH` for a bare name, and the current directory is deliberately not in it. The prefix says "run the file in this directory".

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
