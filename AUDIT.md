# Stage 1 — DART practice-layer audit

Audit date: 2026-10-08  
Scope: the current checkout at baseline commit `1d62764`. This is a point-in-time audit, not a claim about the untouched original ZIP. The repository had been initialized but every project file was untracked, so the existing checkout was recorded in that baseline commit before this staged work. No lecture file (`lessons/*/README.md`) was edited for this audit.

Follow-up structure change: after this snapshot, the user asked to remove the separate top-level `LIG/` directory. Current S1/S2 course exercises are now inside `lessons/S1/` and `lessons/S2/`; this audit's old paths describe the prior snapshot only.

## Findings at a glance

Several defects in the supplied audit are already fixed in the current checkout; several remain, and the current automated checks expose additional problems. In particular, skipped Go tests and comment-only duplicate challenges are no longer present, and evidence paths now resolve. However, 12 Go test functions across 7 test files are not a functioning core-path autograder: 4 lesson packages fail to compile because multiple student/reference files redeclare the same symbols, while tests in 3 packages pass on the current starter (one is a worked example). The Go tests therefore do not meet the required starter-fails/correct-solution-passes standard.

| Area | Current result |
|---|---|
| Lesson/work packet count | 100 lessons; 100 `WORK.md` files |
| Go tests | 7 files, 12 test functions, 0 `t.Skip`; 4 packages fail at compile time, 3 pass as checked |
| C checks | 5 `check.c` files plus 2 `test_*.c` files; GCC is unavailable, so they were not built |
| Challenge duplication | 16 `challenge.*` files; 0 normalize identically to the matching exercise after stripping comments/whitespace |
| Starter size | 102 `exercise.*` files; median 307 bytes |
| Evidence paths | 0 dead `code/...` paths in `EVIDENCE.md` |
| Markdown links | 34 dead relative links, mostly nested reference README quick-reference links |
| Answer keys | 0 files named `ANSWERS.md` |
| Rust | 1 `.rs` file; standalone `rustc` metadata compilation succeeds; no `Cargo.toml` or Rust test project |
| Structure validators | `python tools/validate_all.py` and `python tools/validate_lig_curriculum.py` pass |

## Verification commands and details

### Tests and tooling

- `Get-ChildItem lessons -Directory | Measure-Object` reports 100 lesson directories.
- `Get-ChildItem lessons -Recurse -Filter '*_test.go' | Measure-Object` reports 7 Go test files. A scan for `t.Skip(` / `t.Skipf(` finds 0. Those skips were removed or the old scaffolds deleted; the supplied “18 skipped test files” count does not describe this checkout.
- A static count of `func Test...` declarations finds 12 Go test functions.
- To avoid Go treating the repository root as a module, tests were run from each package directory with `GO111MODULE=off go test`. Results:
  - PASS: lesson 011's worked-example test package, lesson 031, lesson 090.
  - Compile failure: lessons 029, 030, 041, 055. Exercise, challenge, and/or demo files in each flat directory redeclare `main` and/or the function/type under test. These do not reach an assertion failure.
  - Lesson 031 passes while the starter serializes all calls behind a global mutex, so its current test rewards a correct completion property without enforcing the claimed lock-ordering behavior.
  - Lesson 090's two tests pass on the starter implementation. The starter-fails condition is therefore false for this lesson.
- `Get-Command gcc` finds no GCC executable in this environment. The five C check harnesses and two C test files were not compiled or executed. Do not treat their presence as verified feedback.
- Rust toolchain present: `rustc 1.98.1`, `cargo 1.98.1`. `rustc --edition=2021 --emit=metadata -o <temp>/exercise.rmeta lessons/097-rust-ownership-lifetimes-and-systems/code/exercise.rs` exits 0. This proves syntax/type compilation only; it does not provide tests, a Cargo project, or meaningful ownership/lifetime exercises.
- `python tools/validate_all.py` passes all four structural validators. `python tools/validate_lig_curriculum.py` reports 2 tracks, 21 subject packs, 9 Python test files, 0 errors. These outcomes check structure and the separate LIG packs; they do not establish DART exercise quality.

### Challenges, starters, work packets, and answer keys

- The current code has 16 `challenge.*` files. Applying the same simple comment/whitespace normalization used by `tools/validate_code_layer.py` finds 0 exact duplicates with their matching `exercise.*`. Thus the supplied 87/99 duplicate count is stale for this checkout. This does not prove that the remaining challenges are pedagogically harder or have their own checks; the current validator checks only textual non-identity.
- The 102 `exercise.*` files have a median size of 307 bytes. The listed specialist/capstone starters remain small: lesson 064 `exercise.c` 192 bytes; lesson 085 `exercise.c` 374 bytes and `exercise.sql` 368 bytes; lesson 090 `exercise.go` 293 bytes; lesson 076 `exercise.go` 242 bytes; lesson 100 `exercise.go` 230 bytes. These files do not implement xv6, BusTub, Raft, or the capstone service.
- All 100 work packets have distinct full-file hashes, and their topic prompts are not byte-for-byte identical. The recurring Practice A/B/C and Exercise 1–5 structure is nevertheless template-like, and the validator requires that repeated shape. There are no `ANSWERS.md` files. The self-test remains learner-scored, so retrieval questions do not have a checked key.
- The previously reported Lesson 100 `Project 15 - Architecture Review` artifact is absent from its current `WORK.md`. A scan did find malformed or doubled prompt phrasing in the current packets (for example `**Explain ...**` embedded inside a sentence beginning “explain the main mechanism in”). The current validator does not reject that pattern.

### Paths, academic scope, and validator coverage

- A scan of `EVIDENCE.md` for literal backticked `code/...` references finds 0 dead paths. The specific 58-path defect is no longer present in these evidence files.
- `README.md` and `CODE-ARTIFACT-RULES.md` do not contradict each other in the current checkout: all 39 lesson `code/reference/` directories exist. This is different from the supplied audit's earlier snapshot.
- A repository-wide scan of local Markdown links, after excluding inline-code spans and external URLs, finds 34 dead relative links. Many nested worked-example READMEs point to `GO-QUICK-REFERENCE.md` or `QUICK-REFERENCE.md` at the wrong relative depth. Existing validators do not traverse and resolve all Markdown links, which explains why their PASS result does not catch these.
- The academic gap is partly changed: the repository now contains 21 LIG semester packs, including an Automata pack, an Analysis pack, and an Algebra pack, so “zero content on automata/calculus/linear algebra” is too broad for the current checkout. The materials are introductory practice packs, not a full calculus or linear-algebra course. Regex remains absent from the scanned lesson/LIG teaching text. The current `ACADEMIC.md`/LIG mapping should be checked for scope claims during planning rather than treated as proof of full coverage.
- Current `tools/validate_course.py`, `validate_code_layer.py`, and `validate_code_artifacts.py` check numbering, required headings, placeholders, selected evidence paths, simple starter markers, and normalized exercise/challenge identity. They do not run C/Go behavioral checks, detect a passing starter, resolve every Markdown link, require answer keys, or assess challenge difficulty. Their aggregate PASS is a structural result only.

## Upstream resources inspected

The requested read-only checkouts are available as siblings under `C:/Users/med/Documents/dart-upstreams/` rather than in the pasted prompt's proposed `./_ref/` directory. They have not been copied into the learner tree.

| Checkout | Observed layout / lab material | License finding |
|---|---|---|
| `cs50-problems` | Individual problem directories such as `hello`, `mario`, `cash`, `credit`, `filter`, `recover`, `speller`, `finance`, and `fiftyville`; problem folders contain specifications/distribution material. | No `LICENSE*` or `COPYING` file exists in this checkout. Do not redistribute source/spec content absent a confirmed license. |
| `cs50-checks` | Hierarchical `check50/` and `submit50/` definitions; examples include check definitions for Mario, Caesar, Credit, and other problem sets. | No `LICENSE*` or `COPYING` file exists in this checkout. Use as a behavior/design reference and link to upstream; do not copy its checks. |
| `xv6-riscv` | RISC-V teaching OS source tree (`kernel/`, `user/`, `Makefile`, `test-xv6.py`); this checkout itself is the base xv6 source, not the course lab handout/test repository. Its only local branch is `riscv`. | Root `LICENSE` is MIT. |
| `mit-6.5840` | `src/mr`, `src/kvraft1`, `src/raft1`, `src/shardkv1`, plus `labrpc`, `tester1`, and a Makefile. These are course lab/test components, not a small isolated Raft exercise. | No root `LICENSE*` or `COPYING` file exists in this checkout. Do not redistribute its source/tests without verified permission. |
| `bustub` | C++/CMake database project with `src/`, `test/` suites for buffer, concurrency, storage, recovery, SQL, transaction, and other components, plus build tooling. README warns against publishing/forking solutions. | Root `LICENSE` is MIT; the course's solution-publication restriction still applies. |

No upstream source or solution was copied into this audit. Current DART lesson docs link to upstream work, but the xv6 lesson references a separate 2026 lab repository that is not present in the supplied local xv6 checkout. Stage 2 planning must retain that distinction and must not infer missing license terms.

## Stage 1 disposition

Audit complete. The supplied findings are not all current: defects 1, 2, and the specific evidence/reference-path claims are substantially corrected in this checkout; the remaining test and feedback requirements are not met. The active blockers are broken Go package/test isolation, starter tests that pass without the required behavior, absent answer keys, dead Markdown links, thin specialist labs, no single learner-facing check/doctor/progress path, incomplete Rust and academic scope, and upstream license gaps for three snapshots. Stop here as requested by the staged workflow; Stage 2 is the plan and mapping tables.
