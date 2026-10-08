# FSEGS LIG — Statistiques et probabilités

**Semester:** S1  
**Attached engineering lesson:** 019 — Searching, Sorting, and Measuring Work

## Scope

Population/sample; variable types; frequency and cumulative frequency; mean; median; mode; quartiles/percentiles; variance; standard deviation; coefficient of variation; tables/graphs; basic counting; probability axioms; conditional probability; independence; total probability; Bayes; random variables; expectation and variance; common discrete models at introductory level.

## Teaching notes

Always identify the population, sample, variable, and unit before calculating a statistic. Keep the convention for variance explicit when a course or exercise specifies it.

For grouped data, keep class boundaries, class frequencies, cumulative frequencies, and representative values distinct. For probability, distinguish `P(A∩B)`, `P(A∪B)`, and `P(A|B)`.

Independence means `P(A∩B)=P(A)P(B)`, not merely “the events seem unrelated.”

## Worked examples

1. Compute mean, median, quartiles, variance and standard deviation for `[2,3,3,4,8]` and explain the effect of the outlier.
2. A device fails with probability 0.1 independently on each trial. Compute the probability of exactly two failures in five trials.
3. Given `P(A)=0.4`, `P(B)=0.5`, `P(A∩B)=0.2`, test independence and compute `P(A|B)`.

## Exercises

1. Classify ten variables as qualitative/quantitative and discrete/continuous.
2. Build a frequency and cumulative-frequency table from raw observations.
3. Compute weighted mean from grouped observations.
4. Compute median and quartiles for odd and even sample sizes.
5. Compare two data sets using mean and standard deviation.
6. Solve counting problems with permutations/combinations.
7. Solve conditional-probability and Bayes exercises.
8. Determine whether two events are independent.
9. Work with a discrete random variable table and compute expectation and variance.
10. Simulate a small probability model in Python and compare empirical/theoretical results.

## Statistical lab

Use [`lig_stats_prob.py`](code/lig_stats_prob.py) to verify arithmetic after the written solution. These functions check the mean, population variance, and conditional-probability arithmetic after you solve by hand. The unit tests do not replace grouped-data, Bayes, or random-variable problems in this pack.

## Exam patterns

Read a table correctly; choose the appropriate formula; calculate cleanly; interpret the number; justify probability transformations; recognize independence; show counting logic.

## Mastery gate

Complete a full descriptive-statistics table, a conditional-probability/Bayes problem, and a random-variable exercise from fresh data. Explain each formula selection before calculating.
