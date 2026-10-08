# FSEGS LIG — Gestion financière

**Semester:** S2  
**Attached engineering lesson:** 059 — Queues, Scheduling, and Throughput

## Scope

Financial decision vocabulary; interest; simple and compound capitalisation; discounting; present/future value; effective rates; annuity basics; financing versus investment; cash-flow timeline; introductory investment criteria.

## Formula discipline

Use a timeline before writing formulas. Distinguish nominal rate, period rate, number of periods, present value, and future value. Keep the unit of the rate aligned with the unit of time.

Simple interest: `FV = PV(1 + i n)`.

Compound interest: `FV = PV(1+i)^n`.

Discounting: `PV = FV/(1+i)^n`.

For an ordinary annuity, the payment timing must match the formula used. Never import an annuity formula without drawing the cash-flow timeline.

## Worked examples

1. Compare simple and compound growth over several periods.
2. Discount a future payment to present value.
3. Compute the future value of equal periodic deposits.
4. Compare two financing offers using a consistent periodic rate.

## Exercises

1. Solve simple-interest problems with different periods.
2. Solve compound-interest problems and convert rates correctly.
3. Compute present value for future cash flows.
4. Build a payment schedule.
5. Solve an annuity problem from a timeline.
6. Compare financing proposals.
7. Compute a simple investment return from supplied cash flows.
8. Explain the difference between accounting profit and cash flow.
9. Analyse the effect of changing rate or duration.
10. Verify all arithmetic using [`lig_finance_s2.py`](code/lig_finance_s2.py) only after writing the formula manually.

## Exam patterns

Rate conversion, timeline construction, formula selection, PV/FV calculation, annuity calculation, comparison of financial alternatives, interpretation.

## Mastery gate

Solve a mixed set without looking up formulas. For each question, identify the timeline and units before calculating.
