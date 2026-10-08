# Lesson 085 — Database Internals Lab Problems

Read [README.md](README.md). Complete the two short local exercises, then use BusTub for the full database project.

## Problems

1. In `code/exercise.c`, implement `fetch`. Return the existing frame on a hit. On a miss, use the first free frame; when all four frames are occupied, replace frames in round-robin order. Update `hits` and `misses`. The program should report two accesses to page 7 as one miss followed by one hit.
2. In `code/exercise.sql`, add a `CHECK` constraint that rejects negative `hit_count`, then select rows in descending `hit_count` order. Try a negative insert to confirm the constraint works.
3. In the BusTub checkout, complete the buffer-pool manager project and its tests, then the B+ tree project and its tests. Use the handout for the same semester as the checkout.

## Run and check

Build `exercise.c` with your C compiler and run it. Run `exercise.sql` in a disposable PostgreSQL session. For BusTub, follow the handout's build steps and test target (often `make check-tests` from the build directory).

## Course lab

Use the official [BusTub repository](https://github.com/cmu-db/bustub) and matching [CMU 15-445 project guide](https://15445.courses.cs.cmu.edu/). Record the semester, commit, test command, and results in [EVIDENCE.md](EVIDENCE.md). Keep your solution private and follow the course's academic-integrity rules.

## Submit

Record the local outputs and BusTub test results. Explain the invariant that prevents an incorrect page from being returned for a requested page ID.
