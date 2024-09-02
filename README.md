
## Game Of Life

Definition
The universe of the Game of Life is an infinite two-dimensional orthogonal grid of square cells, each of
which is in one of two possible states, alive or dead. Every cell interacts with its eight neighbors, which
are the cells that are horizontally, vertically, or diagonally adjacent.
Rules
At each step in time, the following transitions occur:
1. Any live cell with fewer than two live neighbors dies as if caused by underpopulation.
2. Any live cell with two or three live neighbors lives on to the next generation.
3. Any live cell with more than three live neighbors dies, as if by overcrowding.
4. Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.

## Prerequisites:

1. Golang installed

## Step to run:

1. Clone this repo
2. Run command: go run main.go
