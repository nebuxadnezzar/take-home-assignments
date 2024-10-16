
moves = [(-1, 0), (0, 1), (1, 0), (0, -1)]
wall = '#'
mark = ';'
key = '*'
out_callback = None

def make_maze(dim : tuple) -> list[list]:
    maze = list([])
    s, x, *_ = dim
    for i in range(0, len(s), x):
        z = s[i:i+x]
        maze.append([*z])
    return maze

def make_out_callback(m : list[list]):
    def f(path : list[tuple]):
        print_solution(m, path)
    return f

def print_maze(m : list[list]):
    print("MAZE:")
    for i in range(len(m)):
        mm = m[i]
        print(f"{i:02} {''.join(mm)}")
    print()

def print_solution(m : list[list], path : list[tuple]):
    print_maze(m)
    print("SOLUTION:")
    for d in path:
        m[d[0]][d[1]] = mark
    for i in range(len(m)):
        mm = m[i]
        print(f"{i:02} {''.join(mm)}")
    print()

def count_keys(m : list[list], vis : set )-> int:
    cnt = 0
    for d in vis:
        if m[d[0]][d[1]] == key:
            cnt += 1
    return cnt

def go(pos : tuple, move : tuple) -> tuple:
    return pos[0] + move[0], pos[1] + move[1]

def is_outside_maze(pos : tuple, maze : list[list]) -> bool:
    try:
        _ =maze[pos[0]][pos[1]]
        return False
    except:
        return True

def is_valid(pos : tuple, maze : list[list]) -> bool:
    return not is_outside_maze(pos, maze) and maze[pos[0]][pos[1]] != wall

def find_path(cur : tuple, dest : tuple, path : list, maze : list[list], vis : set):

    if cur == dest:
        #print(f"PATH: {path}")
        print(f"KEY COUNT: {count_keys(maze, vis)}")
        out_callback(path)
        return

    for move in moves:
        next = go(cur, move)
        if is_outside_maze(next, maze):
            path.append(cur)
            vis.add(cur)
            print(f"KEY COUNT: {count_keys(maze, vis)}")
            out_callback(path)
        if next in vis:
            continue
        #print(f"isvalid: {next} {is_valid(next, maze)}")
        if is_valid(next, maze):
            path.append(next)
            vis.add(next)
            #print(f"-> {path}")
            find_path(next, dest, path, maze, vis)
    if len(path) > 0:
       path.pop()


mazes = \
[
(
'#.#############'+\
'#......##...###'+\
'###.#####.#.###'+\
'###...##.....*#'+\
'#....###.##.###'+\
'#*##.....##.###'+\
'###########.###'+\
'########..*.###'+\
'#........######'+\
'####.##.#######',
15, 10,
(0, 0), (9, 4)
)
]


for dim in mazes:
    m = make_maze(dim)
    out_callback = make_out_callback(m)
    find_path(dim[3], dim[4], [], m, set([]))