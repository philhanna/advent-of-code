use std::fs;

// Define a data structure for an Elf.
struct Elf {
    number: i32,
    cals: i32,
}

impl Elf {
    fn new(number: i32) -> Elf {
        Elf { number, cals: 0 }
    }
}

fn main() {
    // Make a list to keep track of elf numbers and their total calories.
    let mut elves: Vec<Elf> = Vec::new();

    // Read the input data into a list, with a blank line separating the entries
    // between the calorie records for each elf.
    let input_data = fs::read_to_string("input.dat").unwrap();
    let input_data: Vec<&str> = input_data.split("\n\n").collect();

    // Calculate the total calories for each elf.
    for (i, entries) in input_data.iter().enumerate() {
        // New elf.
        let mut elf = Elf::new(i as i32 + 1);
        // Calculate all this elf's calories.
        for entry in entries.split("\n") {
            let cals = entry.parse::<i32>().unwrap();
            elf.cals += cals;
        }
        // Add this elf to the list.
        elves.push(elf);
    }

    // Sort the elf list descending by total calories.
    elves.sort_by(|a, b| b.cals.cmp(&a.cals));

    // Part 1.
    let top_elf = &elves[0];
    println!("Part1: Elf {} has {} calories", top_elf.number, top_elf.cals);

    // Part 2.
    let total: i32 = elves.iter().take(3).map(|elf| elf.cals).sum();
    println!("Part2: Top three elves have {} calories", total);
}