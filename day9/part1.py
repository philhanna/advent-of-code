#! /usr/bin/python

from point import Point
from knot import Knot
from grid import get_grid

FILENAME = "testdata/sample.dat"
DEBUG = True

# Initialize a grid with head, tail, and start knots with row and column
# equal to zero.

head = Knot("H", Point(0, 0))
tail = Knot("T", Point(0, 0))
start = Point(0, 0)

# Keep a list of the row and column of each position of the tail.
# Eliminate duplicates.

vlist = []
vlist.append(start)
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
            if tail.point not in vlist:                
                vlist.append(tail.point)

            # (optional) Draw the grid for visual inspection and debugging
            if DEBUG:
                print()
                print("\n".join(get_grid(head.point, tail.point, start)))
