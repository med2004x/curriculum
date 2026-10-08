# Lesson 024 — Virtual Memory and Address Spaces

## Threads
**Primary:** Architecture/Linux/C

**Supporting:** paging; virtual addresses; processes

**Math companion:** functions and mappings

**Time:** 3–4 hours. **Prerequisites:** Lesson 023.

**Time:** 3–4 hours. **Prerequisites:** Lesson 023.

## What you will be able to do

1. Explain the purpose of per-process virtual address spaces.
2. Describe virtual-page to physical-frame translation and the role of page tables and the TLB.
3. Distinguish reserving or mapping address space from making pages resident by accessing them.
4. Explain how a page fault can be serviced or can terminate a process, depending on the access and mapping.

## Retrieve first (no notes)

1. What are spatial locality and temporal locality?
2. Which dimension is contiguous in a C row-major array?
3. Why can a working set larger than cache capacity make access order matter?
4. What does a timing measurement show directly, and what does it only suggest?

## 1. Problem and model

**What problem this solves.** Explain why virtual memory exists, how address translation works, and what a page fault and a TLB miss actually are.

**The contract.** **The problem virtual memory solves.** Without it, every running program would need to know the physical addresses of real RAM, would collide with every other running program's memory, and a buggy program could read or corrupt another program's data or the OS itself.

**Worked example.** On Linux, reserve or allocate a large buffer, record resident memory, then touch one byte on each page in a small prefix. Compare this with touching every page. The experiment illustrates demand paging, but the exact RSS change depends on allocator behavior, page size, kernel policy, and how the memory was obtained.

Self-check: distinguish virtual address-space size from resident set size (RSS); they are different measurements.

## 2. Mechanism and vocabulary

**Mechanism.** A process issues virtual addresses. Hardware translation uses page-table state to map a virtual page to a physical frame with access permissions; a TLB caches recent translations. The operating system maintains the mappings and handles faults that require policy or backing-store work.

**Second point.** Address-space width and layout are architecture- and operating-system-dependent, not one universal range. Virtual memory also supports protection, sharing, mapped files, and sparse allocation; it does not mean every virtual page immediately has a distinct physical RAM page.

On first access to a page that is not currently resident or mapped in the required way, a page fault transfers control to the kernel. A valid demand-zero or file-backed page may be made available and execution may resume. An invalid or disallowed access may instead produce a signal such as `SIGSEGV`; “page fault” itself does not mean “segmentation fault.”

**Worked example.** If a program maps a file, the mapping describes a virtual range backed by file contents. Accessing a page may cause the kernel to bring data into memory. Compare the mapping's virtual size with RSS; do not infer physical allocation solely from a successful mapping call.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: On Linux, write a small C or Go program that allocates a large buffer (e.g.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Record mapping/allocation size and `VmRSS` at defined points: before, after the mapping or allocation, and after touching selected pages. Expect measurements to be approximate and platform-sensitive; record page-touch pattern and avoid interpreting a single RSS delta as a universal rule.

**Failure mode.** Compare a valid mapped-page access with an access beyond the permitted mapping. Explain why the first can cause a recoverable fault and the second can be rejected, rather than describing both simply as “a page fault.”

**Completion criterion.** You can explain virtual page, physical frame, page-table translation, TLB, RSS, and the distinction between a serviced fault and an invalid access.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict whether mapping a large region necessarily makes the entire region resident.
2. Predict how touching one byte on each of several pages affects RSS qualitatively.
3. Predict what a TLB hit avoids compared with a page-table walk.
4. Predict whether a valid but nonresident mapped page must terminate the process.
5. Predict what happens when code accesses an address outside a permitted mapping.
6. Predict why two operating systems may report different RSS changes for the same source program.
7. Distinguish the address-space size from resident memory in your prediction.
8. Record the assumptions about mapping type and page size before interpreting measurements.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Create a Linux program that maps or allocates a buffer and records its initial memory metrics.

2. Touch one byte per page in a selected range, then separately touch the whole range.

3. Record `/proc/self/status` `VmSize` and `VmRSS` before and after each phase; explain what these process metrics indicate and their limitations.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Access an address in a valid mapping that has not yet been touched, then explain how the kernel may service the fault.

2. Attempt an access outside the mapping only as a controlled failure case; distinguish the resulting invalid-access signal from demand paging.

3. Explain the observed metrics in terms of page granularity while noting that allocator, kernel, shared pages, and accounting details can affect the exact values.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one constraint: use a file-backed mapping instead of anonymous memory, or change the fraction of pages touched. Explain what remains the same in address translation and what changes in the backing policy.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What does a page table map?
2. What does a TLB cache?
3. How do virtual size and RSS differ?
4. Why can a valid access to a nonresident page fault and still resume?
5. What makes an invalid access different from demand paging?
6. Why are RSS changes approximate and platform-dependent?

## Evidence to record

Create `lesson024/EVIDENCE.md` with exactly this structure:

```text
# Lesson 024 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Treating virtual allocation as proof that all pages occupy physical RAM.
- Treating every page fault as a fatal segmentation fault.
- Assuming one architecture has a universal virtual-address range.
- Equating `VmSize` with `VmRSS`.
- Claiming a specific RSS delta without accounting for measurement and kernel policy.

## Precision checks

1. Draw the path from virtual address to virtual page to physical frame.
2. Explain when the TLB helps and when the kernel handles a page fault.
3. Distinguish mapped size, resident size, and permitted access.
4. Explain one valid fault and one invalid access.
5. State the platform and measurement caveats for RSS evidence.

Use memory metrics as observations of operating-system accounting, not as direct proof of every internal allocation decision.
## University revision bridge

For the LIG track, connect page-based address translation and process isolation to the formal operating-systems model, then interpret Linux metrics with their platform-specific limits.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can draw virtual-to-physical translation, explain the roles of page tables and the TLB, distinguish virtual size from RSS, and classify a serviced page fault versus an invalid access.
