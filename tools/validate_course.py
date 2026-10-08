#!/usr/bin/env python3
from __future__ import annotations

from pathlib import Path
import re
import sys

ROOT = Path(__file__).resolve().parents[1]
LESSONS = ROOT / 'lessons'
errors: list[str] = []

lesson_dirs = sorted(p for p in LESSONS.iterdir() if p.is_dir() and re.match(r'^\d{3}-', p.name))
if len(lesson_dirs) != 100:
    errors.append(f'expected 100 lesson directories, found {len(lesson_dirs)}')

required_work = [
    '## 1. Session sequence',
    '## 2. Retrieve before learning',
    '## 3. Guided practice',
    '## 4. Independent exercises',
    '## 5. Problem-solving lab',
    '## 6. Coding challenge',
    '## 8. Debugging and failure lab',
    '## 9. Transfer and design reasoning',
    '## 10. Academic drill',
    '## 11. Self-test',
    '## 12. Evidence and grading',
    '## 13. Mastery gate',
    '## 14. Spaced review',
]

for i, d in enumerate(lesson_dirs, 1):
    if not d.name.startswith(f'{i:03d}-'):
        errors.append(f'numbering mismatch: {d.name}, expected {i:03d}-')
    readme = d/'README.md'
    work = d/'WORK.md'
    if not readme.exists(): errors.append(f'{d.name}: missing README.md')
    if not work.exists():
        errors.append(f'{d.name}: missing WORK.md')
        continue
    txt = work.read_text(encoding='utf-8', errors='replace')
    for h in required_work:
        if h not in txt: errors.append(f'{d.name}: missing {h}')
    if txt.count('### Exercise ') < 5:
        errors.append(f'{d.name}: fewer than 5 independent exercises')
    if '### ' not in txt:
        errors.append(f'{d.name}: no activity subsections')
    for marker in ['TODO', 'TBD', 'FIXME', 'FILL IN HERE', 'INSERT DATA HERE']:
        if marker.lower() in txt.lower(): errors.append(f'{d.name}: unresolved placeholder {marker}')
    if 'practice/' in txt or 'reference/' in txt or 'exams/' in txt or 'training/' in txt or 'exam-work/' in txt:
        errors.append(f'{d.name}: stale split-tree reference in WORK.md')
    # Catch obvious template substitution failures that headings-only checks miss.
    if re.search(r'\bsmallest\s+#\s*\d+\s+[—-]\s+', txt, re.I):
        errors.append(f'{d.name}: malformed lesson-title substitution in WORK.md')
    if d.name == '100-graduation-capstone' and 'Project 15 - Architecture Review' in txt:
        errors.append(f'{d.name}: unrelated Project 15 exercise remains in capstone')
    if 'The old repository contained a useful unsolved exercise pack' in txt:
        errors.append(f'{d.name}: stale migrated exercise-pack boilerplate remains')
    if re.search(r'^# Project \d+\s+[—-]', txt, re.M):
        errors.append(f'{d.name}: stale numbered project pack remains in WORK.md')

    evidence = d / 'EVIDENCE.md'
    if evidence.exists():
        etxt = evidence.read_text(encoding='utf-8', errors='replace')
        for ref in re.findall(r'`(code/[^`]+)`', etxt):
            if not (d / ref).exists():
                errors.append(f'{d.name}: dead evidence path: {ref}')

allowed_dirs = {'.git', '.vscode', 'lessons', 'tools', 'LIG'}
extras = [p.name for p in ROOT.iterdir() if p.is_dir() and p.name not in allowed_dirs]
if extras:
    errors.append('unexpected top-level content directories: ' + ', '.join(sorted(extras)))

# Root curriculum docs must exist.
for name in ['README.md', 'CURRICULUM.md', 'LESSON-INDEX.md', 'STUDY-LOG.md', 'C-QUICK-REFERENCE.md', 'GO-QUICK-REFERENCE.md', 'ACADEMIC.md', 'SOURCES.md']:
    if not (ROOT/name).exists(): errors.append(f'missing root file: {name}')
for name in ['LIG/README.md', 'LIG/S1/README.md', 'LIG/S2/README.md', 'LIG/Shared/CHOICE-SLOTS.md']:
    if not (ROOT/name).exists(): errors.append(f'missing semester track file: {name}')

# The original lecture READMEs must retain their structure and minimum depth.
for d in lesson_dirs:
    txt=(d/'README.md').read_text(encoding='utf-8', errors='replace')
    if len(txt.splitlines()) < 120: errors.append(f'{d.name}: lecture unexpectedly short')
    if txt.count('## Threads') != 1: errors.append(f'{d.name}: lecture missing Threads section')
    if '## What you will be able to do' not in txt: errors.append(f'{d.name}: lecture missing learning outcomes')
    if d.name != '074-graceful-shutdown-and-lifecycles' and '## Mastery gate' not in txt: errors.append(f'{d.name}: lecture missing mastery gate')

print(f'lessons: {len(lesson_dirs)}')
print(f'work packets: {sum(1 for d in lesson_dirs if (d/"WORK.md").exists())}')
print(f'top-level content dirs: {len(extras)}')
if errors:
    print(f'ERRORS: {len(errors)}')
    for e in errors[:100]: print('ERROR:', e)
    raise SystemExit(1)
print('PASS')
