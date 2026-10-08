# Lesson 009 — Files, Paths, Arguments, and Permissions

## Threads

**Primary:** C/Linux

**Supporting:** file I/O; argv; paths; permissions

**Math companion:** algebra review

**Time:** 4 hours. **Prerequisites:** Lesson 008.

## What you will be able to do

1. Parse `argc` and `argv` correctly, including shell quoting.
2. Open and read a file while handling `fopen` failure through `errno`.
3. Translate common Unix permission modes between octal and symbolic form.
4. Separate stdout, stderr, and process exit status.

## Retrieve first (no notes)

1. What does a struct group?
2. What did struct assignment do in the verified example?
3. What is the verified size of `struct B`?
4. Why can field order change struct size?

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

Self-check: how many C arguments does the quoted phrase produce?

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

Self-check: which stream receives the diagnostic?

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

Self-check: what exact character increments `lines`?

## 4. Permissions are grouped by owner, group, others

Verified mappings include 750 -> `-rwxr-x---`, 644 -> `-rw-r--r--`, 600 -> `-rw-------`, and 755 -> `-rwxr-xr-x`.

Worked shell example:

```sh
chmod 750 report.txt
ls -l report.txt
```

Self-check: which class has no permission in 750?

## 5. Root changes the meaning of a permission test

The verified permission-denied result requires a non-root user: `Permission denied (errno=13)`, exit status 1. Root can bypass ordinary permission checks, so a root run is not a valid failure demonstration.

Self-check: what changes when you switch from root to `nobody`?

## 6. stderr and exit status are different channels

A process can emit a diagnostic on stderr and still return an integer exit status. Redirecting stderr with `2>/dev/null` hides the message but leaves the verified exit status unchanged at 1.

Worked shell example:

```sh
./l9 nope.txt 2>/dev/null
printf '%s\n' "$?"
```

Self-check: which mechanism controls visibility, and which mechanism communicates process outcome?

## Predict, then observe

1. Predict `argc` for `./l9 three.txt "two words"`.
2. Predict the index containing `two words`.
3. Predict the errno for a missing `nope.txt`.
4. Predict the exit status on the missing-file branch.
5. Predict the symbolic mode for 600.
6. Predict what stderr redirection changes.
7. Predict whether root is a valid permission-denied test.

Run each prediction before checking the result.

## Build from blank

1. `argv_report.c`: print `argc` and each `argv[i]`, including a quoted argument.
2. `line_count.c`: count newline characters in a path from `argv[1]` and return 1 if `fopen` fails.
3. `file_probe.c`: report a successful open or the errno-based failure, then close the file.

## Failure drill

1. Use `argv[1]` without checking `argc`; observe the failure and identify the missing precondition.
2. Remove the `fopen` failure check; explain why the stream cannot be used safely.
3. Run the permission test as root; explain why the expected failure may disappear.
4. Redirect stderr and verify that the exit status remains 1.

## Transfer

The constraint changes: when there is no path argument, read from stdin; when one path argument exists, read that file. Keep the line-count loop single and shared.

## Retrieval questions

1. What is `argc` with no user arguments?
2. How many arguments does `"two words"` produce?
3. What errno was verified for a missing file?
4. What is the verified mode string for 750?
5. Which user must run the permission-denied test?
6. What happens to exit status when stderr is redirected?

## Evidence to record

Create `lesson009/EVIDENCE.md` with exactly:

```text
Lesson 009 evidence
===================

argv:
argc = 1 with no extra arguments
argv[0] = ./l9

./l9 three.txt "two words":
argc = 3
argv[1] = three.txt
argv[2] = two words

Missing file:
errno = 2
exit status = 1

Permission denied:
non-root
errno = 13
exit status = 1

Modes:
750 = -rwxr-x---
644 = -rw-r--r--
600 = -rw-------
755 = -rwxr-xr-x

2>/dev/null hides stderr
exit status remains 1
```

## Common mistakes

- Using `argv[1]` before checking `argc`.
- Using a failed `FILE *`.
- Testing permissions as root.
- Confusing stderr with exit status.
- Reading octal permission digits as one combined number.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can write a CLI file tool from a blank file, handle failure explicitly, and explain arguments, permissions, stderr, and exit status as separate mechanisms. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
