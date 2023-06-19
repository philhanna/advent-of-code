# Define a data structure for an Elf.
Elf = Struct.new(:number, :cals)

# Make a list to keep track of elf numbers and their total calories.
elves = []

# Read the input data into a list, with a blank line separating the entries
# between the calorie records for each elf.
input_data = File.read("input.dat").split("\n\n")

# Calculate the total calories for each elf.
input_data.each_with_index do |entries, i|
  # New elf.
  elf = Elf.new(i + 1, 0)
  # Calculate all this elf's calories.
  entries.split("\n").each do |entry|
    cals = entry.to_i
    elf.cals += cals
  end
  # Add this elf to the list.
  elves << elf
end

# Sort the elf list descending by total calories.
elves.sort_by!(&:cals).reverse!

# Part 1.
top_elf = elves.first
puts "Part1: Elf #{top_elf.number} has #{top_elf.cals} calories"

# Part 2.
total = elves.first(3).map(&:cals).sum
puts "Part2: Top three elves have #{total} calories"
