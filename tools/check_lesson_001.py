#!/usr/bin/env python3
"""Compile Lesson 001 task files and check their stated output contracts."""
from __future__ import annotations

import shutil
import subprocess
import sys
import tempfile
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
CODE = ROOT / "lessons" / "S1" / "001-programs-machines-and-the-terminal" / "code"
PROGRAMS = {
    "task_01.c": "hello\n",
    "hello.c": "Hello, world!\n",
    "greet.c": "Hello, Ada!\n",
    "report.c": "Name: Ada\nScore: 42\n",
    "arith.c": "66\n",
    "exercise.c": "Hello, DART!\n",
}


def main() -> int:
    compiler = shutil.which("gcc") or shutil.which("clang")
    if not compiler:
        print("SETUP ERROR: no gcc or clang found on PATH. Follow SETUP.md.", file=sys.stderr)
        return 2
    failed = False
    with tempfile.TemporaryDirectory(prefix="dart-lesson001-") as temp:
        for filename, expected in PROGRAMS.items():
            source = CODE / filename
            executable = Path(temp) / (source.stem + (".exe" if sys.platform == "win32" else ""))
            build = subprocess.run(
                [compiler, "-std=c17", "-Wall", "-Wextra", "-Wpedantic", str(source), "-o", str(executable)],
                text=True, capture_output=True, check=False,
            )
            if build.returncode:
                print(f"BUILD FAIL: {filename} did not compile.")
                if build.stderr:
                    print(build.stderr, end="", file=sys.stderr)
                failed = True
                continue
            run = subprocess.run([str(executable)], text=True, capture_output=True, check=False)
            if run.returncode != 0 or run.stdout != expected:
                print(f"FAIL: {filename}: expected output {expected!r} and exit status 0.")
                print(f"Received output {run.stdout!r}, exit status {run.returncode}.")
                failed = True
            else:
                print(f"PASS: {filename}")
    if failed:
        print("Complete the lesson tasks, then run this check again.")
        print("If the compiler or standard headers are missing, follow SETUP.md.", file=sys.stderr)
        return 1
    print("PASS: all six Lesson 001 programs meet their output contracts.")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
