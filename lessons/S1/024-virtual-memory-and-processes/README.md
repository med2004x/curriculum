# Lesson 024 — Virtual Memory and Address Spaces

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Explain why virtual memory exists, how address translation works, and what a page fault and a TLB miss actually are.

**The contract.** **The problem virtual memory solves.** Without it, every running program would need to know the physical addresses of real RAM, would collide with every other running program's memory, and a buggy program could read or corrupt another program's data or the OS itself.

**Worked example.** On Linux, reserve or allocate a large buffer, record resident memory, then touch one byte on each page in a small prefix. Compare this with touching every page. The experiment illustrates demand paging, but the exact RSS change depends on allocator behavior, page size, kernel policy, and how the memory was obtained.

## 2. Mechanism and vocabulary

**Mechanism.** A process issues virtual addresses. Hardware translation uses page-table state to map a virtual page to a physical frame with access permissions; a TLB caches recent translations. The operating system maintains the mappings and handles faults that require policy or backing-store work.

**Second point.** Address-space width and layout are architecture- and operating-system-dependent, not one universal range. Virtual memory also supports protection, sharing, mapped files, and sparse allocation; it does not mean every virtual page immediately has a distinct physical RAM page.

On first access to a page that is not currently resident or mapped in the required way, a page fault transfers control to the kernel. A valid demand-zero or file-backed page may be made available and execution may resume. An invalid or disallowed access may instead produce a signal such as `SIGSEGV`; “page fault” itself does not mean “segmentation fault.”

**Worked example.** If a program maps a file, the mapping describes a virtual range backed by file contents. Accessing a page may cause the kernel to bring data into memory. Compare the mapping's virtual size with RSS; do not infer physical allocation solely from a successful mapping call.

## 3. Worked case from the reference material

Worked example: On Linux, write a small C or Go program that allocates a large buffer (e.g.

## 4. Verification, failure, and evidence

**Verification.** Record mapping/allocation size and `VmRSS` at defined points: before, after the mapping or allocation, and after touching selected pages. Expect measurements to be approximate and platform-sensitive; record page-touch pattern and avoid interpreting a single RSS delta as a universal rule.

**Failure mode.** Compare a valid mapped-page access with an access beyond the permitted mapping. Explain why the first can cause a recoverable fault and the second can be rejected, rather than describing both simply as “a page fault.”

**Completion criterion.** You can explain virtual page, physical frame, page-table translation, TLB, RSS, and the distinction between a serviced fault and an invalid access.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
