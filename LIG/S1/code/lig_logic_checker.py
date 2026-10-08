"""Student truth-table / logic verifier for S1 systems logic.

The expressions and expected outputs must come from your written solution.
"""
from __future__ import annotations


def implication(p: bool, q: bool) -> bool:
    return (not p) or q


def to_twos_complement(value: int, width: int) -> str:
    """Encode a signed integer as a fixed-width two's-complement bit string."""
    if width <= 0:
        raise ValueError("width must be positive")
    if not -(1 << (width - 1)) <= value < (1 << (width - 1)):
        raise ValueError("value does not fit the signed width")
    return format(value & ((1 << width) - 1), f"0{width}b")


def from_twos_complement(bits: str) -> int:
    """Decode a non-empty fixed-width two's-complement bit string."""
    if not bits or any(bit not in "01" for bit in bits):
        raise ValueError("bits must be a non-empty binary string")
    unsigned = int(bits, 2)
    return unsigned - (1 << len(bits)) if bits[0] == "1" else unsigned


def build_truth_table() -> list[tuple[bool, bool, bool]]:
    rows: list[tuple[bool, bool, bool]] = []
    for p in (False, True):
        for q in (False, True):
            rows.append((p, q, implication(p, q)))
    return rows


def main() -> None:
    for row in build_truth_table():
        print(row)


if __name__ == "__main__":
    main()
