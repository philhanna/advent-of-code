// Define a data structure for an Elf.
class Elf {
    constructor(number, cals) {
        this.number = number;
        this.cals = cals;
    }
}

// Make an array to keep track of elf numbers and their total calories.
let elves = [];

// Read the input data into a list, with a blank line separating the entries
// between the calorie records for each elf.
const fs = require('fs');
let inputData = fs.readFileSync('input.dat', 'utf8').split('\n\n');

// Calculate the total calories for each elf.
for (let i = 0; i < inputData.length; i++) {
    // New elf.
    let elf = new Elf(i + 1, 0);
    // Calculate all this elf's calories.
    let entries = inputData[i].split('\n');
    for (let j = 0; j < entries.length; j++) {
        let cals = parseInt(entries[j]);
        elf.cals += cals;
    }
    // Add this elf to the list.
    elves.push(elf);
}

// Sort the elf list descending by total calories.
elves.sort((a, b) => b.cals - a.cals);

// Part 1.
let topElf = elves[0];
console.log(`Part1: Elf ${topElf.number} has ${topElf.cals} calories`);

// Part 2.
let total = elves.slice(0, 3).reduce((acc, elf) => acc + elf.cals, 0);
console.log(`Part2: Top three elves have ${total} calories`);