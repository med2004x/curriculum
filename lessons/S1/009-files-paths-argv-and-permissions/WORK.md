# Lesson 009 — Files, Paths, Arguments, and Permissions Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/argv_report.c`, `code/file_probe.c`, `code/line_count.c`, `code/exercise.c`.

## Problems

1. `argv_report.c`: print `argc` and each `argv[i]`, including a quoted argument.
2. `line_count.c`: count newline characters in a path from `argv[1]` and return 1 if `fopen` fails.
3. `file_probe.c`: report a successful open or the errno-based failure, then close the file.

## Run and check

```powershell
Set-Location lessons/S1/009-files-paths-argv-and-permissions/code
gcc -std=c17 -Wall -Wextra -Wpedantic argv_report.c -o argv_report.exe
.\argv_report.exe
gcc -std=c17 -Wall -Wextra -Wpedantic file_probe.c -o file_probe.exe
.\file_probe.exe
gcc -std=c17 -Wall -Wextra -Wpedantic line_count.c -o line_count.exe
.\line_count.exe
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Optional challenge

### File inspector

Build a CLI that walks user-supplied paths and reports type, size, permissions, and errors. It must distinguish nonexistent paths from inaccessible paths.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
