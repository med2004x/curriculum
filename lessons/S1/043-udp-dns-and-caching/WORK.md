# Lesson 043 — Problem Set

**Topic:** UDP, DNS, and Caching.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Use `dig` (or `nslookup`) to query a real domain's A record (its IPv4 address) and observe the raw output, including the TTL value returned.

2. Extend it with one explicit boundary case: Run the same query again immediately and note whether the TTL you observe has decreased (most resolvers report the *remaining* TTL of a cached answer, decreasing in real time, which is itself direct evidence the answer came from cache rather than a fresh upstream query).

3. Rebuild the same behavior with one changed constraint: Use `dig +trace` (which performs and displays the full recursive resolution process, from root servers down to the authoritative answer) on a real domain, and identify, in the output, the transition from root server responses to TLD server responses to the final authoritative answer — write down, in your own words, which server in the trace was the *authorit.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/043-udp-dns-and-caching/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### UDP + DNS probe

Write a Go diagnostic tool that sends a bounded UDP request or performs DNS lookups, records timing, and handles loss or timeout explicitly.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
