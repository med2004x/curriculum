// STUDENT PROBLEM SET — Lesson 071.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Implement a `POST` endpoint with a durable side-effect record and a unique idempotency key scoped to the caller.
// 2. Return the stored response for a repeated same-key/same-fingerprint request; reject same-key/different-fingerprint reuse.
// 3. Add a concurrent reservation path and document how in-progress duplicates wait, receive a conflict, or receive the eventual stored result. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 071: Idempotency, Transactions, and Retries
// Your task: make repeated delivery return the same stored result.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

type KeyStore map[string]string

func applyOnce(store KeyStore, key, result string) (string, bool) {
    // TODO: implement the assigned behavior.
    return "", false
}

func main() {}
