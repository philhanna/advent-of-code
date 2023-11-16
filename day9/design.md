## Definitions

### Knot
- A knot has a row and column, which are updated when it moves.
- There are only two knots in the puzzle; a **head** knot and a **tail** knot.

### Touching
The head and tail knots are considered to be *touching* if any of these
four conditions are true:
- They are in the same row and their columns differ by at most 1
- Their rows differ by at most 1 and they are in the same column
- They are in adjacent rows and their columns differ by at most 1
- Their rows differ by at most 1 and they are in adjacent rows

This can be simplified as saying that the following equation holds
true:

(row<sub><i>H</i></sub> - row<sub><i>T</i></sub>)<sup>2</sup> + 
(col<sub><i>H</i></sub> - col<sub><i>T</i></sub>)<sup>2</sup> &le; 2

or, equivalently:

|row<sub><i>H</i></sub> - row<sub><i>T</i></sub>| +
|col<sub><i>H</i></sub> - col<sub><i>T</i></sub>| &le; 1


### Move

Knots move as follows:
- To move **up**, subtract 1 from the row
- To move **down**, add 1 to the row
- To move **left**, subtract 1 from the column
- To move **right**, add 1 to the column

The head and tail must always be touching after each move is completed.

The head is moved according to commands (Up/Down/Left/Right)
If the head and tail are still touching, the tail does not move.
Otherwise, the tail's position is adjusted according to what just
happened to the head:
- To move **up**, tail moves up and into the head's column
- To move **down**, tail moves down and into the head's column
- To move **left**, tail moves left and into the head's row
- To move **right**, tail moves right and into the head's row

## Functions

### Draw the grid
1. Represent
   - the head knot as ***H***
   - the tail as ***T***,
   - the starting position as ***s***
   - an empty square as ***.***
2. To start with, all ***H***, ***T***, and ***s*** have row and column equal to zero
3. Always draw the three in their current row and column,
   starting with ***s***, then ***T***, and finally ***H***.   
   This will ensure that in case of any overlap, ***T***
   will overlay ***s***, and ***H*** will overlay either
   of them.
4. Add empty squares for the following rows and columns:
   - Rows one less than the lowest row of the three positions
     to one greater than the highest row
   - Columns one less than the lowest column to one greater than the highest
5. Normalize all drawn rows and column so that the lowest is zero