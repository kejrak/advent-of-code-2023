def transpose(pattern: list[str]):
    return list(zip(*pattern))


def get_smudged(a: list[str], b: list[str]) -> bool:
    smudge = 0
    for i in range(len(a)):
        for j in range(len(a[i])):
            if a[i][j] != b[i][j]:
                smudge += 1

    return smudge == 1


def incidence(pattern: list[str]) -> int:
    for i in range(1, len(pattern)):
        above = pattern[:i][::-1]
        below = pattern[i:]

        above = above[: len(below)]
        below = below[: len(above)]

        # - For part two, you shoudl use - #
        # if get_smudge(above, below):
        #     return i

        if above == below:
            return i
    return 0


def solve(pattern: list[str]) -> int:
    sum = 0

    row = incidence(pattern)
    sum += row * 100
    col = incidence(transpose(pattern))
    sum += col

    return sum


def main():
    with open("./input.txt") as f:
        data = f.read().strip()

    patterns: list[list[str]] = [pattern.split("\n") for pattern in data.split("\n\n")]

    solved: int = 0

    for i, pattern in enumerate(patterns):
        solved += solve(pattern)

    print(solved)


if __name__ == "__main__":
    main()
