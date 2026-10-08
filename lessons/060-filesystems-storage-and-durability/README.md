# Lesson 060 — Filesystems, Storage, and Durability

## Threads
**Primary:** OS/DB

**Supporting:** inodes/blocks; persistence; failure

**Math companion:** reliability and probability

**Time:** 3–4 hours. **Prerequisites:** Lesson 059.

**Time:** 3–4 hours. **Prerequisites:** Lesson 059.

## What you will be able to do

1. Distinguish a directory entry/name from an inode-like file identity.
2. Explain the difference between hard links and symbolic links, including unlink behavior.
3. Explain why buffered writes and filesystem journaling do not automatically prove an application's data is durable.
4. Describe the role of `fsync`/`fdatasync` and directory synchronization in a persistence protocol.

## Retrieve first (no notes)

1. Define response time and turnaround time.
2. How does non-preemptive SJF differ from SRTF?
3. What assumptions are needed for the SJF optimality statement?
4. How does a time quantum affect Round Robin scheduling?

## 1. Problem and model

**What problem this solves.** Understand how a filesystem represents files and directories on disk — inodes, allocation strategies, metadata — and reason about crash consistency: what happens if power is cut mid-write.

**The contract.** In Unix-like filesystems, directory entries associate names with filesystem objects; an inode-based filesystem stores file metadata and data-block mapping separately from a name. A hard link creates another directory entry for the same inode on the same filesystem. A symbolic link is a distinct object containing a path reference.

**Worked example.** Create a file and a hard link, compare inode numbers and link counts, then remove one name and read through the other. The data remains while another hard link or open reference keeps the object alive.

Self-check: hard links generally cannot cross filesystem boundaries and are usually restricted for directories; a symlink can refer to a path on another filesystem.

## 2. Mechanism and vocabulary

**Mechanism.** `unlink` removes a directory entry. When no hard links and no open references remain, the filesystem can reclaim the object and its storage. Removing a symlink removes that link object; removing its target name can leave a dangling path reference.

**Second point.** Filesystem durability is distinct from successful `write`: data may be buffered in user space, the kernel page cache, or a device cache. `fsync` requests persistence of file data/metadata as defined by the OS/filesystem; creating or renaming a name may also require syncing the containing directory for crash durability. Exact guarantees depend on filesystem and storage hardware.

Journaling often protects filesystem structural consistency after a crash; it does not necessarily provide application-level transaction atomicity for arbitrary file updates. Applications must design write ordering, temporary-file/rename patterns, and synchronization according to their durability contract.

**Worked example.** For crash-safe replacement, write a temporary file in the target directory, check all writes, synchronize the file, rename it over the destination, then synchronize the directory when the required durability contract demands persistence of the name update.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Create a file, then create a hard link to it (`ln original.txt hardlink.txt`, not `ln -s` which is a symlink, a genuinely different mechanism).

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Record inode/link-count observations before and after unlinking one hard-link name. Separately describe the persistence protocol's ordering: successful write, file synchronization, rename, and directory synchronization as applicable.

**Failure mode.** Repeat with a symlink, remove the target, and observe a dangling path. Separately omit the synchronization step from a replacement protocol and explain why a process-level success return alone cannot prove survival through power loss.

**Completion criterion.** You can explain name/object identity, hard-link/symlink behavior, unlink lifetime, and the boundary between filesystem consistency and application durability.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict inode and link-count changes after creating a hard link.
2. Predict what remains accessible after deleting one of two hard-link names.
3. Predict the outcome when a symlink target is removed.
4. Predict whether a successful buffered write proves crash durability.
5. Predict what synchronization steps follow writing and renaming a replacement file.
6. Predict which filesystem structures journaling may protect and which application invariants remain your responsibility.
7. Predict whether a hard link can name an object on another filesystem.
8. State which operation changes a name and which operation changes file contents.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Create a file and hard link; compare inode identity and link count.

2. Remove one name and verify the other still accesses the same file object.

3. Create and remove a symlink target; explain why the link can become dangling.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Run the hard-link and symlink experiments and compare object identity and unlink behavior.

2. Describe a temporary-file/rename update and list the synchronization steps required by the intended durability contract.

3. Explain why filesystem journaling alone does not guarantee a multi-file application update is atomic.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one persistence constraint: replace a file in place versus via temporary file and rename, or sync only the file versus file plus containing directory. Explain which crash outcomes remain possible.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What relationship does a directory entry have to an inode-based file object?
2. How do hard links and symlinks differ?
3. What does `unlink` remove?
4. Why does successful `write` not necessarily imply durable storage?
5. What is the purpose of synchronizing a file and its containing directory?
6. What does a filesystem journal not guarantee for application-level transactions?

## Evidence to record

Create `lesson060/EVIDENCE.md` with exactly this structure:

```text
# Lesson 060 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Treating a filename as the file's identity.
- Assuming a symlink keeps its target alive.
- Assuming unlink immediately destroys an object still referenced by another hard link or open descriptor.
- Treating successful `write` as proof of power-loss durability.
- Treating filesystem journaling as an application transaction protocol.

## Precision checks

1. Explain directory name to inode mapping.
2. Distinguish link count from symlink target path.
3. State the file and directory synchronization steps for the chosen replacement protocol.
4. Explain the difference between filesystem consistency and application durability.
5. Identify one crash window the protocol is designed to reduce.

Filesystem naming, object lifetime, and durable update are separate mechanisms; a robust design states which guarantees it needs from each.
## University revision bridge

For the LIG track, connect inode-based identity and directory structure to OS file semantics, then distinguish journaling from application-level durability.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 060 — Filesystems, Storage, and Durability without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You can explain the inode/filename separation and have directly demonstrated it with a hard link surviving the original filename's deletion, you can distinguish mtime/ctime/atime and state what specifically changes each one, and you can explain in your own words what a filesystem journal protects against and why database engines still need their own addition. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
