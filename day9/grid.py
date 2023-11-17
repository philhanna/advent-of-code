from point import Point

def get_grid(head: Point, tail: Point, start: Point) -> []:
    """ Creates a grid with the head, tail, and starting points for
    visual inspection and debugging. Empty squares are shown as ".".
    The grid is returned as a list of strings.
    """

    # Find the lowest and highest row and column
    points = [head, tail, start]
    rows = [p.row for p in points]
    cols = [p.col for p in points]
    minRow = min(rows)
    minCol = min(cols)
    maxRow = max(rows)
    maxCol = max(cols)

    # Create a list of strings
    grid = []

    # Create each row of the list
    for row in range(minRow, maxRow+1):
        gridRow = ""
        for col in range(minCol, maxCol+1):
            point = Point(row, col)
            gridPoint = "."
            if point == start:
                gridPoint = "s"
            if point == tail:
                gridPoint = "T"
            if point == head:
                gridPoint = "H"
            gridRow += gridPoint
        grid.append(gridRow)

    # Return the list of grid lines
    return grid

def get_visited_grid(vlist: []) -> []:
    """ Creates a grid that shows all the places the tail has visited at
    least once. The grid is returned as a list of strings. """
    
    rows = [p.row for p in vlist]
    cols = [p.col for p in vlist]
    minRow = min(rows)
    minCol = min(cols)
    maxRow = max(rows)
    maxCol = max(cols)

    # Create a list of strings
    grid = []

    # Create each row of the list
    for row in range(minRow, maxRow+1):
        gridRow = ""
        for col in range(minCol, maxCol+1):
            point = Point(row, col)
            gridPoint = "."
            if point in vlist:
                gridPoint = "#"
            gridRow += gridPoint
        grid.append(gridRow)

    # Return the list of grid lines
    return grid