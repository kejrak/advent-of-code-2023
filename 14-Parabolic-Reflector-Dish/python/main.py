def transpose(grid: list[str]):
    return list(zip(*grid))


def rotate_grid(grid: list[str]):
    return list(zip(*grid[::-1]))


def move_stones(row: list[str]):
    counter = 0
    load = 0
    for i, symbol in enumerate(row):
        if symbol == ".":
            counter += 1
        elif symbol == "O":
            load += counter + len(row) - i
        else:
            counter = 0
    return load


def tilt(grid: tuple[str]) -> tuple[str]:
    for _ in range(4):
        grid = tuple(map("".join, zip(*grid)))
        grid = tuple(
            "#".join(
                [
                    "".join(sorted(tuple(group), reverse=True))
                    for group in row.split("#")
                ]
            )
            for row in grid
        )
        grid = tuple(row[::-1] for row in grid)
    return grid


def solve_stones(grid: list[str]) -> list[str]:
    seen: set[list[str]] = {grid}
    array: list[list[str]] = [grid]

    it = 0

    while True:
        it += 1
        grid = tilt(grid)
        if grid in seen:
            break
        seen.add(grid)
        array.append(grid)

    first = array.index(grid)
    print(it, first)

    grid = array[(1000000000 - first) % (it - first) + first]

    return grid


def main():
    with open("./input.txt") as f:
        grid = tuple(f.read().splitlines())

    # grid = [
    #     "#".join(
    #         ["".join(sorted(list(group), reverse=True)) for group in row.split("#")]
    #     )
    #     for row in grid
    # ]
    # grid = list(map("".join, zip(*grid)))

    grid = solve_stones(grid)
    print(grid)

    print(sum(row.count("O") * (len(grid) - r) for r, row in enumerate(grid)))


if __name__ == "__main__":
    main()
