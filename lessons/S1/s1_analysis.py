# S1 COURSE EXERCISE - ANALYSIS
# Study the prompts below, then complete the exercises in this file.

# # FSEGS LIG — Analyse 1
#
# **Semester:** S1  
# **Attached engineering lesson:** 006 — Functions, the Call Stack, and Debugging
#
# ## Scope
#
# Real functions; domain and codomain; elementary function families; limits; continuity; derivatives; derivative rules; monotonicity; extrema; tangent line; basic curve study; introductory antiderivative/integration ideas where included by the section.
#
# ## Teaching sequence
#
# 1. Function notation and reading graphs.
# 2. Domain restrictions and composition.
# 3. Limits and one-sided limits.
# 4. Continuity and discontinuities.
# 5. Derivative as rate of change and tangent slope.
# 6. Product, quotient, chain and elementary derivatives.
# 7. Monotonicity and extrema through the sign of the derivative.
# 8. Structured curve study: domain → limits → derivative → sign table → variation table.
#
# ## Worked examples
#
# ### Example A — Domain
#
# For `f(x)=sqrt(2x-1)/(x-3)`, identify every restriction before simplifying anything.
#
# ### Example B — Derivative and monotonicity
#
# For `f(x)=x^3-3x²+2`, compute `f'`, factor it, create a sign table, and locate local extrema.
#
# ### Example C — Limit
#
# Evaluate a rational limit with a removable indeterminate form by factoring before substitution.
#
# ## Coding workspace
#
# s1_analysis.py is a student starter for polynomial differentiation and evaluation. Run the S1 tests with `python tools/run_semester_tests.py s1`; the tests specify the helper contracts, not the full calculus syllabus.
#
# ## Exercises
#
# 1. Find domains of rational, radical, logarithmic, and composed functions.
# 2. Compute five standard limits, including one one-sided limit.
# 3. Determine continuity at a point for a piecewise function.
# 4. Differentiate polynomial, rational, exponential/logarithmic, and composed functions.
# 5. Find tangent equations at specified points.
# 6. Build derivative sign tables and monotonicity intervals.
# 7. Locate local extrema and interpret them in context.
# 8. Complete a curve-study table from a given derivative graph.
# 9. Solve an optimization problem with a clearly stated physical/business constraint.
# 10. Repeat one missed problem with one changed parameter.
#
# ## Problem-solving lab
#
# Solve one complete curve study from blank, including domain, limits, derivative, sign table, extrema, and final variation table. Then write a five-sentence interpretation of the result.
#
# ## Exam patterns
#
# Expect direct calculation plus structured reasoning. Marks often come from the method: stating restrictions, writing the derivative correctly, organizing sign tables, and interpreting a result.
#
# ## Mastery gate
#
# Complete two full curve studies, five derivative calculations, and three limit problems from fresh data without notes. Show every intermediate step.

# --- Executable exercise/checking helpers ---
"""Student workspace for S1 Analyse.

This is intentionally incomplete. Write the mathematics in your work notes first, then use
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
