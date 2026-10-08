# Lesson 012 — Integrated Lab 1: Small C CLI

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

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

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
