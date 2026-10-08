#!/usr/bin/env python3
"""Run all structural curriculum validators in a consistent order."""
from __future__ import annotations

from pathlib import Path
import subprocess
import sys

ROOT = Path(__file__).resolve().parents[1]
CHECKS = [
    "validate_course.py",
    "validate_code_layer.py",
    "validate_code_artifacts.py",
    "validate_lig_curriculum.py",
]


def main() -> int:
    failed = False
    for check in CHECKS:
        print(f"\n=== {check} ===", flush=True)
        result = subprocess.run([sys.executable, str(ROOT / "tools" / check)], cwd=ROOT, check=False)
        failed |= result.returncode != 0
    return 1 if failed else 0


if __name__ == "__main__":
    raise SystemExit(main())
