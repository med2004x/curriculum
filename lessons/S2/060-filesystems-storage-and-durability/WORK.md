# Lesson 060 — Problem Set

**Topic:** Filesystems, Storage, and Durability.
**Your files:** `code/exercise.c`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Create a file and hard link; compare inode identity and link count.

2. Remove one name and verify the other still accesses the same file object.

3. Create and remove a symlink target; explain why the link can become dangling.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/060-filesystems-storage-and-durability/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Stretch problem

### Filesystem model

Write a small model of blocks, metadata, and file growth. Simulate allocation and crash points and explain what survives under each durability policy.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
