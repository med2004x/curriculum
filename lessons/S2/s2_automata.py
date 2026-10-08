# S2 COURSE EXERCISE - AUTOMATA
# Study the prompts below, then complete the exercises in this file.

# # FSEGS LIG — Théories des langages et des automates
#
# **Semester:** S2
# **Attached engineering lesson:** 038 — Recursion, Dynamic Programming, and Recurrences
#
# ## Scope
#
# Alphabet, strings and languages; regular languages; deterministic and nondeterministic finite automata; transition functions; accepted language; regular expressions; equivalence intuition; automata construction; simple language recognition; transition-table tracing.
#
# ## Teaching notes
#
# A finite automaton is defined by states, an alphabet, transitions, an initial state, and accepting states. A string is accepted when the transition process ends in an accepting state after consuming the entire input.
#
# Distinguish the language described by a regular expression from the mechanism used to recognise it. For a DFA, exactly one transition is followed for each state/input pair where the transition is defined.
#
# ## Worked examples
#
# 1. Trace a DFA on several strings and identify accepted/rejected inputs.
# 2. Build a DFA recognizing binary strings ending in `01`.
# 3. Write a regular expression for a simple pattern and test it against sample strings.
# 4. Convert a transition table into a diagram.
#
# ## Exercises
#
# 1. Define alphabet/string/language precisely.
# 2. Complete missing DFA transitions.
# 3. Trace ten strings through a supplied automaton.
# 4. Design a DFA for a parity condition.
# 5. Design a DFA for a suffix/prefix pattern.
# 6. Write a regular expression for a simple finite pattern.
# 7. Decide whether a supplied string belongs to a described language.
# 8. Minimise a tiny DFA when the course expects it.
# 9. Build a small automaton simulator from a transition table.
# 10. Give a counterexample when two claimed automata are not equivalent.
#
# ## Coding workspace
#
# s2_automata.py is a student simulator starter. Missing transitions reject a word; the tests check traces and the empty-word case. The formal transition system remains the answer.
#
# ## Mastery gate
#
# Given an unseen language description, define the alphabet, construct an automaton, trace sample strings, and explain acceptance formally.

# --- Executable exercise/checking helpers ---
"""Student finite-automaton simulator starter for S2."""
from __future__ import annotations

TransitionTable = dict[tuple[str, str], str]


def accepts(
    transitions: TransitionTable,
    start: str,
    accepting: set[str],
    word: str,
) -> bool:
    state = start
    for symbol in word:
        next_state = transitions.get((state, symbol))
        if next_state is None:
            return False
        state = next_state
    return state in accepting


def main() -> None:
    print("S2 automata student workspace")


if __name__ == "__main__":
    main()
