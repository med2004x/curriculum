"""Student workspace for S2 Accounting Information Systems."""
from __future__ import annotations

from math import isfinite


def validate_transaction(event: dict) -> tuple[bool, list[str]]:
    errors = []
    event_id = event.get("event_id")
    amount = event.get("amount")
    source = event.get("source_document")
    approver = event.get("approved_by")
    if not isinstance(event_id, str) or not event_id.strip():
        errors.append("event_id is required")
    if not isinstance(amount, (int, float)) or not isfinite(amount) or amount < 0:
        errors.append("amount must be a finite non-negative number")
    if not isinstance(source, str) or not source.strip():
        errors.append("source_document is required")
    if not isinstance(approver, str) or not approver.strip():
        errors.append("approved_by is required")
    return not errors, errors


def reconcile(operational_total: float, accounting_total: float) -> float:
    return operational_total - accounting_total


def main() -> None:
    print("S2 accounting information systems student workspace")


if __name__ == "__main__":
    main()
