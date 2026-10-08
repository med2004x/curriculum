#!/usr/bin/env python3
"""Compile Lesson 001's student file and check its exact output."""
from __future__ import annotations

import shutil
import subprocess
import sys
import tempfile
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
SOURCE = ROOT / "lessons" / "001-programs-machines-and-the-terminal" / "code" / "exercise.c"
EXPECTED = "Hello, DART!\n"


def main() -> int:
    compiler = shutil.which("gcc") or shutil.which("clang")
    if not compiler:
        print("SETUP ERROR: no gcc or clang found on PATH. Follow SETUP.md.", file=sys.stderr)
        return 2
    with tempfile.TemporaryDirectory(prefix="dart-lesson001-") as temp:
        executable = Path(temp) / ("exercise.exe" if sys.platform == "win32" else "exercise")
        build = subprocess.run(
            [compiler, "-std=c17", "-Wall", "-Wextra", "-Wpedantic", str(SOURCE), "-o", str(executable)],
            text=True, capture_output=True, check=False,
        )
        if build.returncode:
            print("BUILD FAIL: exercise.c did not compile.")
            if build.stdout:
                print(build.stdout, end="")
            if build.stderr:
                print(build.stderr, end="", file=sys.stderr)
            print("Check that your C compiler includes the standard headers. See SETUP.md.", file=sys.stderr)
            return 1
        run = subprocess.run([str(executable)], text=True, capture_output=True, check=False)
        if run.returncode != 0:
            print(f"RUN FAIL: program exited with {run.returncode}; expected 0.")
            if run.stderr:
                print(run.stderr, end="", file=sys.stderr)
            return 1
        if run.stdout != EXPECTED:
            print("OUTPUT FAIL: expected exactly 'Hello, DART!\n'.")
            print(f"Received: {run.stdout!r}")
            return 1
    print("PASS: exercise.c compiled, printed Hello, DART!, and exited with 0.")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
