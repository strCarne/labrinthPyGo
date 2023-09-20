from tkinter import *
from random import randint
import lib
from time import sleep

quit = exit

# SETTINGS
player_color = "Purple"
exit_color = "Red"
blocked_zone_color = "Green"
available_zone_color = "White"
way_out_color = "Yellow"
cell_size = 20 # in pixels
size = 21
floors_played = 0
floors = 3

# Availability
is_available = '\u0001'
is_not_available = '\u0002'

#StartingPoint


exit = lib.generate_exit()
map, pos = lib.make_maze(size, exit)

scr = pos // size
scc = pos % size
ecr, ecc = lib.calc_exit(exit, size)

way_out = lib.find_way_out(scr * size + scc, ecr * size + ecc)
way_out_is_drown = False

def win_screen():
    ws = Tk()
    ws.title("CONGRATS!!!")
    wl = Label(ws, text="You won!", font=("Helvetica", 72))
    wl.pack()
    wl.mainloop()

def recalc_way_out(dir):
    global way_out
    if lib.calc_dir_char(dir) == way_out[0]:
        way_out = way_out[1:]
    else:
        way_out = lib.calc_dir_char(3 - dir) + way_out

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

draw(ecr, ecc, exit_color)
# print(revisited_cells)

def regenerate():
    global floors, floors_played
    global scr, scc
    global ecr, ecc
    global x1, y1
    global map
    global way_out, way_out_is_drown
    global exit

    floors_played += 1
    if floors_played == floors:
        root.destroy()
        win_screen()
        quit()

    if floors_played + 1 == floors:
        exit = lib.center
    else: 
        exit = lib.generate_exit()
    map, pos = lib.make_maze(size, exit)

    scr = pos // size
    scc = pos % size
    ecr, ecc = lib.calc_exit(exit, size)

    way_out = lib.find_way_out(scr * size + scc, ecr * size + ecc)
    way_out_is_drown = False

    create()

    y1 = scr * cell_size 
    x1 = scc * cell_size
    draw(scr, scc, player_color)

    draw(ecr, ecc, exit_color)

    lbl.config(text=f"CURRENT FLOOR: {floors - floors_played}")

def draw_full_way_out(color):
    r, c = scr, scc
    for dir in way_out[:len(way_out)-1]:
        if dir == 'r':
            c += 1
        elif dir == 'l':
            c -= 1
        elif dir == 'u':
            r -= 1
        elif dir == 'd':
            r += 1
        draw(r, c, color)

def draw_way_out():
    global way_out_is_drown
    if not way_out_is_drown:
        draw_full_way_out(way_out_color)
        way_out_is_drown = True
    else:
        draw_full_way_out(available_zone_color)
        way_out_is_drown = False

def draw_rect():
    ffs.create_rectangle((x1, y1, x1 + cell_size, y1 + cell_size), fill=player_color)

def del_rect():
    ffs.create_rectangle((x1, y1, x1 + cell_size, y1 + cell_size), fill=available_zone_color)

def move(event):
    global way_out
    global x1, y1
    global scr, scc
    global way_out_is_drown
    # print(event.char)
    del_rect()
    col = w = x1//cell_size
    row = h = y1//cell_size
    
    if event.char == 'o':
        draw_way_out()
    elif event.char == 'a':
        if map[row * size + col - 1] == is_available:
            l_old = len(way_out)
            recalc_way_out(lib.left)
            if l_old < len(way_out) and way_out_is_drown:
                ffs.create_rectangle((x1, y1, x1 + cell_size, y1 + cell_size), fill=way_out_color)
            x1 -= cell_size
            scc -= 1
    elif event.char == 'd':
        if map[row * size + col + 1] == is_available:
            l_old = len(way_out)
            recalc_way_out(lib.right)
            if l_old < len(way_out) and way_out_is_drown:
                ffs.create_rectangle((x1, y1, x1 + cell_size, y1 + cell_size), fill=way_out_color)
            x1 += cell_size
            scc += 1
    elif event.char == 'w':
        if map[(row - 1) * size + col] == is_available:
            l_old = len(way_out)
            recalc_way_out(lib.up)
            if l_old < len(way_out) and way_out_is_drown:
                ffs.create_rectangle((x1, y1, x1 + cell_size, y1 + cell_size), fill=way_out_color)
            y1 -= cell_size
            scr -= 1
    elif event.char == 's':
        if map[(row + 1) * size + col] == is_available:
            l_old = len(way_out)
            recalc_way_out(lib.down)
            if l_old < len(way_out) and way_out_is_drown:
                ffs.create_rectangle((x1, y1, x1 + cell_size, y1 + cell_size), fill=way_out_color)
            y1 += cell_size
            scr += 1
    if scr == ecr and scc == ecc:
        regenerate()
    draw_rect()
    col = w = x1//cell_size
    row = h = y1//cell_size


root.bind("<Key>", move)

def find_way_out():
    print(lib.find_way_out(scr * size + scc, ecr * size + ecc))

btn = Button(root, text="Show out/off a way out (press \"o\")", command=draw_way_out)
# lbl = Label(root, text=f"CURRENT FLOOR: {3 - floors_played}")

btn.pack()

lbl = Label(root, text="CURRENT FLOOR: 3")
lbl.pack()


root.mainloop()