import pytest

from day9 import Point
from day9 import Knot
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

if __name__ == '__main__':
    pytest.main()
