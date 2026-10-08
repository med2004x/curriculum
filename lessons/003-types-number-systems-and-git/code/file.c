git init                    create a repository in this directory
git status                  what changed? what is staged?
git add file.c              stage a file (or: git add -p to choose hunks)
git commit -m "message"     snapshot what is staged
git log --oneline           history, one line per commit
git diff                    unstaged changes;   git diff --staged  = staged changes
git restore file.c          discard unstaged edits to a file (destructive)
