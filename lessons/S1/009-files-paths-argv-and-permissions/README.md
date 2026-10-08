# Lesson 009 — Files, Paths, Arguments, and Permissions

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Arguments arrive after shell parsing

The program receives an argument count and vector. With no user arguments, the verified values are `argc=1` and `argv[0]="./l9"`. With `./l9 three.txt "two words"`, the verified values are `argc=3`, `argv[1]="three.txt"`, `argv[2]="two words"`.

Worked example:

```c
#include <stdio.h>

int main(int argc, char **argv) {
    for (int i = 0; i < argc; i++) printf("%s\n", argv[i]);
    return 0;
}
```

## 2. File opening can fail

`fopen` returns a stream or failure. Test the result before reading. The verified missing-file case produced `fopen(nope.txt): No such file or directory (errno=2)` and exit status 1.

Worked example:

```c
#include <errno.h>
#include <stdio.h>
#include <string.h>

int main(void) {
    FILE *fp = fopen("nope.txt", "r");
    if (fp == NULL) {
        fprintf(stderr, "fopen(nope.txt): %s (errno=%d)\n", strerror(errno), errno);
        return 1;
    }
    fclose(fp);
    return 0;
}
```

## 3. Reading lines is counting state

A simple line counter can count newline characters. In the verified three-line file, counting `\n` produced `3 lines`.

Worked example:

```c
#include <stdio.h>

int main(void) {
    FILE *fp = fopen("input.txt", "r");
    if (fp == NULL) return 1;
    int ch;
    int lines = 0;
    while ((ch = fgetc(fp)) != EOF) {
        if (ch == '\n') lines++;
    }
    fclose(fp);
    printf("%d lines\n", lines);
    return 0;
}
```

## 4. Permissions are grouped by owner, group, others

Verified mappings include 750 -> `-rwxr-x---`, 644 -> `-rw-r--r--`, 600 -> `-rw-------`, and 755 -> `-rwxr-xr-x`.

Worked shell example:

```sh
chmod 750 report.txt
ls -l report.txt
```

## 5. Root changes the meaning of a permission test

The verified permission-denied result requires a non-root user: `Permission denied (errno=13)`, exit status 1. Root can bypass ordinary permission checks, so a root run is not a valid failure demonstration.

## 6. stderr and exit status are different channels

A process can emit a diagnostic on stderr and still return an integer exit status. Redirecting stderr with `2>/dev/null` hides the message but leaves the verified exit status unchanged at 1.

Worked shell example:

```sh
./l9 nope.txt 2>/dev/null
printf '%s\n' "$?"
```

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
