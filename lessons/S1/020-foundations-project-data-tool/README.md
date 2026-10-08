# Lesson 020 — Foundations Project: In-Memory Data Tool

## Threads
**Primary:** C/Algorithms/Linux

**Supporting:** integrate data structures, files, tests, CLI

**Math companion:** problem decomposition

**Time:** 3–4 hours. **Prerequisites:** Lesson 019.

**Time:** 3–4 hours. **Prerequisites:** Lesson 019.

## What you will be able to do

1. Design and implement a small C log indexer from a written contract.
2. Parse records into an in-memory array and reject malformed input without corrupting prior records.
3. Separate parsing, searching, command-line handling, and reporting across source and header files.
4. Explain the data flow from `argv` through parsing to search results and exit status.

## Retrieve first (no notes)

1. Why does binary search require sorted input?
2. What condition must hold before reporting that a key was not found?
3. What should a parser do when a record has a missing field?
4. Why should errors go to stderr while normal results go to stdout?

## 1. Problem and model

**What problem this solves.** This capstone combines C parsing, arrays, search, file I/O, command-line arguments, modular compilation, and tests in one bounded program.

**The contract.** Read one text file named on the command line. Parse each valid line into a record, retain records in memory, support a small search command, and report malformed records or I/O failures without pretending the operation succeeded. The assigned specification requires at least two source files and a header, tests for parsing and search, and no unnecessary dynamic allocation or external libraries.

**Worked example.** Before coding, define the record format, such as `timestamp|level|message`, and decide whether empty fields are legal. Use a tiny input with two valid records and one malformed record. Write down the expected accepted-record count, search result, diagnostic stream, and exit status for each command.

Self-check: does the specification distinguish a valid empty message from a missing delimiter? If not, make the policy explicit before implementation.

## 2. Mechanism and vocabulary

**Mechanism.** Keep parsing separate from I/O. A parser should receive a line and either produce a fully initialized record or a clear failure result. The file-reading loop owns input acquisition; the CLI layer selects a command; the search routine operates only on validated records.

**Second point.** For an array with `n` records, a linear search takes at most `n` key comparisons. If records are sorted by the searched key, binary search takes logarithmically many comparisons, but sorting has a cost and is only justified if the interface needs ordered data or repeated searches.

Useful boundaries: end-of-file is not the same as a read error; a malformed line is not the same as an unreadable file; “not found” is not the same as parse failure.

**Worked example.** When parsing `2026-01-02|INFO|started`, locate the first two delimiters, validate that each required field exists, then copy the fields into bounded destinations with room for the terminating null byte. If a delimiter is absent, return failure before exposing a partially initialized record.

Self-check: identify who owns each buffer and when it becomes invalid. Do not retain pointers into a line buffer that will be overwritten on the next read.

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

Self-check: if the line is one character longer than the maximum message capacity, does the parser reject it, truncate it, or accept it? Pick and document one behavior.

## 4. Verification, failure, and evidence

**Verification.** Test valid input, an empty file, a missing file, a malformed line, a missing search key, and a successful search. Unit-test the parser and search routine independently of CLI output. Keep normal output on stdout and diagnostics on stderr; use nonzero exit status for failed invocation or unrecoverable I/O errors.

**Failure mode.** Feed a line with a missing delimiter and ensure the parser does not read beyond the buffer or append a partially initialized record. Then request a file that does not exist and check that the program reports an I/O error rather than “zero records found.”

**Completion criterion.** You can rebuild from the specification, explain module boundaries, and show that malformed input cannot silently become a valid record.

Worked example: capture the command, input file, stdout, stderr, and exit status for one successful search and one malformed-input case.

Self-check: which test would distinguish “record was rejected correctly” from “parser stopped reading the file early”?

## Predict, then observe

1. Predict the record count after a file with two valid lines and one malformed line.
2. Predict the exit status when the input file cannot be opened.
3. Predict the result of searching for a key absent from a valid nonempty file.
4. Predict what happens if the final line has no newline terminator.
5. Predict whether a line exceeding the fixed buffer is split across reads or rejected as oversized.
6. Predict what should appear on stdout versus stderr for a parse error.
7. Predict which module should own the decision that a record is malformed.
8. Write down the expected output and status before implementing the corresponding test.

Keep the expected behavior in a small input/output table; compare it with the implementation after the program exists.

## Build from blank

1. Define `LogRecord`, the line grammar, field bounds, and malformed-line policy in a short design sketch.

2. Implement a parser that accepts one valid record and rejects one malformed record without modifying the output record on failure.

3. Add file loading, an in-memory array with an explicit capacity policy, a search command, and clear stdout/stderr behavior.

4. Split the program into at least two `.c` files and a header; add automated tests for parser and search behavior.

Save the design sketch before coding. Keep memory allocation bounded and justified by the file-size/record-count policy.

## Failure drill

1. Supply a line missing its second delimiter and verify that no record is appended.

2. Supply a line at the exact maximum field length, then one byte beyond it; check that the chosen boundary policy is consistent.

3. Search an empty file and a file with records; distinguish “no match” from “failed to open.”

4. Plant a parsing or search bug, reduce it to the smallest input that demonstrates it, then explain the correction and evidence.

## Transfer

Change exactly one constraint: change the delimiter format, allow quoted messages, or replace linear search with sorted-index search. Explain which parser or ordering invariant changes and add a test that would fail under the original assumption.

Your transfer answer is complete when the changed rule is explicit and the test demonstrates the difference.

## Retrieval questions

1. What is the difference between malformed input and file I/O failure?
2. Why should a parser avoid partially updating its output on failure?
3. What is the ownership risk of retaining a pointer into a reused line buffer?
4. What should stdout and stderr each contain?
5. When is binary search valid for the records in this project?
6. Which parser and search behaviors must have automated tests?

## Evidence to record

Create `lesson020/EVIDENCE.md` with exactly this structure:

```text
# Lesson 020 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Treating EOF, malformed data, and read failure as the same condition.
- Retaining pointers into a line buffer that is reused on the next read.
- Writing into a record before all fields have passed validation.
- Silently truncating overlong fields without documenting that policy.
- Reporting a missing search key as a file or parsing error.

## Precision checks

1. State the accepted record grammar and maximum field sizes.
2. Explain the parser's success/failure contract and how it avoids partial records.
3. Give the search algorithm's cost and the precondition for binary search.
4. Explain how the CLI distinguishes missing file, malformed line, and no match.
5. Show the tests that protect parser boundaries and search results.

Design first, then implementation: the contract determines the parser and its tests.
## University revision bridge

For the LIG track, connect array search and parsing to the formal algorithm and data representation topics. The engineering project supplies a concrete case; it does not replace course-specific academic exercises.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can rebuild the log indexer from its written specification, demonstrate valid and malformed records, distinguish “not found” from I/O failure, and explain the data flow and ownership of each buffer. The project should include tests for parsing and search, a modular source layout, and a reproducible planted-bug diagnosis.
