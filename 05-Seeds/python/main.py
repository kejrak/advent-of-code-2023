class Filter:
    name: str
    intervals: list[str]

    def __init__(self, name, intervals: list[str]):
        self.name = name
        self.intervals = intervals
    
    def get_mapping(self, value: int) -> int:
        result = value
        for interval in self.intervals:
            i = list(map(lambda x: int(x), interval.split()))
            d = value - i[1]
            if 0 <= d <= i[2]:
                return d + i[0]
        return result

def go_thrugh_filters(filters: list[Filter], value: int) -> int:
    filtered_value = value
    for filter in filters:
        filtered_value = filter.get_mapping(filtered_value)
    return filtered_value

def main():
    with open("./input.txt") as f:
        data = f.read()
    nums = data.split("\n")[0].split(":")[1].split()
    seeds = [x for x in nums if x != " "]

    filters: list[Filter] = []
    lines = data.split("\n")[2:]

    while len(lines) > 0:
        i = lines.index("") if "" in lines else len(lines)
        f = lines[0:i]
        filter = Filter(f[0], f[1:])
        filters.append(filter)
        lines = lines[i+1:]

    result = []
    for seed in seeds:
        res = go_thrugh_filters(filters, int(seed))
        result.append(res)

    print(min(result))

if __name__ == "__main__":
    main()