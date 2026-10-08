#!/usr/bin/env python3
"""Run the semester-specific LIG practice checks."""
from __future__ import annotations

import os
from pathlib import Path
import subprocess
import sys

ROOT = Path(__file__).resolve().parents[1]
TRACKS = {"s1": ROOT / "LIG" / "S1" / "code", "s2": ROOT / "LIG" / "S2" / "code"}


def main() -> int:
    selected = sys.argv[1].lower() if len(sys.argv) > 1 else "all"
    tracks = list(TRACKS) if selected == "all" else [selected]
    if any(track not in TRACKS for track in tracks):
        print("Usage: python tools/run_lig_tests.py [s1|s2|all]", file=sys.stderr)
        return 2
    env = os.environ.copy()
    env["PYTHONDONTWRITEBYTECODE"] = "1"
    outcomes = []
    for track in tracks:
        code_dir = TRACKS[track]
        command = [sys.executable, "-m", "unittest", "discover", "-s", str(code_dir), "-p", "test_*.py"]
        print(f"\n[{track.upper()}] {' '.join(command)}", flush=True)
        result = subprocess.run(command, cwd=ROOT, env=env, check=False)
        outcomes.append((track, result.returncode))
    for track, code in outcomes:
        print(f"{track.upper()}: {'PASS' if code == 0 else 'FAIL'}")
    return int(any(code != 0 for _, code in outcomes))


if __name__ == "__main__":
    raise SystemExit(main())
