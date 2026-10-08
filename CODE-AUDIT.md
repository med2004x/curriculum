# Code and assessment audit

Updated 2026-10-08. Structural validator PASS means repository shape and literal paths pass; it is not a rating of assignment quality or a claim that unfinished student work passes.

## Engineering route

- 100 numbered lessons have code workspaces.
- The false skipped Go test scaffolds were removed; no `t.Skip` remains in lesson Go tests.
- Seven Go test files remain across lesson code, including reference examples. Focused learner tests cover concurrency, workers, deadlock behavior, TCP request/response, repository behavior, and the Raft term/vote transition.
- One C assertion harness checks a page-record serialization contract and out-of-bounds rejection.
- Duplicate challenge source files were removed; 16 distinct challenge files remain. The coding challenge assignment is specified in each work packet.
- Lessons 064, 085, and 090 use the real MIT xv6, CMU BusTub, and MIT 6.5840 assignment/test workflows. DART does not vendor or recreate those upstream test harnesses.

## Semester course exercises

- Subject packs are physically separated under `lessons/S1` and `lessons/S2`: 11 S1 packs and 10 S2 packs, plus shared elective-slot guidance.
- Semester-specific code and unit checks sit beside their course files.
- Nine Python test files specify 34 behaviors across analysis, accounting, statistics/probability, logic/architecture, multimedia, algebra, automata, finance, and accounting information systems. The corresponding worked helpers now implement the contracts and pass these checks.
- Run a track with `python tools/run_semester_tests.py s1` or `s2`, or run both with `python tools/run_semester_tests.py all`.
- The current official semester maquette was not present in the repository. Treat these exercise prompts as study support, not an authoritative syllabus; verify them with current institutional materials.

## Validators

- `validate_course.py` checks lesson numbering, work packet sections, placeholders, stale migrated text, evidence paths, and required semester entry points.
- `validate_code_layer.py` checks lesson code directories, student TODOs, skipped Go tests, challenge duplication after comment removal, and omitted source placeholders.
- `validate_code_artifacts.py` checks lesson artifacts and referenced paths. C syntax checks run only when GCC is installed; it does not equate compilation with behavioral verification.
- `validate_semester_courses.py` checks semester exercise counts, test presence, and Python syntax.

## Environment limits

- Go checks run for lessons 029, 030, 031, 041, 055, and 090. The incomplete starters in 029, 030, 041, and 055 fail their new acceptance tests; 031 and 090 pass.
- No usable C toolchain is available here. GCC is absent, and the installed standalone Clang lacks `stdio.h`; the C page assertion harness, syntax checks, and Lesson 001 checker cannot run until setup is completed.
- Rust examples remain limited and are not compiled by the structural validators.
- Upstream projects are external dependencies. Record the exact course release/commit and the upstream test output in the corresponding lesson evidence.
