class Point:
    def __init__(self, row, col):
        self.row = row
        self.col = col

    def touches(self, other):
        rowdif = self.row - other.row
        coldif = self.col - other.col
        return rowdif*rowdif + coldif*coldif <= 2
    
    def __eq__(self, other):
        return self.row == other.row and self.col == other.col
    
    def __str__(self):
        return repr(self)
    
    def __repr__(self):
        return f"({self.row},{self.col})"
    
    def __hash__(self):
        return hash(str(self))
