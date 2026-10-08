// STUDENT PROBLEM SET — Lesson 097.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: Take a small C program from C track Lesson 4 (e.g. your dynamic vector implementation) and translate it into Rust using a `Vec<T>` (Rust's standard growable array type, which internally manages exactly the allocation/reallocation logic your C version implemented manually).
// 2. Extend it with one explicit boundary case: Confirm the Rust version, using the standard library type, requires no manual `free`/`malloc` equivalent at all — ownership and the `Drop` trait handle cleanup automatically.
// 3. Rebuild the same behavior with one changed constraint: Deliberately reproduce C track Lesson 4's use-after-free bug's *intent* in Rust (attempt to use a value after it's been moved elsewhere) and confirm the Rust compiler refuses to compile it, with an error message explicitly citing the move — read the actual compiler error text and confirm it correctly identifies both the move point and the invalid subsequent. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 097: Rust Ownership, Lifetimes, and Systems Thinking
// Your task: demonstrate borrowing/ownership/lifetime behavior from the lesson.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
fn transform(input: &str) -> String {
    // TODO: implement the behavior you specified in this lesson's WORK.md.
    todo!("implement the assigned behavior");
}

fn main() {
    let input = String::from("systems");
    println!("{}", transform(&input));
}
