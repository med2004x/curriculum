# Lesson 043 — UDP, DNS, and Caching Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Use `dig` (or `nslookup`) to query a real domain's A record (its IPv4 address) and observe the raw output, including the TTL value returned.

2. Run the same query again immediately and note whether the TTL you observe has decreased (most resolvers report the *remaining* TTL of a cached answer, decreasing in real time, which is itself direct evidence the answer came from cache rather than a fresh upstream query).

3. Use `dig +trace` (which performs and displays the full recursive resolution process, from root servers down to the authoritative answer) on a real domain, and identify, in the output, the transition from root server responses to TLD server responses to the final authoritative answer — write down, in your own words, which server in the trace was the *authorit.

## Run and check

```powershell
Set-Location lessons/S1/043-udp-dns-and-caching/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### UDP + DNS probe

Write a Go diagnostic tool that sends a bounded UDP request or performs DNS lookups, records timing, and handles loss or timeout explicitly.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
