# Sources and External References

The learner path is local; external labs are linked rather than vendored. This file collects the external references that are useful for exact semantics, additional practice, and curriculum design.

## Go

- Go documentation: https://go.dev/doc/
- A Tour of Go: https://go.dev/tour/
- Go by Example: https://gobyexample.com/
- Exercism Go: https://exercism.org/tracks/go

## Curriculum design

- Udemy course outlining and backward design: https://teach.udemy.com/course-creation/outline-your-course/
- Udemy practice-activity quality standards: https://support.udemy.com/hc/en-us/articles/229605248-Practice-activities-Quality-standards
- Coursera mastery learning and formative feedback: https://blog.coursera.org/how-to-integrate-mastery-learning-into-course-design/
- Coursera learning-objective and assessment alignment: https://blog.coursera.org/courseras-commitment-to-learning-skill-development/
- freeCodeCamp curriculum structure: https://contribute.freecodecamp.org/how-to-work-on-coding-challenges/

## How to use external material

For lessons 064, 085, and 090, the official course project is the actual lab assignment. For other topics, external exercises are supplementary: use them for additional repetitions, then return to the lesson's transfer task.

## Course labs and feedback systems

These primary course repositories are the assessment references for the DART lessons listed below. Use the handout and tests from the same release; do not mix versions.

- CS50 problem specifications: https://github.com/cs50/problems and check suites: https://github.com/cs50/checks
- MIT 6.1810 system-call lab and tests: https://pdos.csail.mit.edu/6.S081/2026/labs/syscall.html; lab repo: `git://g.csail.mit.edu/xv6-labs-2026` (lesson 064)
- MIT 6.5840 distributed systems labs: https://pdos.csail.mit.edu/6.5840/ (source checkout: git://g.csail.mit.edu/6.5840-golabs-2026) (lesson 090)
- CMU BusTub database project and tests: https://github.com/cmu-db/bustub (lesson 085)

DART's lessons 064, 085, and 090 direct learners to these upstream assignments and ask them to record the selected commit and test output in `EVIDENCE.md`. The structural validators do not run these external test suites. For smaller lessons, use the per-lesson test file or checker where one is present; an unfinished starter is not considered passed merely because it compiles.


## Semester course exercises

S1 and S2 course exercises are stored in `lessons/S1/` and `lessons/S2/`. Teaching notes and exercise prompts are comments at the top of each exercise file. The current official FSEGS maquette and selected elective titles were not supplied with this repository; confirm the subject list and exam scope with your instructor.

The `dart-upstreams` directory is a local audit/reference checkout outside this repository. It is not included in Git. The lesson instructions identify the upstream repository and ask learners to record the exact commit they use.
