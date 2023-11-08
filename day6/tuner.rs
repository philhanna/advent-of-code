use std::fs;

const LIMIT: usize = 4;

fn main() {
    let filename = "input.txt";
    let data = fs::read_to_string(filename).expect("Unable to read file");
    let (p, prev) = find_marker(&data);
    println!("p={}, prev={}", p, prev);
}

fn find_marker(s: &str) -> (usize, &str) {
    for p in LIMIT..s.len() {
        let prev = &s[p - LIMIT..p];
        if !dups(prev) {
            return (p, prev);
        }
    }
    (0, "")
}

fn dups(s: &str) -> bool {
    for (i, ci) in s.char_indices() {
        for (j, cj) in s.char_indices() {
            if i != j && ci == cj {
                return true;
            }
        }
    }
    false
}
