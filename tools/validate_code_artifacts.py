#!/usr/bin/env python3
"""Validate that lesson coding artifacts actually exist and basic source checks pass."""
from pathlib import Path
import re
import subprocess
import sys
import shutil

ROOT = Path(__file__).resolve().parents[1]
LESSONS = ROOT / "lessons"
EXTS = {".c", ".h", ".go", ".rs", ".py", ".sql", ".sh", ".asm", ".s", ".json"}
LOCAL_FILE = re.compile(r"`([^`]+\\.(?:c|h|go|rs|py|sql|sh|asm|s|json|toml|mod|yaml|yml))`", re.I)

failures = []
lesson_dirs = sorted(LESSONS.glob("[0-9][0-9][0-9]-*"))
if len(lesson_dirs) != 100:
    failures.append(f"expected 100 lesson directories, found {len(lesson_dirs)}")

for lesson in lesson_dirs:
    for required in ("README.md", "WORK.md", "EVIDENCE.md", "code"):
        if not (lesson / required).exists():
            failures.append(f"{lesson.name}: missing {required}")

    text = "\n".join((lesson / name).read_text(errors="replace") for name in ("README.md", "WORK.md"))
    for filename in ("README.md", "WORK.md", "EVIDENCE.md"):
        doc = lesson / filename
        if not doc.exists():
            continue
        doc_text = doc.read_text(errors="replace")
        for ref in re.findall(r"`(code/[^`]+)`", doc_text):
            if not (lesson / ref).exists():
                failures.append(f"{lesson.name}: referenced code path missing: {ref}")
    for ref in sorted(set(LOCAL_FILE.findall(text))):
        if "/" in ref or "\\" in ref:
            continue
        if not list(lesson.rglob(ref)):
            failures.append(f"{lesson.name}: referenced code file missing: {ref}")

    exercise = list((lesson / "code").glob("exercise.*")) + list((lesson / "code").glob("exercise/main.*"))
    if not exercise:
        failures.append(f"{lesson.name}: no exercise starter under code/")

    if "[starter content intentionally omitted]" in text:
        failures.append(f"{lesson.name}: omitted starter content marker remains")

# Syntax checks for generated lesson starters. Full project builds are not assumed for PostgreSQL/network/system dependencies.
if shutil.which("gcc"):
    for src in sorted(LESSONS.glob("*/code/exercise.c")):
        r = subprocess.run(["gcc", "-std=c17", "-Wall", "-Wextra", "-Wpedantic", "-fsyntax-only", str(src)], capture_output=True, text=True)
        if r.returncode:
            failures.append(f"C starter does not compile: {src}: {r.stderr.strip()}")
else:
    print("C syntax checks: SKIPPED (gcc is not installed)")
print("Go behavior tests: run per lesson; structural validation does not treat compilation as a test")

code_files = [p for p in LESSONS.rglob("*") if p.is_file() and (p.suffix.lower() in EXTS or p.name in {"go.mod", "go.sum", "Cargo.toml", "Makefile"})]
print(f"lessons={len(lesson_dirs)}")
print(f"coding_files={len(code_files)}")
print(f"failures={len(failures)}")
for failure in failures:
    print(f"FAIL: {failure}")
sys.exit(1 if failures else 0)
