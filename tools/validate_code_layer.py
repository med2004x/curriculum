#!/usr/bin/env python3
from pathlib import Path
import re
import hashlib

ROOT = Path(__file__).resolve().parents[1]
lessons = sorted(p for semester in (ROOT / "lessons" / "S1", ROOT / "lessons" / "S2") for p in semester.glob("[0-9][0-9][0-9]-*"))
issues = []
for lesson in lessons:
    if not lesson.is_dir():
        continue
    code = lesson / "code"
    if not code.is_dir():
        issues.append(f"{lesson.name}: missing code/")
        continue
    files = [p for p in code.rglob("*") if p.is_file()]
    student = [p for p in files if p.name in {"exercise.c", "exercise.go", "exercise.rs", "exercise.sql", "challenge.c", "challenge.go", "challenge.rs", "challenge.sql"}]
    for p in student:
        s = p.read_text(errors="ignore")
        if "TODO" not in s and "CHALLENGE TODO" not in s and "t.Skip" not in s and "not implemented" not in s:
            issues.append(f"{lesson.name}: student file looks complete: {p.relative_to(code)}")
        if "gcc " in s and p.suffix == ".c":
            issues.append(f"{lesson.name}: shell command found inside C source: {p.relative_to(code)}")
        if re.search(r"return input\b", s):
            issues.append(f"{lesson.name}: identity stub in student file: {p.relative_to(code)}")
    for p in files:
        if p.suffix in {".c", ".go", ".rs", ".sql", ".py", ".sh"}:
            s=p.read_text(errors="ignore")
            if "starter content intentionally omitted" in s.lower():
                issues.append(f"{lesson.name}: placeholder content: {p.relative_to(code)}")
        if p.name.endswith("_test.go") and re.search(r'\bt\.Skip(?:f)?\s*\(', s):
            issues.append(f"{lesson.name}: skipped Go test is not verification: {p.relative_to(code)}")

    # A challenge must add a distinct task, not just rename the exercise and
    # replace TODO comments. Compare normalized source after removing comments.
    for challenge in (p for p in files if p.name.startswith("challenge.")):
        exercise = code / ("exercise" + challenge.suffix)
        if exercise.exists():
            def normalized(path):
                source = path.read_text(errors="ignore")
                source = re.sub(r'/\*.*?\*/|//[^\n]*|\#[^\n]*', '', source, flags=re.S)
                return re.sub(r'\s+', '', source)
            if normalized(challenge) == normalized(exercise):
                issues.append(f"{lesson.name}: challenge duplicates exercise after comments: {challenge.name}")

    evidence = lesson / "EVIDENCE.md"
    if evidence.exists():
        for ref in re.findall(r'`(code/[^`]+)`', evidence.read_text(errors="replace")):
            if not (lesson / ref).exists():
                issues.append(f"{lesson.name}: dead evidence path: {ref}")
if issues:
    print("code validation FAILED")
    for x in issues: print("-", x)
    raise SystemExit(1)
print(f"code validation passed: {len([x for x in lessons if x.is_dir()])}/100 lessons")
