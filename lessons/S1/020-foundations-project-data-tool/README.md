# Lesson 020 — Foundations Project: In-Memory Data Tool

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** This capstone combines C parsing, arrays, search, file I/O, command-line arguments, modular compilation, and tests in one bounded program.

**The contract.** Read one text file named on the command line. Parse each valid line into a record, retain records in memory, support a small search command, and report malformed records or I/O failures without pretending the operation succeeded. The assigned specification requires at least two source files and a header, tests for parsing and search, and no unnecessary dynamic allocation or external libraries.

**Worked example.** Before coding, define the record format, such as `timestamp|level|message`, and decide whether empty fields are legal. Use a tiny input with two valid records and one malformed record. Write down the expected accepted-record count, search result, diagnostic stream, and exit status for each command.

## 2. Mechanism and vocabulary

**Mechanism.** Keep parsing separate from I/O. A parser should receive a line and either produce a fully initialized record or a clear failure result. The file-reading loop owns input acquisition; the CLI layer selects a command; the search routine operates only on validated records.

**Second point.** For an array with `n` records, a linear search takes at most `n` key comparisons. If records are sorted by the searched key, binary search takes logarithmically many comparisons, but sorting has a cost and is only justified if the interface needs ordered data or repeated searches.

Useful boundaries: end-of-file is not the same as a read error; a malformed line is not the same as an unreadable file; “not found” is not the same as parse failure.

**Worked example.** When parsing `2026-01-02|INFO|started`, locate the first two delimiters, validate that each required field exists, then copy the fields into bounded destinations with room for the terminating null byte. If a delimiter is absent, return failure before exposing a partially initialized record.

## 3. Worked case from the reference material

Use the assignment's log-indexer contract rather than borrowing an unrelated dependency-planner example. A suitable first record type is:

```c
#define LEVEL_MAX 8
#define MESSAGE_MAX 160

typedef struct {
    char timestamp[20];
    char level[LEVEL_MAX];
    char message[MESSAGE_MAX];
} LogRecord;
```

This is a starting representation, not a complete specification: confirm timestamp syntax and maximum field lengths before treating it as the final contract. Keep parsing in a function that returns success/failure and writes the output record only on success.

Worked example check: identify the input line, the parsed fields, and the exact condition that rejects malformed input.

## 4. Verification, failure, and evidence

**Verification.** Test valid input, an empty file, a missing file, a malformed line, a missing search key, and a successful search. Unit-test the parser and search routine independently of CLI output. Keep normal output on stdout and diagnostics on stderr; use nonzero exit status for failed invocation or unrecoverable I/O errors.

**Failure mode.** Feed a line with a missing delimiter and ensure the parser does not read beyond the buffer or append a partially initialized record. Then request a file that does not exist and check that the program reports an I/O error rather than “zero records found.”

**Completion criterion.** You can rebuild from the specification, explain module boundaries, and show that malformed input cannot silently become a valid record.

Worked example: capture the command, input file, stdout, stderr, and exit status for one successful search and one malformed-input case.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
