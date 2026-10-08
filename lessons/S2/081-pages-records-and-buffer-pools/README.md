# Lesson 081 — Pages, Records, and Buffer Pools

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** A storage engine moves data in pages rather than issuing disk operations for individual fields. Page and record layout determine how records are found, modified, and persisted.

**The contract.** A slotted page can place its slot directory at one end and variable-length record bytes at the other, with free space between them. The page size is an engine/build format choice; it is not necessarily selected to equal the filesystem block size.

**Worked example.** Define a fixed page byte array, header fields, slot entries `(offset, length)`, and a free-space invariant. Insert records from one end while the slot directory grows from the other; reject an insertion before the regions overlap.

## 2. Mechanism and vocabulary

**Mechanism.** A buffer pool maps page identifiers to in-memory frames. A pin/reference count prevents eviction while a page is in use; a dirty bit records modifications that may need flushing. A replacement policy selects an unpinned frame when another page must be loaded.

**Second point.** Variable-width records need metadata describing their location and length. Deleting a record may leave fragmented free space; total free bytes alone do not prove a new record fits contiguously.

Keep the terms page, frame, pin, dirty, slot, and record distinct.

**Worked example.** Choose a page size such as 4096 bytes. Diagram three records with actual offsets and show how the slot directory and record area grow toward each other.

## 3. Worked case from the reference material

Design a slotted page on paper first: include a page header, slot count, free-space boundaries, slot entries, and serialized record bytes. Then insert three records of different lengths and calculate every offset.

For each insertion, state the page state before and after. A record is valid only if its byte range lies within the page and does not overlap another live record or the slot directory.

## 4. Verification, failure, and evidence

**Verification.** Show offsets, slot entries, round-trip values, used/free bytes, and storage-overhead calculation for fixed- and variable-width records. Check an empty page, exact fit, one-byte overflow, and corrupted offset.

**Failure mode.** Delete middle records without compacting, then attempt to insert a record that fits in total free bytes but not in any contiguous region. Explain why free-space total is insufficient and how compaction relocates records while updating slot offsets.

**Completion criterion.** You can preserve page/slot bounds, round-trip records, explain pin/dirty/eviction roles, and identify when fragmented free space requires compaction.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
