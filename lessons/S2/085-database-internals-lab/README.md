# Lesson 085 — Database Internals Lab

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Integrate Lessons 1-6 into one substantial, tested database component with documented invariants and measured performance — the database-internals track's equivalent of the integration capstones in earlier tracks.

**The contract.** Choose one component with an observable invariant. For example, a buffer pool must not evict a pinned frame; a B+ tree must route searches according to separators and maintain equal leaf depth; an iterator must emit the same rows as its logical query; a recovery component must preserve acknowledged committed state under its stated assumptions.

**Worked example.** Start with the smallest concrete case described by the source material: If using BusTub directly, complete at least the buffer pool manager project and the B+ tree index project from the official course materials (`https://github.com/cmu-db/bustub`), using its existing test infrastructure to verify correctness — this is the recommended default path, since BusTub's tests are considerably more rigorous than what's practical to con.

## 2. Mechanism and vocabulary

**Mechanism.** Use the repository's chosen teaching system or a bounded component built from prior lessons. Keep the component boundary explicit: page format, buffer-pool lifecycle, index invariant, execution operator, or recovery record. Do not claim a full DBMS implementation when only one component is implemented.

**Second point.** Integration risk lives at assumptions between components: page IDs versus frame IDs, pin lifetime versus eviction, index keys versus table records, iterator end-of-stream versus errors, and log sequence numbers versus persistence order.

**Worked example.** If building your own integrated system instead (a reasonable alternative if you want the components to be ones you fully wrote yourself, end to end), integrate your own Lessons 1-6 implementations into one coherent system with this minimum scope:

## 3. Worked case from the reference material

Worked example: If using BusTub directly, complete at least the buffer pool manager project and the B+ tree index project from the official course materials (`https://github.com/cmu-db/bustub`), using its existing test infrastructure to verify correctness — this is the recommended default path, since BusTub's tests are considerably more rigorous than what's practical to con.

## 4. Verification, failure, and evidence

**Verification.** Record the selected invariant, the cross-component test that checks it, the exact test scope, and a reproducible measurement with workload size and environment. A passing test covers its cases, not all possible states.

**Failure mode.** Break one interface assumption in a controlled branch—for example, unpin a page before its consumer finishes using the frame, or return an iterator row after end-of-stream—and show which invariant/test fails.

**Completion criterion.** You can explain one integrated component path, state its invariants, show a cross-component test and a deliberate failure, and support one performance claim with measured evidence.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
