#!/usr/bin/env python3
"""CS50-style exact-output checks for Lesson 004's four logic problems."""
from __future__ import annotations

import shutil
import subprocess
import sys
import tempfile
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
CODE = ROOT / "lessons" / "S1" / "004-conditions-logic-and-proofs" / "code"
TRUTH = """0 0 0 -> 1
0 0 1 -> 0
0 1 0 -> 1
0 1 1 -> 0
1 0 0 -> 1
1 0 1 -> 0
1 1 0 -> 1
1 1 1 -> 1
"""


def compile_one(compiler: str, source: Path, output: Path) -> bool:
    result = subprocess.run([compiler, "-std=c17", "-Wall", "-Wextra", "-Wpedantic", str(source), "-o", str(output)], text=True, capture_output=True)
    if result.returncode:
        print(f"FAIL: {source.name} does not compile.")
        print(result.stderr, end="", file=sys.stderr)
        return False
    return True


def check(executable: Path, args: list[str], expected: str, label: str) -> bool:
    result = subprocess.run([str(executable), *args], text=True, capture_output=True)
    if result.returncode != 0 or result.stdout != expected:
        print(f"FAIL: {label}; expected {expected!r}, got {result.stdout!r} (exit {result.returncode}).")
        return False
    print(f"PASS: {label}")
    return True


def main() -> int:
    compiler = shutil.which("gcc") or shutil.which("clang")
    if not compiler:
        print("SETUP ERROR: install GCC or Clang and its standard headers.", file=sys.stderr)
        return 2
    passed = True
    with tempfile.TemporaryDirectory(prefix="dart-lesson004-") as temp:
        temp_path = Path(temp)
        for filename in ("truth.c", "demorgan.c", "classify.c", "leap.c"):
            exe = temp_path / (Path(filename).stem + (".exe" if sys.platform == "win32" else ""))
            if not compile_one(compiler, CODE / filename, exe):
                passed = False
                continue
            if filename == "truth.c":
                passed &= check(exe, [], TRUTH, "all eight truth-table rows")
            elif filename == "demorgan.c":
                passed &= check(exe, [], "ok\n", "both De Morgan laws for all Boolean pairs")
            elif filename == "classify.c":
                for value, expected in (("-3", "negative odd\n"), ("0", "zero even\n"), ("4", "positive even\n")):
                    passed &= check(exe, [value], expected, f"classify {value}")
            elif filename == "leap.c":
                expected = "1900: common\n2000: leap\n2023: common\n2024: leap\n"
                passed &= check(exe, [], expected, "century and ordinary leap-year boundaries")
    if not passed:
        print("Fix the first failing contract, then rerun this command.")
        return 1
    print("PASS: all Lesson 004 contracts.")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
