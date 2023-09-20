import ctypes
import random

# Linking the GO library
golib = ctypes.cdll.LoadLibrary("/Users/strCarne/VSCodeProjects/Python/labrinth/golib/golib.so")

# Test of connection with the library
if __name__ == "__main__":

    a_go = golib.A
    def a():
        a_go()

    hello_go = golib.Hello
    hello_go.argtypes = [ctypes.c_char_p]
    def hello(name):
        hello_go(name.encode("utf-8"))

    factorial_go = golib.Factorial
    factorial_go.argtypes = [ctypes.c_char_p]
    factorial_go.restype = ctypes.c_void_p
    def factorial(n):
        n = str(n)
        result = factorial_go(n.encode("utf-8"))
        return int(ctypes.string_at(result).decode("utf-8"))

    print(factorial(5))
    hello("Pusya")
    a()
    a()
    a()

#############
# Constants #
#############

# Exit positions
down_left = 0
up_left = 1
up_right = 2
down_right = 3

#############
# Functions #
#############
def generate_exit() -> int:
    return random.randrange(4)

def calc_exit(exit: int, size: int) -> list[int]:
    if exit == down_left:
        return [size - 2, 1]
    elif exit == up_left:
        return [1, 1]
    elif exit == up_right:
        return [1, size - 2]
    else:
        return [size - 2, size - 2]
    
def calc_position(i: int, j: int, size: int) -> int:
    return i * size + j

def cell_is_available(maze, size,  i, j) -> bool:
    return maze[i*size + j] == "\u0001"

make_maze_go = golib.MakeMaze
make_maze_go.argtypes = [ctypes.c_char_p, ctypes.c_char_p, ctypes.c_char_p, ctypes.c_char_p]
make_maze_go.restype = ctypes.c_void_p
def make_maze(size: int, i: int, j: int, exit: int) -> str:
    size = str(size)
    i = str(i)
    j = str(j)
    exit = str(exit)
    result = make_maze_go(size.encode("utf-8"), i.encode("utf-8"), j.encode("utf-8"), exit.encode("utf-8"))
    return ctypes.string_at(result).decode("utf-8")