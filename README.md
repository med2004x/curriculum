# LIG S1/S2 + Systems Engineering Curriculum

This repository now has two clear study routes:

- [LIG Semester 1](LIG/S1/README.md) ? subject packs, semester-specific practice code, engineering links, and the S1 mastery rehearsal.
- [LIG Semester 2](LIG/S2/README.md) ? its own subject packs, practice code, engineering links, and the S2 mastery rehearsal.
- [Full engineering spine](CURRICULUM.md) ? 100 sequential lessons from C foundations through production systems.

The LIG tracks are organized by university semester. The engineering lessons remain a shared supporting sequence and may be referenced by both semesters; they are not mechanically split at lesson 050.

## Start

If you are preparing for LIG, choose the correct [semester](LIG/README.md), then use the course pack's exercises and mastery gate. If you are studying systems/backend engineering, start with [Lesson 001](lessons/001-programs-machines-and-the-terminal/README.md).

## Feedback and assessment

Use the problem statements and supplied checks from the linked upstream labs for xv6, MIT 6.5840, and BusTub. Local lesson tests are intended to fail until the student implementation meets the contract. Record commands and outputs in `EVIDENCE.md`; a structural validator PASS only checks repository integrity.

## Maintainer references

- [LIG track map](LIG/README.md)
- [100-lesson index](LESSON-INDEX.md)
- [Course design and assessment](CURRICULUM.md)
- [Code artifact rules](CODE-ARTIFACT-RULES.md)
- [Code audit](CODE-AUDIT.md)
- [External course sources](SOURCES.md)
- [Course validators](tools/)

Run `python tools/validate_all.py` to check repository structure and local links. Run `python tools/run_lig_tests.py s1`, `s2`, or `all` to execute the subject-code checks.
