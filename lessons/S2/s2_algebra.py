# S2 COURSE EXERCISE - ALGEBRA
# Study the prompts below, then complete the exercises in this file.

# # FSEGS LIG — Algèbre
#
# **Semester:** S2
# **Attached engineering lesson:** 037 — Trees, Graphs, and Discrete Structure
#
# ## Scope
#
# Vectors; linear combinations; span; subspaces; linear dependence/independence; matrices; operations; determinants; systems of linear equations; Gaussian elimination; rank; linear maps; kernel/image; basis; dimension.
#
# ## Teaching notes
#
# A vector can be treated as a coordinate object, but the important question is which space and which operations are being used. A linear combination is built by scalar multiplication and vector addition.
#
# For a matrix system `Ax=b`, row operations preserve the solution set when applied correctly. Rank measures the number of independent constraints/columns or rows, depending on the viewpoint.
#
# ## Worked examples
#
# 1. Determine whether three vectors are linearly independent.
# 2. Solve a 3×3 linear system by Gaussian elimination and verify by substitution.
# 3. Compute a determinant and interpret singular versus invertible.
# 4. Find the kernel of a simple linear transformation.
#
# ## Exercises
#
# 1. Compute vector sums/scalar multiples.
# 2. Express a vector as a linear combination of others or prove impossible.
# 3. Test linear independence.
# 4. Perform matrix multiplication by hand.
# 5. Compute determinants of 2×2 and 3×3 matrices.
# 6. Solve systems by elimination.
# 7. Determine rank from row-echelon form.
# 8. Find a basis for a given subspace.
# 9. Compute kernel/image for a small linear map.
# 10. Interpret a singular system in terms of existence/uniqueness of solutions.
#
# ## Coding workspace
#
# s2_algebra.py is a student workspace for matrix multiplication and rank. Write the row operations in your work notes first, then use `python tools/run_semester_tests.py s2` for the included checks.
#
# ## Mastery gate
#
# Solve a fresh matrix/system problem, one independence problem, and one kernel/image problem with all row-operation reasoning visible.

# --- Executable exercise/checking helpers ---
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
