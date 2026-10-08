#!/usr/bin/env python3
"""Validate semester exercise files stored inside the lessons tree."""
from __future__ import annotations

import ast
from pathlib import Path
import sys

ROOT = Path(__file__).resolve().parents[1]
EXPECTED = {"S1": (11, 5), "S2": (10, 4)}
errors: list[str] = []
for semester, (course_count, test_count) in EXPECTED.items():
    folder = ROOT / "lessons" / semester
    courses = [p for p in folder.glob("*.py") if not p.name.startswith("test_")]
    tests = list(folder.glob("test_*.py"))
    if len(courses) != course_count:
        errors.append(f"{semester}: expected {course_count} course exercise files, found {len(courses)}")
    if len(tests) != test_count:
        errors.append(f"{semester}: expected {test_count} test files, found {len(tests)}")
    for source in courses + tests:
        try:
            ast.parse(source.read_text(encoding="utf-8"), filename=str(source))
        except (SyntaxError, UnicodeError) as exc:
            errors.append(f"{source.relative_to(ROOT)}: invalid Python syntax: {exc}")
    for source in courses:
        text = source.read_text(encoding="utf-8", errors="replace")
        if "## Exercises" not in text and "## Part A" not in text and "## Algorithmique" not in text:
            errors.append(f"{source.relative_to(ROOT)}: course explanation or exercise prompts missing")
        if not text.startswith("#"):
            errors.append(f"{source.relative_to(ROOT)}: course notes must be visible as source comments")

print("Semester folders: 2 (inside lessons/)")
print(f"Course exercise files: {sum(count for count, _ in EXPECTED.values())}")
print(f"Python test files: {sum(count for _, count in EXPECTED.values())}")
print(f"Errors: {len(errors)}")
for error in errors:
    print("ERROR:", error)
sys.exit(1 if errors else 0)
