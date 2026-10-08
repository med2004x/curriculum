# Lesson 063 — Virtual Memory, Allocators, and Garbage Collection

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Go deeper than computer architecture Lesson 6's MMU/TLB mechanics into the OS's role specifically: managing page tables per process, handling page faults as software (not just hardware) events, and inspecting real memory mappings on a running Linux process.

**The contract.** The MMU performs address translation using page-table state; the kernel creates and maintains mappings and handles faults that require policy or I/O. `/proc/<pid>/maps` reports virtual mappings and permissions, not which pages are resident. The `[heap]` label is only one possible allocator region; runtimes may obtain memory through multiple mappings.

**Worked example.** Start with the smallest concrete case described by the source material: Run `cat /proc/self/maps` (or `/proc/<pid>/maps` for any running process) and identify at least: the executable's code region (permissions should show `r-xp`, executable but not writable — direct evidence of the write-protection on code mentioned in computer architecture Lesson 6), the heap region, the stack region, and at least one shared library mapping.

## 2. Mechanism and vocabulary

**Mechanism.** A minor page fault can be serviced without disk I/O, for example when a page is already cached or a mapping is established. A major fault requires I/O to resolve the access in Linux accounting. Neither counter alone reveals every cause or gives a portable prediction of timing.

**Second point.** A C allocator manages application allocations using implementation-specific arenas/blocks and may request more virtual memory from the OS. Go's runtime similarly manages its heap and exposes garbage-collected objects when they are no longer reachable; collection and returning pages to the OS are distinct events.

Go's garbage collector trades CPU work and memory headroom against allocation pressure. `GOGC` changes the target heap-growth policy; `runtime.ReadMemStats` and `GODEBUG=gctrace=1` can provide observations, but metrics are runtime-version-sensitive and do not explain every allocation by themselves.

**Worked example.** Inspect `/proc/<pid>/maps` for a chosen process and label mapped files, permissions, anonymous ranges, and stack where visible. Do not assume the stack/heap ordering or a `[heap]` mapping fully describes a language runtime's allocations.

## 3. Worked case from the reference material

Worked example: Run `cat /proc/self/maps` (or `/proc/<pid>/maps` for any running process) and identify at least: the executable's code region (permissions should show `r-xp`, executable but not writable — direct evidence of the write-protection on code mentioned in computer architecture Lesson 6), the heap region, the stack region, and at least one shared library mapping.

## 4. Verification, failure, and evidence

**Verification.** Record the process and workload associated with mapping/fault/heap metrics, the measurement interval, and whether data was warmed. A file read may cause major faults when I/O is needed, but cached pages and OS accounting can change the result; report measurements rather than a fixed expected count.

**Failure mode.** Compare a process's virtual map size with its resident memory and explain why they differ. In Go, keep references to allocated objects in one run and drop them in another; explain that collection depends on reachability and runtime policy, not merely on leaving a function.

**Completion criterion.** You can interpret mappings and memory metrics with their limits, explain allocator/GC layering, and distinguish object reclamation from returning memory to the operating system.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
