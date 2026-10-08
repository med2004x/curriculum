#!/usr/bin/env python3
"""Check that semester packs, code workspaces, and local Markdown links agree."""
from __future__ import annotations

from pathlib import Path
import re
import sys
import ast

ROOT = Path(__file__).resolve().parents[1]
errors: list[str] = []
expected = {
    "S1": {"ACCOUNTING.md", "ADS1.md", "ANALYSIS.md", "BUSINESS-COMMUNICATION.md", "DB-MULTIMEDIA.md", "DIGITAL-SKILLS.md", "LOGIC-ARCH.md", "MANAGEMENT.md", "OS1.md", "S1-MASTER.md", "STATS-PROB.md", "README.md"},
    "S2": {"ACCOUNTING-IS.md", "ADS2.md", "ALGEBRA.md", "AUTOMATA.md", "BUSINESS-COMMUNICATION.md", "FINANCE.md", "IS-NETWORKS.md", "MATH-LOGIC.md", "OS2.md", "S2-MASTER.md", "README.md"},
}

for semester, required in expected.items():
    directory = ROOT / "LIG" / semester
    actual = {path.name for path in directory.glob("*.md")}
    missing = required - actual
    if missing:
        errors.append(f"{semester}: missing packs: {', '.join(sorted(missing))}")
    for doc in directory.rglob("*.md"):
        text = doc.read_text(encoding="utf-8", errors="replace")
        for target in re.findall(r"\]\(([^)]+)\)", text):
            if target.startswith(("https://", "http://", "mailto:")) or target.startswith("#"):
                continue
            local = (doc.parent / target.split("#", 1)[0]).resolve()
            if not local.exists():
                errors.append(f"{doc.relative_to(ROOT)}: dead link {target}")
    if not (directory / "code").is_dir():
        errors.append(f"{semester}: missing code workspace")
    if not list((directory / "code").glob("test_*.py")):
        errors.append(f"{semester}: no executable practice tests")
    for source in (directory / "code").glob("*.py"):
        try:
            ast.parse(source.read_text(encoding="utf-8"), filename=str(source))
        except (SyntaxError, UnicodeError) as exc:
            errors.append(f"{source.relative_to(ROOT)}: invalid Python syntax: {exc}")

if not (ROOT / "LIG" / "README.md").is_file():
    errors.append("LIG: missing top-level README.md semester index")

for doc in (ROOT / "LIG").rglob("*.md"):
    text = doc.read_text(encoding="utf-8", errors="replace")
    for target in re.findall(r"\]\(([^)]+)\)", text):
        if target.startswith(("https://", "http://", "mailto:")) or target.startswith("#"):
            continue
        local = (doc.parent / target.split("#", 1)[0]).resolve()
        if not local.exists():
            errors.append(f"{doc.relative_to(ROOT)}: dead link {target}")

for stub in (ROOT / "lessons").glob("*/ACADEMIC-*.md"):
    text = stub.read_text(encoding="utf-8", errors="replace")
    for target in re.findall(r"\]\(([^)]+)\)", text):
        if not (stub.parent / target).resolve().exists():
            errors.append(f"{stub.relative_to(ROOT)}: dead redirect {target}")

print(f"LIG tracks: {len(expected)}")
print(f"subject packs: {sum(len(items) - 1 for items in expected.values())}")
print(f"practice tests: {sum(1 for p in (ROOT / 'LIG').rglob('test_*.py'))}")
print(f"errors: {len(errors)}")
for error in errors:
    print("ERROR:", error)
sys.exit(1 if errors else 0)
