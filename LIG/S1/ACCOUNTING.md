# FSEGS LIG — Comptabilité générale

**Semester:** S1  
**Attached engineering lesson:** 012 — Integrated Lab 1: Small C CLI

## Scope

Accounting equation; assets; liabilities; equity; revenue; expenses; accounts; debit/credit; journal; ledger; trial balance; inventory basics; depreciation basics; adjustments; income statement; balance sheet; closing logic.

## Core rule

Every transaction must preserve the accounting equation:

`Assets = Liabilities + Equity`

The debit/credit side is determined by the account class and the accounting rule being used. Never decide debit/credit from whether cash “feels like money coming in”. Identify the accounts first, then apply the rule.

## Worked transaction set

1. Owner invests cash.
2. Business buys equipment for cash.
3. Business buys inventory on credit.
4. Business makes a cash sale above cost.
5. Business pays a supplier.
6. Business pays an operating expense.

For each transaction: identify accounts, debit, credit, reason, and effect on the accounting equation. Then post to a ledger and prepare a trial balance.

## Exercises

1. Classify 20 items as asset, liability, equity, revenue, or expense.
2. Record 12 transactions in a journal.
3. Post the transactions to ledger accounts.
4. Prepare a trial balance and find one deliberately introduced imbalance.
5. Record a sale with separate revenue and cost-of-goods information.
6. Record a purchase on credit and subsequent settlement.
7. Apply a simple depreciation rule to an equipment item.
8. Prepare a minimal income statement.
9. Prepare a minimal balance sheet.
10. Explain how each transaction preserves the accounting equation.

## Digital workspace

Use [`lig_accounting_s1.py`](code/lig_accounting_s1.py) to total debit and credit columns and flag imbalances. Create your own CSV transaction set and write every journal entry before checking the debit/credit totals.

## Problem-solving lab

Given a month's mixed transactions, reconstruct the journal from source descriptions, post to a ledger, prepare the trial balance, then derive the basic financial statements. Deliberately corrupt one transaction and locate the error without looking at a solution.

## Exam patterns

Account classification; journal entry; ledger posting; balance extraction; trial balance; statement preparation; explain the impact of a transaction; identify and repair an accounting error.

## Mastery gate

Complete a fresh 15-transaction set from blank and produce a balanced trial balance plus basic statements. Explain every non-obvious entry.
