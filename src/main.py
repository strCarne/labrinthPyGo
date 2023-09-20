from tkinter import *
from random import randint
import lib

# SETTINGS
player_color = "Purple"
exit_color = "Red"
blocked_zone_color = "Green"
available_zone_color = "White"
cell_size = 15 # in pixels
size = 50

# Availability
is_available = '\u0001'
is_not_available = '\u0002'

#StartingPoint

# starting color of row
scr = randint(1, size)

# starting random column
scc = randint(1, size)

exit = lib.generate_exit()
map = lib.make_maze(size, scr, scc, exit)

def create():
    "Create a rectangle with draw function (below) with random color"
    for row in range(size):
        for col in range(size):
            if map[row * size + col] == is_available:
                color = available_zone_color
            elif map[row * size + col] == is_not_available:
                color = blocked_zone_color
            draw(row, col, color)


def draw(row, col, color):
    x1 = col * cell_size
    y1 = row * cell_size
    x2 = x1 + cell_size
    y2 = y1 + cell_size
    ffs.create_rectangle(x1, y1, x2, y2, fill=color)

root = Tk()
root.title('Maze')
canvas_side = size*cell_size
ffs = Canvas(root, width = canvas_side, height = canvas_side, bg = 'grey')
ffs.pack()


create()

y1 = scr * cell_size 
x1 = scc * cell_size
draw(scr, scc, player_color)

ecr, ecc = lib.calc_exit(exit, size)
draw(ecr, ecc, exit_color)
# print(revisited_cells)

def draw_rect():
    ffs.create_rectangle((x1, y1, x1 + cell_size, y1 + cell_size), fill=player_color)

def del_rect():
    ffs.create_rectangle((x1, y1, x1 + cell_size, y1 + cell_size), fill=available_zone_color)

def move(event):
    global x1, y1
    # print(event.char)
    del_rect()
    col = w = x1//cell_size
    row = h = y1//cell_size
    print("before", map[row * size + col])
    if event.char == 'a':
        if map[row * size + col - 1] == is_available:
            x1 -= cell_size
    elif event.char == 'd':
        if map[row * size + col + 1] == is_available:
            x1 += cell_size
    elif event.char == 'w':
        if map[(row - 1) * size + col] == is_available:
            y1 -= cell_size
    elif event.char == 's':
        if map[(row + 1) * size + col] == is_available:
            y1 += cell_size

    draw_rect()
    col = w = x1//cell_size
    row = h = y1//cell_size
    print(w, h)
    print("after", map[row * size + col])


root.bind("<Key>", move)


root.mainloop()