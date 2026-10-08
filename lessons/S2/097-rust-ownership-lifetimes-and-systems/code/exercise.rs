fn transform(input: &str) -> String {
    // TODO: demonstrate borrowing/ownership/lifetime behavior from the lesson.
    input.to_owned()
}

fn main() {
    let input = String::from("systems");
    println!("{}", transform(&input));
}
