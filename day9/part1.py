#! /usr/bin/python

from point import Point
from knot import Knot
from grid import get_grid, get_visited_grid

FILENAME = "testdata/sample.dat"
DEBUG = False

# Initialize a grid with head, tail, and start knots with row and column
# equal to zero.

head = Knot("H", Point(0, 0))
tail = Knot("T", Point(0, 0))
start = Point(0, 0)

# Keep a list of the row and column of each position of the tail.

vlist = set()
vlist.add(start)
if DEBUG:
    print("\n".join(get_grid(head.point, tail.point, start)))

# Read the puzzle input one line at a time

with open(FILENAME) as fp:
    for line in fp:
        line = line.rstrip()        # remove newline

        # Parse the input line into a direction (UDLR) and a count

        direction, count = line.split()

        # Move the head according the the direction. After the move, if
        # the head and tail are no longer touching (see definition),
        # then move the tail as described in the definition.

        for i in range(int(count)):
            head.move(direction)
            tail.follow(head, direction)
            vlist.add(Point(tail.point.row, tail.point.col))

            # (optional) Draw the grid for visual inspection and debugging
            if DEBUG:
                print()
                print("\n".join(get_grid(head.point, tail.point, start)))
                print(f"{vlist}")

# (optional) Show the visited grid
if DEBUG:
    vgrid = get_visited_grid(vlist)
    print()
    print("\n".join(vgrid))

# Print the number of points the tail visited at least once
print(f"Total visited points = {len(vlist)}")