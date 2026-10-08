# Start here

Choose one route. You do not need to read the whole repository first.

## Systems/backend lessons

1. Read the [Lesson 001 lecture](lessons/001-programs-machines-and-the-terminal/README.md).
2. Follow the short [Lesson 001 work card](lessons/001-programs-machines-and-the-terminal/WORK.md).
3. Edit [exercise.c](lessons/001-programs-machines-and-the-terminal/code/exercise.c) in that folder's `code/` directory.
4. Run `python tools/check_lesson_001.py` from the repository root. It should fail on the starter and pass after the exercise prints exactly `Hello, DART!`.
5. Write your result in [Lesson 001 evidence](lessons/001-programs-machines-and-the-terminal/EVIDENCE.md), then open Lesson 002 in the [lesson index](LESSON-INDEX.md).

The lecture's Section 6 suggests creating a separate `lesson001/` directory. For this repository, use the existing `code/` directory as the work folder.

### C setup on this Windows machine

This machine has a `clang.exe` binary but no C standard headers, and no GCC. It cannot currently build C programs. Follow [SETUP.md](SETUP.md) to install a complete compiler, then run these commands in PowerShell:

```powershell
Set-Location ./lessons/001-programs-machines-and-the-terminal/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
./exercise.exe
$LASTEXITCODE
```

The starter prints `hello, Student` and exits with 0. To run its pass/fail check, return to the repository root (`Set-Location ../../..`) and run `python tools/check_lesson_001.py`.

## Semester courses

Use the semester you are enrolled in. Open a course file; its explanation and exercises are written as comments at the top of that file. Implement the exercise directly below those comments.

### S1 courses

- [Algorithms and data structures 1](lessons/S1/s1_ads1.py)
- [Analysis 1](lessons/S1/s1_analysis.py)
- [General accounting](lessons/S1/s1_accounting.py)
- [Business communication 1](lessons/S1/s1_business_communication.py)
- [Databases and multimedia](lessons/S1/s1_db_multimedia.py)
- [Digital skills](lessons/S1/s1_digital_skills.py)
- [Logic and computer architecture](lessons/S1/s1_logic_arch.py)
- [Management](lessons/S1/s1_management.py)
- [Operating systems 1](lessons/S1/s1_os1.py)
- [Statistics and probability](lessons/S1/s1_stats_prob.py)
- [S1 mastery exercises](lessons/S1/mastery_s1.py)

Run the S1 checks from the repository root: `python tools/run_semester_tests.py s1`.

### S2 courses

- [Algorithms and data structures 2](lessons/S2/s2_ads2.py)
- [Algebra](lessons/S2/s2_algebra.py)
- [Automata and language theory](lessons/S2/s2_automata.py)
- [Accounting information systems](lessons/S2/s2_accounting_is.py)
- [Business communication 2](lessons/S2/s2_business_communication.py)
- [Finance](lessons/S2/s2_finance.py)
- [Information systems and networks](lessons/S2/s2_is_networks.py)
- [Mathematical logic](lessons/S2/s2_math_logic.py)
- [Operating systems 2](lessons/S2/s2_os2.py)
- [S2 mastery exercises](lessons/S2/mastery_s2.py)

Run the S2 checks from the repository root: `python tools/run_semester_tests.py s2`.

Confirm semester topics and exam scope against your instructor's current syllabus.
