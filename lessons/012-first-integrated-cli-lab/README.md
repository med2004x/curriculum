# Lesson 012 — Integrated Lab 1: Small C CLI

## Threads

**Primary:** C/Linux

**Supporting:** arguments; files; parsing; arrays; functions; tests; Git

**Math companion:** problem decomposition

**Time:** 3–4 hours. **Prerequisites:** Lesson 011.

## What you will be able to do

1. Turn a command-line contract into a multi-file C program.
2. Parse a delimited record without confusing malformed input with valid data.
3. Keep a bounded set of records in memory and search it by message text.
4. Report usage and data errors through stderr and nonzero exit status.
5. Test parsing and search independently from file I/O.

## Retrieve first (no notes)

1. What does `argc` count, and what is the value of `argv[0]`?
2. Why should a missing-file error go to stderr rather than stdout?
3. What stage reports an unresolved external function after compilation?
4. What should a test for a boundary do before the implementation is changed?

Write from memory: Lesson 009 covered argv and stderr, Lesson 010 compilation versus linking, and Lesson 011 boundary tests.

## 1. Project contract: a small log indexer

Build `log-indexer FILE TERM`. Read a text file, parse each valid record, store at most 1,000 records, then print records whose message contains `TERM`. Matching is case-sensitive. The program uses a fixed-size array; it does not need dynamic allocation for this bounded exercise.

Each input line has exactly this shape:

```text
1712345678|WARN|disk near limit
```

The fields are a nonnegative decimal Unix timestamp, one of `INFO`, `WARN`, or `ERROR`, and a nonempty message. The first two `|` characters separate fields; later `|` characters belong to the message. A malformed line is diagnosed with its line number, skipped, and makes the final exit status nonzero. A valid file with no matches exits zero and prints no records.

## 2. Trace one record before coding

For `1712345678|WARN|disk near limit`, the first separator ends the timestamp, the second ends the level, and the remaining bytes are the message. Parsing yields timestamp `1712345678`, level `WARN`, and message `disk near limit`. Searching for `disk` matches; searching for `Disk` does not.

Before implementation, write these expected observations:

| Input condition | Observable result |
|---|---|
| valid record containing `disk` | one output record, status 0 |
| valid record without `disk` | no output, status 0 |
| malformed timestamp | line-number diagnostic on stderr, nonzero status |
| missing file | operating-system error on stderr, nonzero status |
| wrong number of arguments | usage on stderr, nonzero status |

## 3. Separate the program into files

Use `log.h` for the shared type and function declarations, `log.c` for parsing and matching, `main.c` for arguments and file I/O, and `test_log.c` for unit tests. This boundary lets the tests exercise parsing without creating a file or launching the CLI.

Start with the shared interface:

```c
#ifndef LOG_H
#define LOG_H

typedef enum { INFO, WARN, ERROR } Level;

typedef struct {
    long timestamp;
    Level level;
    char message[256];
} Record;

int parse_record(char *line, Record *out);
int record_matches(const Record *record, const char *term);
const char *level_name(Level level);

#endif
```

`parse_record` returns 1 for a valid record and 0 otherwise. It may modify its input buffer while replacing delimiters, so callers pass a writable line buffer. The output is written only after all fields validate; on failure, callers must ignore it.

## 4. Implement and test the parser

In `log.c`, split at the first two `|` characters and use `strtol` for the timestamp. Reject empty, negative, overflowed, or trailing timestamp text; unknown levels; empty messages; and messages that do not fit. Trim CR/LF first. `record_matches` uses case-sensitive `strstr`.

Test the parser without file I/O. Use a new writable array for each call because parsing changes the separators:

```c
char valid[] = "1712345678|WARN|disk near limit\n";
Record record;
assert(parse_record(valid, &record));
assert(record.timestamp == 1712345678L);
assert(record.level == WARN);
assert(record_matches(&record, "disk"));
assert(!record_matches(&record, "Disk"));

char bad_level[] = "1712345678|BOGUS|disk near limit";
assert(!parse_record(bad_level, &record));
```

Add separate tests for missing separators, empty message, nonnumeric timestamp, and a 256-byte message. Each test should isolate one invalid condition so its failure points to one rule.

## 5. Build the CLI around explicit failure paths

In `main.c`, require exactly three arguments: program name, path, and search term. Open `argv[1]` with `fopen`; use `perror(argv[1])` if it fails. Read with `fgets` into a fixed line buffer, increment a line counter, and reject an overlong line rather than parsing its first fragment as a complete record.

Keep parsed records in `Record records[1000]` and track the number stored with `size_t count`. If capacity is reached, report it and return failure. For each malformed line, print `path:line: invalid record` to stderr, remember an error, and continue. Check `ferror` after reading and before closing.

For each stored record that matches, print timestamp, level name, and message to stdout. Return 1 if any input/read/capacity error occurred, otherwise return 0. Do not print a usage message or parse error to stdout: stdout is the machine-readable result stream.

## 6. Compile and observe

Build the CLI and the unit tests as separate executables:

```sh
gcc -std=c17 -Wall -Wextra -Wpedantic -g main.c log.c -o log-indexer
gcc -std=c17 -Wall -Wextra -Wpedantic -g test_log.c log.c -o test-log
./test-log
```

Create `sample.log` with the three lines below, then run the search:

```text
1712345678|INFO|started service
1712345680|WARN|disk near limit
1712345688|ERROR|disk full
```

```sh
./log-indexer sample.log disk
echo $?
```

Expected output contains the WARN and ERROR records in file order; status is 0. Searching for `Disk` prints nothing and still exits 0 because no matches is not an input failure.

## Predict, then observe

1. Predict which fields the parser extracts from the first sample record.
2. Predict whether `disk` and `Disk` both match.
3. Predict the output stream and exit status for a missing file.
4. Predict whether a malformed middle line prevents later valid records from being searched.
5. Predict what happens when a file contains 1,001 valid records.
6. Predict which unit test detects accepting `BOGUS` as a level.

Record each prediction before running. Compare stdout, stderr, and status separately; they are distinct observations.

## Build from blank

1. Write the record type, enum, and prototypes in `log.h`.
2. Implement and test `parse_record` before opening any file.
3. Implement `record_matches` and verify case sensitivity.
4. Write the CLI argument check, file loop, fixed-capacity array, and output path.
5. Compile the CLI and tests with the exact commands above; save their results.

## Failure drill

1. Treat `BOGUS` as `INFO`; confirm the invalid-level test fails.
2. Remove the second separator from one input; confirm parsing rejects it.
3. Search for an absent term; confirm status is still zero.
4. Rename the input file; confirm the filename diagnostic appears on stderr and the process exits nonzero.
5. Make a line longer than the input buffer; confirm it is rejected as one malformed line, not read as two records.

For each experiment, identify whether the defect is in parsing, file I/O, output routing, or command-line validation. Restore the source and rerun all tests.

## Transfer

Change the record format to `LEVEL|timestamp|message` while preserving all other behavior. Update the parser tests first, including one valid record and one invalid level. Only then change the implementation. Your evidence must show the same search results on equivalent data and a nonzero status for malformed input.

## Retrieval questions

1. What are the four valid fields or constraints in one log record?
2. Why does the parser receive a writable line buffer?
3. Which channel receives parse errors, and what process status follows any malformed line?
4. Why does a search with no matches still return status zero?
5. Which source file owns file I/O, and which one owns record parsing?
6. What should happen when the record capacity of 1,000 is exceeded?

## Evidence to record

Create `lesson012/EVIDENCE.md` containing:

```text
Design sketch:

Build commands:

Unit-test result:

Search input and output:

Malformed-line diagnostic and exit status:

One bug reproduced and repaired:

Transfer format test:
```

## Common mistakes

- Searching a raw line before validating its timestamp and level.
- Treating an unknown level as a default instead of rejecting the record.
- Losing track of the line number after skipping malformed input.
- Forgetting that `fgets` may return a fragment of an overlong line.
- Returning failure for a valid search that simply has no matches.
- Mixing diagnostics with records on stdout.
- Compiling `main.c` without `log.c` and mistaking the linker error for a parser bug.

## University revision bridge

Represent parsing as a partial function from strings to records: valid strings map to one `Record`; malformed strings have no result. The tests exhibit selected valid and invalid inputs, but they do not prove correctness for all strings. Relate each test to one precondition in the file format.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the record contract, rebuild the parser and CLI from blank across the header and two source files, test malformed data and search behavior, and distinguish stdout, stderr, and exit status. Complete the transfer format without changing unrelated behavior, then reproduce the build from a clean directory.
