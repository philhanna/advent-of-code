# Define a data structure for an Elf.
class Elf:
    def __init__(self, number, cals):
        self.number = number
        self.cals = cals

# Make a list to keep track of elf numbers and their total calories.
elves = []

# Read the input data into a list, with a blank line separating the entries
# between the calorie records for each elf.
with open("input.dat") as f:
    inputData = f.read().split("\n\n")

# Calculate the total calories for each elf.
for i, entries in enumerate(inputData):
    # New elf.
    elf = Elf(i + 1, 0)
    # Calculate all this elf's calories.
    for entry in entries.split("\n"):
        cals = int(entry)
        elf.cals += cals
    # Add this elf to the list.
    elves.append(elf)

# Sort the elf list descending by total calories.
elves.sort(key=lambda x: x.cals, reverse=True)

# Part 1.
topElf = elves[0]
print(f"Part1: Elf {topElf.number} has {topElf.cals} calories")

# Part 2.
total = sum(elf.cals for elf in elves[:3])
print(f"Part2: Top three elves have {total} calories")
