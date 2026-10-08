"""Student arithmetic workspace for S2 Algebra."""
from __future__ import annotations


def matmul(a: list[list[float]], b: list[list[float]]) -> list[list[float]]:
    def width(matrix: list[list[float]]) -> int:
        if not matrix:
            return 0
        expected = len(matrix[0])
        if any(len(row) != expected for row in matrix):
            raise ValueError("matrix rows must have equal length")
        return expected

    a_width, b_width = width(a), width(b)
    if a_width != len(b):
        raise ValueError("incompatible matrix dimensions")
    return [[sum(a[i][k] * b[k][j] for k in range(a_width))
             for j in range(b_width)] for i in range(len(a))]


def rank_placeholder(matrix: list[list[float]]) -> int:
    if not matrix:
        return 0
    columns = len(matrix[0])
    if any(len(row) != columns for row in matrix):
        raise ValueError("matrix rows must have equal length")
    work = [list(map(float, row)) for row in matrix]
    pivot_row = 0
    for column in range(columns):
        pivot = next((r for r in range(pivot_row, len(work)) if abs(work[r][column]) > 1e-12), None)
        if pivot is None:
            continue
        work[pivot_row], work[pivot] = work[pivot], work[pivot_row]
        pivot_value = work[pivot_row][column]
        work[pivot_row] = [value / pivot_value for value in work[pivot_row]]
        for row in range(len(work)):
            if row != pivot_row:
                factor = work[row][column]
                work[row] = [value - factor * pivot_value_item
                             for value, pivot_value_item in zip(work[row], work[pivot_row])]
        pivot_row += 1
        if pivot_row == len(work):
            break
    return pivot_row


def main() -> None:
    print("S2 algebra student workspace")


if __name__ == "__main__":
    main()
