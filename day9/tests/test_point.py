import pytest

from point import Point


@pytest.mark.parametrize("row1,col1,row2,col2,expected", [
    (3, 4, 3, 4, True),     # Same point
    (2, 3, 2, 4, True),     # One to the right
    (2, 3, 2, 2, True),     # One to the left
    (2, 3, 1, 3, True),     # Up one row
    (2, 3, 3, 3, True),     # Down one row
    (1, 5, 2, 6, True),     # Diagonally adjacent
    (3, 5, 1, 4, False),    # Too far apart
])
def test_touches(row1, col1, row2, col2, expected):
    head = Point(row1, col1)
    tail = Point(row2, col2)
    actual = head.touches(tail)
    assert expected == actual


if __name__ == '__main__':
    pytest.main()
