class Point:
    def __init__(self, row, col):
        self.row = row
        self.col = col

    def touches(self, other):
        rowdif = self.row - other.row
        coldif = self.col - other.col

        # Use absolute values of the differences

        if rowdif < 0:
            rowdif = -rowdif
        if coldif < 0:
            coldif = -coldif

        return rowdif + coldif <= 2
    
    def __eq__(self, other):
        return self.row == other.row and self.col == other.col
