import pytest

from point import Point
from grid import get_grid

@pytest.mark.parametrize("head,tail,start,expected", [
    (Point(1, 4), Point(2, 5), Point(5, 1), [
"...H.",
"....T",
".....",
".....",
"s...."
    ]),
])
def test_get_grid(head, tail, start, expected):
    actual = get_grid(head, tail, start)
    assert expected == actual
