class Tetromino:
    def __init__(self, shape):
        self.shape = shape  # single shape (no rotations)

# Define the 7 tetrominoes with their fixed shapes
I = Tetromino([[1, 1, 1, 1]])
J = Tetromino([[1, 0, 0], [1, 1, 1]])
L = Tetromino([[0, 0, 1], [1, 1, 1]])
O = Tetromino([[1, 1], [1, 1]])
S = Tetromino([[0, 1, 1], [1, 1, 0]])
T = Tetromino([[0, 1, 0], [1, 1, 1]])
Z = Tetromino([[1, 1, 0], [0, 1, 1]])

tetrominoes = [I, J, L, O, S, T, Z]
grid_size = 6
grid = [[0 for _ in range(grid_size)] for _ in range(grid_size)]

def can_place(tetromino, grid, row, col):
    shape = tetromino.shape
    for r in range(len(shape)):
        for c in range(len(shape[0])):
            if shape[r][c] == 1:
                if row + r >= len(grid) or col + c >= len(grid[0]) or grid[row + r][col + c] == 1:
                    return False
    return True

def place(tetromino, grid, row, col):
    shape = tetromino.shape
    for r in range(len(shape)):
        for c in range(len(shape[0])):
            if shape[r][c] == 1:
                grid[row + r][col + c] = 1

def remove(tetromino, grid, row, col):
    shape = tetromino.shape
    for r in range(len(shape)):
        for c in range(len(shape[0])):
            if shape[r][c] == 1:
                grid[row + r][col + c] = 0

def solve(tetrominoes, grid, index=0):
    if index == len(tetrominoes):
        return True
    for row in range(len(grid)):
        for col in range(len(grid[0])):
            if can_place(tetrominoes[index], grid, row, col):
                place(tetrominoes[index], grid, row, col)
                if solve(tetrominoes, grid, index + 1):
                    return True
                remove(tetrominoes[index], grid, row, col)
    return False

if solve(tetrominoes, grid):
    for row in grid:
        print(" ".join(str(cell) for cell in row))
else:
    print("No solution found")
