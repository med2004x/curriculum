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
