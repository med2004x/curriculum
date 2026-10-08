"""Student workspace for S1 Analyse.

This is intentionally incomplete. Write the mathematics in Markdown first, then use
this file only to verify arithmetic for your own solved cases.
"""
from __future__ import annotations


def derivative_polynomial(coefficients: list[float]) -> list[float]:
    """Return derivative coefficients, highest power first."""
    if not coefficients:
        raise ValueError("a polynomial needs at least one coefficient")
    degree = len(coefficients) - 1
    if degree == 0:
        return [0]
    return [coefficient * (degree - index) for index, coefficient in enumerate(coefficients[:-1])]


def evaluate_polynomial(coefficients: list[float], x: float) -> float:
    """Evaluate a polynomial represented by highest-power-first coefficients."""
    if not coefficients:
        raise ValueError("a polynomial needs at least one coefficient")
    value = 0
    for coefficient in coefficients:
        value = value * x + coefficient
    return value


def main() -> None:
    print("S1 Analyse student workspace")
    print("Write and solve one exercise before implementing these helpers.")


if __name__ == "__main__":
    main()
