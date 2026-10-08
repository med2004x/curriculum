"""Student workspace for S1 Statistics and Probability.

No exercise answers are embedded here. The program is an arithmetic verifier for
values and probability cases you have already solved on paper.
"""
from __future__ import annotations

from math import prod
from math import isfinite


def mean(values: list[float]) -> float:
    if not values or any(not isfinite(value) for value in values):
        raise ValueError("values must be a non-empty finite sample")
    return sum(values) / len(values)


def variance_population(values: list[float]) -> float:
    center = mean(values)
    return sum((value - center) ** 2 for value in values) / len(values)


def conditional_probability(intersection: float, event_b: float) -> float:
    if not (0 <= intersection <= event_b <= 1) or event_b == 0:
        raise ValueError("require 0 <= P(A and B) <= P(B) and P(B) > 0")
    return intersection / event_b


def main() -> None:
    print("S1 statistics/probability student workspace")
    print("Use your own datasets and compare the program against your written result.")
    print(f"Sanity arithmetic available: {prod([2, 3, 4])}")


if __name__ == "__main__":
    main()
