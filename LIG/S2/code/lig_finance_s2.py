"""Student workspace for S2 Gestion financière.

Always build the cash-flow timeline first.
"""
from __future__ import annotations


def future_value_compound(present_value: float, periodic_rate: float, periods: int) -> float:
    if present_value < 0 or periodic_rate <= -1 or periods < 0:
        raise ValueError("require PV >= 0, rate > -1, and periods >= 0")
    return present_value * (1 + periodic_rate) ** periods


def present_value(future_value: float, periodic_rate: float, periods: int) -> float:
    if future_value < 0 or periodic_rate <= -1 or periods < 0:
        raise ValueError("require FV >= 0, rate > -1, and periods >= 0")
    return future_value / (1 + periodic_rate) ** periods


def main() -> None:
    print("S2 finance student workspace")


if __name__ == "__main__":
    main()
