# Define the Tetris blocks
tetrominoes = [
    [[1, 1], [1, 1]], # O
    [[1, 0], [1, 1], [1, 0]], # T
    [[0, 1, 1], [1, 1, 0]], # Z
    [[0, 0, 1], [0, 1, 1], [1, 1, 0]], # S
    [[1, 1, 1], [0, 0, 1]], # L
    [[1, 1], [0, 1], [0, 1]], # Mirrored L
]

# Initialize the 7x7 grid
grid = [[0 for _ in range(7)] for _ in range(7)]

def canPlace(tetromino, grid, row, col):
    for i in range(len(tetromino)):
        for j in range(len(tetromino[i])):
            if tetromino[i][j] == 1:
                if row + i >= len(grid) or col + j >= len(grid[0]) or grid[row + i][col + j] == 1:
                    return False
    return True

def place(tetromino, grid, row, col):
    for i in range(len(tetromino)):
        for j in range(len(tetromino[i])):
            if tetromino[i][j] == 1:
                grid[row + i][col + j] = 1
            

def remove(tetromino, grid, row, col):
    for i in range(len(tetromino)):
        for j in range(len(tetromino[i])):
            if tetromino[i][j] == 1:
                grid[row + i][col + j] = 0
            print(tetromino, grid, row, col)

def solve(tetrominoes, grid, index):
    if index == len(tetrominoes):
        return True
    for row in range(len(grid)):
        for col in range(len(grid[0])):
            if canPlace(tetrominoes[index], grid, row, col):
                place(tetrominoes[index], grid, row, col)
                if solve(tetrominoes, grid, index + 1):
                    return True
                remove(tetrominoes[index], grid, row, col)
    return False

# Flatten the tetrominoes list
flattened_tetrominoes = tetrominoes * 2  # Repeat the list to get 12 blocks

# Solve the puzzle
if solve(flattened_tetrominoes, grid, 0):
    for row in grid:
        print(' '.join(str(x) for x in row))
else:
    print("No solution found.")
