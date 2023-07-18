import re

def expand(pair):
    n1, n2 = map(int, pair.split('-'))
    return n1, n2

def contains(range1, range2):
    n1a, n1b = expand(range1)
    n2a, n2b = expand(range2)
    return (n1a <= n2a and n1b >= n2b) or (n2a <= n1a and n2b >= n1b)

with open("input", "r") as file:
    data = file.read()

count = 0

for line in data.split("\n"):
    match = re.search(r"(\d+-\d+),(\d+-\d+)", line)
    if match:
        range1, range2 = match.groups()
        if contains(range1, range2):
            count += 1

print(f"Count={count}")
