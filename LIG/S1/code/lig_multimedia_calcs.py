"""Student arithmetic verifier for multimedia calculations.

Write units and formulas in the academic workbook before using this file.
"""
from __future__ import annotations

from math import isfinite


def raw_rgb_bytes(width: int, height: int, bytes_per_pixel: int = 3) -> int:
    if width <= 0 or height <= 0 or bytes_per_pixel <= 0:
        raise ValueError("dimensions and bytes_per_pixel must be positive")
    return width * height * bytes_per_pixel


def pcm_bytes(seconds: float, sample_rate: int, channels: int, bytes_per_sample: int) -> int:
    if not isfinite(seconds) or seconds < 0:
        raise ValueError("seconds must be finite and non-negative")
    if sample_rate <= 0 or channels <= 0 or bytes_per_sample <= 0:
        raise ValueError("sample rate, channels, and bytes per sample must be positive")
    return int(seconds * sample_rate * channels * bytes_per_sample)


def main() -> None:
    print("S1 multimedia calculation verifier")
    print("Enter your own problem values after solving them by hand.")


if __name__ == "__main__":
    main()
