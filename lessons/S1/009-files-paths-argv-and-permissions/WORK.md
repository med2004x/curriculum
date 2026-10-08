# Lesson 009 — Problem Set

**Topic:** Files, Paths, Arguments, and Permissions.
**Your files:** `code/argv_report.c`, `code/file_probe.c`, `code/line_count.c`, `code/exercise.c`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. `argv_report.c`: print `argc` and each `argv[i]`, including a quoted argument.
2. `line_count.c`: count newline characters in a path from `argv[1]` and return 1 if `fopen` fails.
3. `file_probe.c`: report a successful open or the errno-based failure, then close the file.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

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

## Stretch problem

### File inspector

Build a CLI that walks user-supplied paths and reports type, size, permissions, and errors. It must distinguish nonexistent paths from inaccessible paths.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
