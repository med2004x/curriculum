# Curriculum and Assessment Design

## Two routes, one shared engineering library

This repository supports two LIG semester tracks and one longer engineering sequence.

| Route | Entry point | Scope |
|---|---|---|
| LIG S1 | [S1 folder](lessons/S1/) | Lessons 001?050 and first-semester subject exercises |
| LIG S2 | [S2 folder](lessons/S2/) | Lessons 051?100 and second-semester subject exercises |
| Full engineering sequence | [Lesson index](LESSON-INDEX.md) | Continue from S1 into S2; cross-linked to semester subjects where useful |

The numbered lessons are physically organized by semester: 001?050 are in `lessons/S1/`, and 051?100 are in `lessons/S2/`. Each semester folder also contains its LIG course exercises. Use the lesson index to continue in order.

## Engineering sequence

| Block | Lessons | Capability built |
|---|---:|---|
| Programming and machine foundations | 001?012 | C, shell, toolchain, testing, first CLI |
| Memory, data structures, algorithms | 013?020 | pointers, ownership, structures, complexity, project work |
| Architecture, OS, Go foundations | 021?032 | assembly, caches, processes, Go, concurrency |
| Go and algorithms | 033?040 | abstraction, generics, graphs, recursion, assessment |
| Networking | 041?047 | sockets, protocols, HTTP, TLS, packet reasoning |
| PostgreSQL and data services | 048?057 | schema, SQL, indexes, transactions, repository, project |
| Operating systems | 058?064 | processes, storage, synchronization, xv6 lab |
| Backend engineering | 065?076 | APIs, auth, consistency, resilience, integration, capstone |
| Security and distributed data | 077?092 | threat models, database internals, replication, Raft labs |
| Production engineering | 093?100 | capacity, observability, Rust, releases, upstreaming, capstone |

Each lesson is intended as a 3?6 hour study unit. The full engineering spine is therefore an extended pathway, not a single university semester.

## Assessment loop

Every practical assignment should expose a checkable contract. A healthy cycle is:

1. Predict a result or state an invariant before running code.
2. Complete the smallest implementation or hand calculation.
3. Run automated checks, tests, or an instructor rubric.
4. Read the failure signal and repair the smallest cause.
5. Change an input or constraint and verify transfer.
6. Save the command, result, explanation, and correction in evidence.

The CS50 material in `Documents/dart-upstreams` demonstrates specification-driven problems and executable check suites. The large systems projects use their own supplied harnesses: MIT's xv6 labs, MIT 6.5840, and CMU BusTub. Follow each course's matching release and academic-integrity rules; DART does not reproduce their test harnesses. The smaller DART tests cover focused lesson contracts.

## Readiness standard

A learner is ready to advance when they can reproduce the core method without notes, pass the relevant executable or written checks, explain one failure, and solve a changed case. Self-rating alone is not a passing signal.

For the LIG tracks, use each subject pack's exercises and the semester mastery rehearsal. The packs are study materials, not official course specifications. Compare their subject scope with the current instructor syllabus, assessment formats, and semester schedule. No official grades, credit hours, or weightings are inferred here.

## Source-backed lab assignments

- Lesson 064 uses the actual xv6 lab branch, handout, and `make grade` workflow.
- Lesson 085 uses the matching CMU BusTub project and its `check-tests` harness.
- Lesson 090 uses the MIT 6.5840 Raft lab and the versioned `raft1` test target in the local upstream checkout.

Run `python tools/validate_all.py` for repository checks. Run `python tools/run_semester_tests.py s1` or `s2` for subject-code acceptance tests; the initial failures are the unfinished exercises. See [SOURCES.md](SOURCES.md) for upstream links and [CODE-AUDIT.md](CODE-AUDIT.md) for what has and has not been verified locally.
