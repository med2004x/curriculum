#!/usr/bin/env python3
from pathlib import Path
import re

ROOT = Path(__file__).resolve().parents[1]
REQUIRED = {
    'S1': [
        'lessons/004-conditions-logic-and-proofs/ACADEMIC-ADS1-S1.md',
        'lessons/009-files-paths-argv-and-permissions/ACADEMIC-OS1-S1.md',
        'lessons/022-logic-gates-and-cpu-datapath/ACADEMIC-LOGIC-ARCH-S1.md',
        'lessons/006-functions-call-stack-and-debugging/ACADEMIC-ANALYSIS-S1.md',
        'lessons/019-search-sorting-and-profiling/ACADEMIC-STATS-PROB-S1.md',
        'lessons/020-foundations-project-data-tool/ACADEMIC-MANAGEMENT-S1.md',
        'lessons/012-first-integrated-cli-lab/ACADEMIC-ACCOUNTING-S1.md',
        'lessons/003-types-number-systems-and-git/ACADEMIC-DIGITAL-SKILLS-S1.md',
        'lessons/011-testing-verification-and-code-reading/ACADEMIC-BUSINESS-COMMUNICATION-S1.md',
        'lessons/048-postgresql-server-and-relational-thinking/ACADEMIC-DB-MULTIMEDIA-S1.md',
        'lessons/001-programs-machines-and-the-terminal/ACADEMIC-CHOICE-SLOTS.md',
        'lessons/040-algorithms-assessment/ACADEMIC-S1-MASTER.md',
    ],
    'S2': [
        'lessons/036-algorithms-in-go/ACADEMIC-ADS2-S2.md',
        'lessons/037-trees-graphs-and-discrete-structure/ACADEMIC-ALGEBRA-S2.md',
        'lessons/038-recursion-dp-and-recurrences/ACADEMIC-MATH-LOGIC-S2.md',
        'lessons/050-crud-transactions-and-safe-state/ACADEMIC-ACCOUNTING-IS-S2.md',
        'lessons/059-queues-scheduling-and-throughput/ACADEMIC-FINANCE-S2.md',
        'lessons/065-go-http-json-and-sql/ACADEMIC-IS-NETWORKS-S2.md',
        'lessons/065-go-http-json-and-sql/ACADEMIC-BUSINESS-COMMUNICATION-S2.md',
        'lessons/058-process-models-and-scheduling/ACADEMIC-AUTOMATA-S2.md',
        'lessons/058-process-models-and-scheduling/ACADEMIC-OS2-S2.md',
        'lessons/001-programs-machines-and-the-terminal/ACADEMIC-CHOICE-SLOTS.md',
        'lessons/100-graduation-capstone/ACADEMIC-S2-MASTER.md',
    ],
}
fail = []
for sem, paths in REQUIRED.items():
    for rel in paths:
        p = ROOT/rel
        if not p.exists():
            fail.append(f'MISSING {sem}: {rel}')
            continue
        text = p.read_text(encoding='utf-8', errors='ignore')
        for term in ('## Exercises', '## Mastery gate'):
            if term not in text and 'MASTER' not in p.name:
                fail.append(f'WEAK {sem}: {rel} lacks {term}')
        if len(text.splitlines()) < 35:
            fail.append(f'TOO THIN {sem}: {rel}')

# Ensure no stale separate academic content tree was created.
for bad in ('academic','practice','exams','training','reference'):
    p=ROOT/bad
    if p.exists():
        fail.append(f'FORBIDDEN ROOT CONTENT TREE: {bad}')

if fail:
    print('\n'.join(fail))
    raise SystemExit(1)
print('LIG YEAR 1 VALIDATION PASSED: all named S1/S2 packs and checkpoints present; no separate academic/practice/exam trees.')
