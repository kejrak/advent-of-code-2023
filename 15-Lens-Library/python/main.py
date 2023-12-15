def hash_alogirthm(line):
    current = 0
    for value in line:
        current += ord(value)
        current *= 17
        current %= 256

    return current


def main():
    with open("./input.txt") as f:
        data = f.read().split(",")

    boxes = [[] for _ in range(256)]
    focal = {}

    total_sum = 0
    for line in data:
        if "-" in line:
            label = line[:-1]
            index = hash_alogirthm(label)
            if label in boxes[index]:
                boxes[index].remove(label)
        else:
            label, length = line.split("=")
            length = int(length)

            index = hash_alogirthm(label)

            if label not in boxes[index]:
                boxes[index].append(label)

            focal[label] = length

    total = 0

    for box_number, box in enumerate(boxes, 1):
        for lens_slot, label in enumerate(box, 1):
            total += box_number * lens_slot * focal[label]

    print(total)


if __name__ == "__main__":
    main()
