# FSEGS LIG — Système d'information comptable

**Semester:** S2  
**Attached engineering lesson:** 050 — CRUD, Transactions, and Safe State

## Scope

Business transaction flow; source documents; transaction capture; validation; journals/ledgers; master data; controls; audit trail; segregation of duties; information flow; reports; reconciliation; integrity between operational and accounting systems.

## Teaching notes

An accounting information system is an information-flow system, not merely accounting software. Start with the business event, identify the source evidence, validate it, record it, transform it, and report it.

Controls can be preventive, detective, or corrective. Audit trails make changes traceable. Segregation of duties reduces the risk that one person can initiate, approve, execute, and conceal a transaction.

## Exercises

1. Draw the transaction flow for a sale from order to accounting record.
2. Identify source documents for purchase, sale, payment, and receipt.
3. Design validation rules for customer, invoice, and payment records.
4. Classify controls as preventive/detective/corrective.
5. Identify segregation-of-duties conflicts in three cases.
6. Build a reconciliation between operational totals and accounting totals.
7. Design an audit trail for edits to an invoice.
8. Identify missing data in a transaction-processing system.
9. Explain why a database transaction can support accounting integrity.
10. Design a small schema linking business events to accounting entries.

## Coding workspace

[`lig_accounting_is_s2.py`](code/lig_accounting_is_s2.py) demonstrates one explicit validation contract: event ID, finite non-negative amount, source document, and approver. Treat these as a practice-case policy, not an official accounting rule. Reconciliation returns operational total minus accounting total.

## Problem-solving lab

Given a scenario where sales totals and accounting totals differ, reconstruct the information flow, enumerate possible failure points, and design a reconciliation procedure.

## Mastery gate

Model a fresh business process end-to-end and identify actors, data, controls, evidence, and reports without hand-waving.
