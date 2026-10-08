# FSEGS LIG — Algèbre

**Semester:** S2  
**Attached engineering lesson:** 037 — Trees, Graphs, and Discrete Structure

## Scope

Vectors; linear combinations; span; subspaces; linear dependence/independence; matrices; operations; determinants; systems of linear equations; Gaussian elimination; rank; linear maps; kernel/image; basis; dimension.

## Teaching notes

A vector can be treated as a coordinate object, but the important question is which space and which operations are being used. A linear combination is built by scalar multiplication and vector addition.

For a matrix system `Ax=b`, row operations preserve the solution set when applied correctly. Rank measures the number of independent constraints/columns or rows, depending on the viewpoint.

## Worked examples

1. Determine whether three vectors are linearly independent.
2. Solve a 3×3 linear system by Gaussian elimination and verify by substitution.
3. Compute a determinant and interpret singular versus invertible.
4. Find the kernel of a simple linear transformation.

## Exercises

1. Compute vector sums/scalar multiples.
2. Express a vector as a linear combination of others or prove impossible.
3. Test linear independence.
4. Perform matrix multiplication by hand.
5. Compute determinants of 2×2 and 3×3 matrices.
6. Solve systems by elimination.
7. Determine rank from row-echelon form.
8. Find a basis for a given subspace.
9. Compute kernel/image for a small linear map.
10. Interpret a singular system in terms of existence/uniqueness of solutions.

## Coding workspace

[`lig_algebra_s2.py`](code/lig_algebra_s2.py) is a student workspace for matrix multiplication and rank. Write the row operations in Markdown first, then use `python tools/run_lig_tests.py s2` for the included checks.

## Mastery gate

Solve a fresh matrix/system problem, one independence problem, and one kernel/image problem with all row-operation reasoning visible.
