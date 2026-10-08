# Lesson 060 — Filesystems, Storage, and Durability Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.c`.

## Problems

1. Create a file and hard link; compare inode identity and link count.

2. Remove one name and verify the other still accesses the same file object.

3. Create and remove a symlink target; explain why the link can become dangling.

## Run and check

```powershell
Set-Location lessons/S2/060-filesystems-storage-and-durability/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Optional challenge

### Filesystem model

Write a small model of blocks, metadata, and file growth. Simulate allocation and crash points and explain what survives under each durability policy.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
