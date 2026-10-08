#!/usr/bin/env python3
from __future__ import annotations

from pathlib import Path
import re
import sys

ROOT = Path(__file__).resolve().parents[1]
LESSONS = ROOT / 'lessons'
errors: list[str] = []

lesson_dirs = sorted((p for semester in ('S1', 'S2') for p in (LESSONS / semester).iterdir() if p.is_dir() and re.match(r'^\d{3}-', p.name)), key=lambda p: p.name)
if len(lesson_dirs) != 100:
    errors.append(f'expected 100 lesson directories, found {len(lesson_dirs)}')

required_problem_set = ['## Problems', '## Run and check', '## Submit']

for i, d in enumerate(lesson_dirs, 1):
    if not d.name.startswith(f'{i:03d}-'):
        errors.append(f'numbering mismatch: {d.name}, expected {i:03d}-')
    expected_semester = 'S1' if i <= 50 else 'S2'
    if d.parent.name != expected_semester:
        errors.append(f'{d.name}: expected under lessons/{expected_semester}/, found lessons/{d.parent.name}/')
    readme = d/'README.md'
    work = d/'WORK.md'
    if not readme.exists(): errors.append(f'{d.name}: missing README.md')
    if not work.exists():
        errors.append(f'{d.name}: missing WORK.md')
        continue
    txt = work.read_text(encoding='utf-8', errors='replace')
    work_headings = required_problem_set
    for h in work_headings:
        if h not in txt: errors.append(f'{d.name}: missing {h}')
    if not re.search(r'^\s*\d+\.\s+|^### Task\s+\d+', txt, re.M):
        errors.append(f'{d.name}: problem set has no numbered tasks')
    for marker in ['TBD', 'FIXME', 'FILL IN HERE', 'INSERT DATA HERE']:
        if marker.lower() in txt.lower(): errors.append(f'{d.name}: unresolved placeholder {marker}')
    if 'practice/' in txt or 'exams/' in txt or 'training/' in txt or 'exam-work/' in txt:
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

allowed_dirs = {'.git', '.vscode', 'lessons', 'tools'}
extras = [p.name for p in ROOT.iterdir() if p.is_dir() and p.name not in allowed_dirs]
if extras:
    errors.append('unexpected top-level content directories: ' + ', '.join(sorted(extras)))

# Root curriculum docs must exist.
for name in ['README.md', 'START-TODAY.md', 'SETUP.md', 'CURRICULUM.md', 'LESSON-INDEX.md', 'STUDY-LOG.md', 'C-QUICK-REFERENCE.md', 'GO-QUICK-REFERENCE.md', 'SOURCES.md']:
    if not (ROOT/name).exists(): errors.append(f'missing root file: {name}')
for name in ['lessons/S1', 'lessons/S2']:
    if not (ROOT/name).is_dir(): errors.append(f'missing semester exercise folder: {name}')

# Lecture pages keep the explanation and point directly to the problem card.
for d in lesson_dirs:
    txt=(d/'README.md').read_text(encoding='utf-8', errors='replace')
    if '## Next step' not in txt: errors.append(f'{d.name}: missing direct path to its problem card')
    if not re.search(r'^## (?:\d+\.|Problem and model)', txt, re.M): errors.append(f'{d.name}: missing lesson explanation')
    if re.search(r'^## (?:\d+\. )?(?:Retrieval questions|Common mistakes|Evidence to record|Mastery gate)', txt, re.M):
        errors.append(f'{d.name}: repeated study boilerplate remains in README.md')

print(f'lessons: {len(lesson_dirs)}')
print(f'work packets: {sum(1 for d in lesson_dirs if (d/"WORK.md").exists())}')
print(f'top-level content dirs: {len(extras)}')
if errors:
    print(f'ERRORS: {len(errors)}')
    for e in errors[:100]: print('ERROR:', e)
    raise SystemExit(1)
print('PASS')
