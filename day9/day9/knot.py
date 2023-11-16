class Knot:
    def __init__(self, name, point):
        self.name = name
        self.point = point

    def move(self, direction):
        """ Moves the knot in the specified direction """
        if direction == "U":
            self.point.row -= 1
        elif direction == "D":
            self.point.row += 1
        elif direction == "L":
            self.point.col -= 1
        elif direction == "R":
            self.point.col += 1
        else:
            raise RuntimeError(f"Undefined direction \"{direction}\"")