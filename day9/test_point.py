import pytest

from point import Point

@pytest.mark.parametrize("row1,col1,row2,col2,expected", [
    (1, 5, 2, 6, True),
])
def test_touches(row1, col1, row2, col2, expected):
    head = Point(row1, col1)
    tail = Point(row2, col2)
    actual = head.touches(tail)
    assert expected == actual

if __name__ == '__main__':
    pytest.main()

