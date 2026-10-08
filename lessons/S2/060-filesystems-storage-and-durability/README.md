# Lesson 060 — Filesystems, Storage, and Durability

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Understand how a filesystem represents files and directories on disk — inodes, allocation strategies, metadata — and reason about crash consistency: what happens if power is cut mid-write.

**The contract.** In Unix-like filesystems, directory entries associate names with filesystem objects; an inode-based filesystem stores file metadata and data-block mapping separately from a name. A hard link creates another directory entry for the same inode on the same filesystem. A symbolic link is a distinct object containing a path reference.

**Worked example.** Create a file and a hard link, compare inode numbers and link counts, then remove one name and read through the other. The data remains while another hard link or open reference keeps the object alive.

## 2. Mechanism and vocabulary

**Mechanism.** `unlink` removes a directory entry. When no hard links and no open references remain, the filesystem can reclaim the object and its storage. Removing a symlink removes that link object; removing its target name can leave a dangling path reference.

**Second point.** Filesystem durability is distinct from successful `write`: data may be buffered in user space, the kernel page cache, or a device cache. `fsync` requests persistence of file data/metadata as defined by the OS/filesystem; creating or renaming a name may also require syncing the containing directory for crash durability. Exact guarantees depend on filesystem and storage hardware.

Journaling often protects filesystem structural consistency after a crash; it does not necessarily provide application-level transaction atomicity for arbitrary file updates. Applications must design write ordering, temporary-file/rename patterns, and synchronization according to their durability contract.

**Worked example.** For crash-safe replacement, write a temporary file in the target directory, check all writes, synchronize the file, rename it over the destination, then synchronize the directory when the required durability contract demands persistence of the name update.

## 3. Worked case from the reference material

Worked example: Create a file, then create a hard link to it (`ln original.txt hardlink.txt`, not `ln -s` which is a symlink, a genuinely different mechanism).

## 4. Verification, failure, and evidence

**Verification.** Record inode/link-count observations before and after unlinking one hard-link name. Separately describe the persistence protocol's ordering: successful write, file synchronization, rename, and directory synchronization as applicable.

**Failure mode.** Repeat with a symlink, remove the target, and observe a dangling path. Separately omit the synchronization step from a replacement protocol and explain why a process-level success return alone cannot prove survival through power loss.

**Completion criterion.** You can explain name/object identity, hard-link/symlink behavior, unlink lifetime, and the boundary between filesystem consistency and application durability.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
