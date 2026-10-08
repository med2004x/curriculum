"""Student checker for S1 general accounting.

Input format: one CSV-like line per transaction: account,debit,credit
The implementation is deliberately incomplete. Your accounting reasoning is the
primary work; this file only helps check totals after you have solved the case.
"""
from __future__ import annotations
from math import isclose


def total_debit(rows: list[tuple[str, float, float]]) -> float:
    return sum(debit for _, debit, _ in rows)


def total_credit(rows: list[tuple[str, float, float]]) -> float:
    return sum(credit for _, _, credit in rows)


def balanced(rows: list[tuple[str, float, float]]) -> bool:
    return isclose(total_debit(rows), total_credit(rows), rel_tol=1e-9, abs_tol=1e-9)


def main() -> None:
    print("S1 accounting student workspace")
    print("Create your own transaction set, journal, ledger, and statements first.")


if __name__ == "__main__":
    main()
