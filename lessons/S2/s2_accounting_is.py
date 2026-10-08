# S2 COURSE EXERCISE - ACCOUNTING-IS
# Study the prompts below, then complete the exercises in this file.

# # FSEGS LIG — Système d'information comptable
#
# **Semester:** S2
# **Attached engineering lesson:** 050 — CRUD, Transactions, and Safe State
#
# ## Scope
#
# Business transaction flow; source documents; transaction capture; validation; journals/ledgers; master data; controls; audit trail; segregation of duties; information flow; reports; reconciliation; integrity between operational and accounting systems.
#
# ## Teaching notes
#
# An accounting information system is an information-flow system, not merely accounting software. Start with the business event, identify the source evidence, validate it, record it, transform it, and report it.
#
# Controls can be preventive, detective, or corrective. Audit trails make changes traceable. Segregation of duties reduces the risk that one person can initiate, approve, execute, and conceal a transaction.
#
# ## Exercises
#
# 1. Draw the transaction flow for a sale from order to accounting record.
# 2. Identify source documents for purchase, sale, payment, and receipt.
# 3. Design validation rules for customer, invoice, and payment records.
# 4. Classify controls as preventive/detective/corrective.
# 5. Identify segregation-of-duties conflicts in three cases.
# 6. Build a reconciliation between operational totals and accounting totals.
# 7. Design an audit trail for edits to an invoice.
# 8. Identify missing data in a transaction-processing system.
# 9. Explain why a database transaction can support accounting integrity.
# 10. Design a small schema linking business events to accounting entries.
#
# ## Coding workspace
#
# s2_accounting_is.py demonstrates one explicit validation contract: event ID, finite non-negative amount, source document, and approver. Treat these as a practice-case policy, not an official accounting rule. Reconciliation returns operational total minus accounting total.
#
# ## Problem-solving lab
#
# Given a scenario where sales totals and accounting totals differ, reconstruct the information flow, enumerate possible failure points, and design a reconciliation procedure.
#
# ## Mastery gate
#
# Model a fresh business process end-to-end and identify actors, data, controls, evidence, and reports without hand-waving.

# --- Executable exercise/checking helpers ---
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
