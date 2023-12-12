from functools import cache


@cache
def count_possibilities(s: str, nums: tuple[int]) -> int:
    if s == "":
        return 1 if nums == () else 0

    if nums == ():
        return 0 if "#" in s else 1

    result = 0

    if s[0] in ".?":
        result += count_possibilities(s[1:], nums)

    if s[0] in "#?":
        if (
            nums[0] <= len(s)
            and "." not in s[: nums[0]]
            and (nums[0] == len(s) or s[nums[0]] != "#")
        ):
            result += count_possibilities(s[nums[0] + 1 :], nums[1:])
    return result


def main():
    with open("./input.txt") as f:
        data = f.read().strip()

    records = [
        (x, tuple(map(int, y.split(","))))
        for row in data.split("\n")
        for x, y in [row.split(" ")]
    ]

    partOne = 0
    partTwo = 0

    for springs, groups in records:
        partOne += count_possibilities(springs, groups)
        partTwo += count_possibilities("?".join([springs] * 5), (groups) * 5)

    print(partOne)
    print(partTwo)


if __name__ == "__main__":
    main()
