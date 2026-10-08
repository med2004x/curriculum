# Lesson 063 — Virtual Memory, Allocators, and Garbage Collection

## Threads
**Primary:** OS/C/Go

**Supporting:** pages; allocation strategies; GC trade-offs

**Math companion:** growth rates

**Time:** 3–4 hours. **Prerequisites:** Lesson 062.

**Time:** 3–4 hours. **Prerequisites:** Lesson 062.

## What you will be able to do

1. Interpret selected regions and permissions in `/proc/<pid>/maps`.
2. Distinguish virtual mapping, resident pages, minor faults, and major faults.
3. Explain how C allocators and Go's runtime obtain and reuse memory at different abstraction layers.
4. Explain Go garbage collection as reachability-based reclamation and interpret heap/GC metrics without assuming memory is immediately returned to the OS.

## Retrieve first (no notes)

1. What does a mutex unlock/lock synchronization edge guarantee?
2. Why must a condition predicate be checked in a loop?
3. What is a Linux futex used for when a thread must block?
4. Why should application code rely on Go's sync-package contract rather than futex internals?

## 1. Problem and model

**What problem this solves.** Go deeper than computer architecture Lesson 6's MMU/TLB mechanics into the OS's role specifically: managing page tables per process, handling page faults as software (not just hardware) events, and inspecting real memory mappings on a running Linux process.

**The contract.** The MMU performs address translation using page-table state; the kernel creates and maintains mappings and handles faults that require policy or I/O. `/proc/<pid>/maps` reports virtual mappings and permissions, not which pages are resident. The `[heap]` label is only one possible allocator region; runtimes may obtain memory through multiple mappings.

**Worked example.** Start with the smallest concrete case described by the source material: Run `cat /proc/self/maps` (or `/proc/<pid>/maps` for any running process) and identify at least: the executable's code region (permissions should show `r-xp`, executable but not writable — direct evidence of the write-protection on code mentioned in computer architecture Lesson 6), the heap region, the stack region, and at least one shared library mapping.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** A minor page fault can be serviced without disk I/O, for example when a page is already cached or a mapping is established. A major fault requires I/O to resolve the access in Linux accounting. Neither counter alone reveals every cause or gives a portable prediction of timing.

**Second point.** A C allocator manages application allocations using implementation-specific arenas/blocks and may request more virtual memory from the OS. Go's runtime similarly manages its heap and exposes garbage-collected objects when they are no longer reachable; collection and returning pages to the OS are distinct events.

Go's garbage collector trades CPU work and memory headroom against allocation pressure. `GOGC` changes the target heap-growth policy; `runtime.ReadMemStats` and `GODEBUG=gctrace=1` can provide observations, but metrics are runtime-version-sensitive and do not explain every allocation by themselves.

**Worked example.** Inspect `/proc/<pid>/maps` for a chosen process and label mapped files, permissions, anonymous ranges, and stack where visible. Do not assume the stack/heap ordering or a `[heap]` mapping fully describes a language runtime's allocations.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Run `cat /proc/self/maps` (or `/proc/<pid>/maps` for any running process) and identify at least: the executable's code region (permissions should show `r-xp`, executable but not writable — direct evidence of the write-protection on code mentioned in computer architecture Lesson 6), the heap region, the stack region, and at least one shared library mapping.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Record the process and workload associated with mapping/fault/heap metrics, the measurement interval, and whether data was warmed. A file read may cause major faults when I/O is needed, but cached pages and OS accounting can change the result; report measurements rather than a fixed expected count.

**Failure mode.** Compare a process's virtual map size with its resident memory and explain why they differ. In Go, keep references to allocated objects in one run and drop them in another; explain that collection depends on reachability and runtime policy, not merely on leaving a function.

**Completion criterion.** You can interpret mappings and memory metrics with their limits, explain allocator/GC layering, and distinguish object reclamation from returning memory to the operating system.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict what `/proc/<pid>/maps` says about virtual ranges and permissions.
2. Predict whether a mapped range must be resident in RAM.
3. Distinguish minor from major page faults in Linux accounting.
4. Predict what a warm file cache might change on a second read.
5. Predict whether a Go object becomes collectible while a reference remains reachable.
6. Predict whether garbage collection immediately returns all freed heap memory to the OS.
7. Predict what changes when `GOGC` is lowered, qualitatively.
8. Separate the measured metric from the allocator/GC explanation.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Inspect `/proc/<pid>/maps` for a selected running process and classify several mappings by permissions and backing.

2. Record virtual-size and resident-memory metrics for a process, and explain why they are not interchangeable.

3. Compare C allocation behavior with Go heap/GC observations; state what each tool measures and what it cannot establish.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Rerun step 3's file-reading test a second time, immediately after the first run, without doing anything to clear caches in between.

2. Observe the major fault count drop sharply (likely to near zero) on the second run.

3. Explain why: the OS's page cache now holds the file's contents in RAM from the first read, so the second read's page faults are resolved as *minor* faults (the data's already in memory, just needs mapping into this process) rather than major faults requiring an actual disk read — this is the same mechanism that makes "warm cache" reads of recently-accessed f.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one memory-pressure condition: change working-set size, retain versus release references in Go, or read a previously cached file. Explain which metric should change and which allocator/runtime policy may obscure it.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What does `/proc/<pid>/maps` show?
2. What is the difference between virtual size and resident set size?
3. What distinguishes minor and major faults in Linux accounting?
4. How does a Go garbage collector decide whether an object is reachable?
5. Why may collected heap space remain reserved by the runtime?
6. What does `GOGC` influence?

## Evidence to record

Create `lesson063/EVIDENCE.md` with exactly this structure:

```text
# Lesson 063 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Treating a mapped virtual range as proof of resident physical memory.
- Assuming `[heap]` accounts for every allocator/runtime mapping.
- Treating page-fault counts as portable latency guarantees.
- Assuming every collection returns all memory to the OS immediately.
- Equating dropping a local variable with proving an object unreachable.

## Precision checks

1. Interpret mapping permissions and backing for selected regions.
2. Distinguish virtual, resident, and fault metrics.
3. Explain C allocator versus Go runtime responsibilities.
4. Explain reachability and the separation between collection and OS release.
5. State the workload and runtime settings behind memory observations.

Memory metrics are observations at different layers; a process map, OS residency count, and language heap statistic answer different questions.
## University revision bridge

For the LIG track, connect page mappings and fault handling to allocator and garbage-collector behavior, keeping OS accounting distinct from language-runtime metrics.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 063 — Virtual Memory, Allocators, and Garbage Collection without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You can read and interpret `/proc/<pid>/maps` for a real running process, you've measured a concrete difference in minor vs. major fault counts between memory-only and disk-backed access patterns, and you've directly observed page-cache warming reduce major faults on a repeated read, connecting it to real-world caching behavior you can now recognize rather t. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
