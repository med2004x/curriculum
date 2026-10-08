# Lesson 081 — Pages, Records, and Buffer Pools

## Threads
**Primary:** DB Internals/C

**Supporting:** storage layout; pages; buffer pool

**Math companion:** matrix/index reasoning

**Time:** 3–4 hours. **Prerequisites:** Lesson 080.

## What you will be able to do

1. Design a slotted page with an explicit header, slot directory, and record bytes.
2. Serialize and deserialize variable-length records while preserving page and slot bounds.
3. Explain buffer-pool frames, pins, dirty state, and eviction.
4. Demonstrate fragmented free space and explain when compaction is needed.

## Retrieve first (no notes)

1. Which kernel do ordinary containers on one host share?
2. What do namespaces isolate?
3. What do cgroups account for or limit?
4. What does an image digest or SBOM establish, and what does it not prove?

## 1. Problem and model

**What problem this solves.** A storage engine moves data in pages rather than issuing disk operations for individual fields. Page and record layout determine how records are found, modified, and persisted.

**The contract.** A slotted page can place its slot directory at one end and variable-length record bytes at the other, with free space between them. The page size is an engine/build format choice; it is not necessarily selected to equal the filesystem block size.

**Worked example.** Define a fixed page byte array, header fields, slot entries `(offset, length)`, and a free-space invariant. Insert records from one end while the slot directory grows from the other; reject an insertion before the regions overlap.

Self-check: write the page layout and record encoding before calculating offsets.

## 2. Mechanism and vocabulary

**Mechanism.** A buffer pool maps page identifiers to in-memory frames. A pin/reference count prevents eviction while a page is in use; a dirty bit records modifications that may need flushing. A replacement policy selects an unpinned frame when another page must be loaded.

**Second point.** Variable-width records need metadata describing their location and length. Deleting a record may leave fragmented free space; total free bytes alone do not prove a new record fits contiguously.

Keep the terms page, frame, pin, dirty, slot, and record distinct.

**Worked example.** Choose a page size such as 4096 bytes. Diagram three records with actual offsets and show how the slot directory and record area grow toward each other.

Self-check: which invariant prevents a slot from pointing outside the page?

## 3. Worked case from the reference material

Design a slotted page on paper first: include a page header, slot count, free-space boundaries, slot entries, and serialized record bytes. Then insert three records of different lengths and calculate every offset.

For each insertion, state the page state before and after. A record is valid only if its byte range lies within the page and does not overlap another live record or the slot directory.

Self-check: what should happen when the next record would make the slot directory meet the record area?

## 4. Verification, failure, and evidence

**Verification.** Show offsets, slot entries, round-trip values, used/free bytes, and storage-overhead calculation for fixed- and variable-width records. Check an empty page, exact fit, one-byte overflow, and corrupted offset.

**Failure mode.** Delete middle records without compacting, then attempt to insert a record that fits in total free bytes but not in any contiguous region. Explain why free-space total is insufficient and how compaction relocates records while updating slot offsets.

**Completion criterion.** You can preserve page/slot bounds, round-trip records, explain pin/dirty/eviction roles, and identify when fragmented free space requires compaction.

## Predict, then observe

1. Predict the first record's offset in your chosen layout.
2. Predict how a slot entry changes after insertion.
3. Predict what happens when the free-space regions meet.
4. Predict whether a deleted record's bytes need immediate erasure for the slot to become invalid.
5. Predict when a buffer frame may be evicted.
6. Predict what a dirty bit signals.
7. Predict how fragmentation can block insertion despite sufficient aggregate free bytes.
8. State which bytes are metadata and which bytes encode records.

## Build from blank

1. Define a page header and slot record format.
2. Implement insertion and retrieval for variable-length records with bounds checks.
3. Test empty, exact-fit, overflow, deletion, and compaction cases.
4. Add a small buffer-pool model with page ID, pin count, dirty state, and an eviction decision.

Do not rely on host struct layout as a portable on-disk format; define byte order and field encoding explicitly.

## Failure drill

1. Insert records until the page is nearly full.
2. Delete records from the middle without compacting.
3. Attempt an insertion that fits in total free bytes but not in a contiguous gap.
4. Compact live records, update all slot offsets, and verify every value again.

## Transfer

Change the record encoding or page size. Recalculate offsets and show which page invariants remain unchanged.

## Retrieval questions

1. What fields does a slot entry need for variable-length records?
2. What prevents slot and record storage from overlapping?
3. What does a pin count protect against?
4. What does a dirty bit indicate?
5. Why can free-space fragmentation prevent insertion?
6. What metadata must compaction update?

## Evidence to record

Create `lesson081/EVIDENCE.md` with exactly this structure:

```text
# Lesson 081 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Treating a page size as universally equal to a filesystem block size.
- Writing a host struct directly as a portable disk representation.
- Failing to bounds-check offsets and lengths before decoding.
- Assuming total free space is one contiguous region.
- Evicting a pinned frame or forgetting to track dirty pages.

## Precision checks

1. State the page's free-space invariant.
2. Show that every live slot refers to a valid, non-overlapping record range.
3. Explain the pin and dirty indicators.
4. Demonstrate a fragmentation case and its compaction repair.
5. Calculate storage overhead from the chosen format.

This lesson models storage layout and buffer management, not PostgreSQL's SQL query planner.

## University revision bridge

Connect page layout to block-oriented storage and record organization. Keep byte offsets, address bounds, and invariants explicit.

## Assigned material

- `09-postgresql-engineering/lesson-11-indexes-explain/README.md`
- `10-operating-systems/lesson-05-virtual-memory/README.md`

## Mastery gate

You pass when you can encode and decode a mixed record set in a slotted page, prove all bounds, explain buffer-frame lifecycle, and reproduce/repair fragmented free space.
