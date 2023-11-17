import pytest

from point import Point
from knot import Knot


@pytest.mark.parametrize("row,col,direction,want_row,want_col", [
    (1, 5, "U", 0, 5),
    (1, 5, "D", 2, 5),
    (1, 5, "L", 1, 4),
    (1, 5, "R", 1, 6),
])
def test_move(row, col, direction, want_row, want_col):
    point = Point(row, col)
    head = Knot("H", point)
    head.move(direction)
    assert want_row == point.row
    assert want_col == point.col

@pytest.mark.parametrize("hrow,hcol,trow,tcol,direction,want_row,want_col", [
    (1, 3, 1, 3, "U", 1, 3),    # Same starting point
    (1, 3, 1, 3, "D", 1, 3),    # Same starting point
    (1, 3, 1, 3, "L", 1, 3),    # Same starting point
    (1, 3, 1, 3, "R", 1, 3),    # Same starting point
    (2, 6, 2, 5, "U", 2, 5),    # tail starts one to the left, head moves up
    (2, 6, 2, 7, "U", 2, 7),    # tail starts one to the right, head moves up
    (2, 6, 1, 6, "U", 1, 6),    # tail starts one up, head moves up
    (1, 6, 2, 5, "U", 1, 6),    # tail southeast one, head moves up
    (3, 3, 4, 2, "U", 3, 3),    # tail southeast, head moves up
    (3, 3, 4, 2, "R", 3, 3),    # tail southeast, head moves right
])
def test_follows(hrow, hcol, trow, tcol, direction, want_row, want_col):
    head = Knot("H", Point(hrow, hcol))
    tail = Knot("T", Point(trow, tcol))
    head.move(direction)
    tail.follow(head, direction)
    assert want_row == tail.point.row
    assert want_col == tail.point.col

if __name__ == '__main__':
    pytest.main()
