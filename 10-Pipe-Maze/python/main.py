from itertools import combinations
from collections import defaultdict
import numpy as np

def polygonArea(vertices):
    x, y = zip(*vertices)
    print(x, y)
    return 0.5 * abs(
        sum(x[i] * y[i - 1] - x[i - 1] * y[i] for i in range(len(vertices)))
    )

def pick_theorem(area: float, loop_size: int) -> int:
    return int(area - 0.5 *loop_size + 1)


def find_index_of_start(array: list[list[str]], startPoint: str) -> tuple[int, int]:
    for i, row in enumerate(array):
       for j, value in enumerate(row):
           if value == startPoint:
               return (i, j)
    return None

def get_symbol_at_index(array, index):
    row, col = index
    if 0 <= row < len(array) and 0 <= col < len(array[row]):
        return array[row][col]
    return None

def lookup_symbols_around(array, start_index):

    symbol = get_symbol_at_index(array ,start_index)

    if start_index is not None:
        row, col = start_index
        allowed_symbols_up = [] if symbol == 'F' or symbol == '7' or symbol == '-' else ['|', 'F', '7', 'S']
        allowed_symbols_down = [] if symbol == 'L' or symbol == 'J' or symbol == '-' else ['|', 'L', 'J', 'S']
        allowed_symbols_left = [] if symbol == '|' or symbol == 'L' or symbol == 'F' else ['-', 'L', 'F', 'S']
        allowed_symbols_right = [] if symbol == '|' or symbol == '7' or symbol == 'J' else ['-', 'J', '7', 'S']
        directions_map = {
            'up': (-1, 0, allowed_symbols_up),
            'down': (1, 0, allowed_symbols_down),
            'left': (0, -1, allowed_symbols_left),
            'right': (0, 1, allowed_symbols_right),
        }

        symbols_and_indexes = {direction: (None, None) for direction in directions_map}

        for direction, (row_change, col_change, allowed_symbols) in directions_map.items():
            new_row = row + row_change
            new_col = col + col_change
            if (
                0 <= new_row < len(array)
                and 0 <= new_col < len(array[new_row])
                and (allowed_symbols is None or array[new_row][new_col] in allowed_symbols)
            ):
                symbols_and_indexes[direction] = (array[new_row][new_col], (new_row, new_col))

        return symbols_and_indexes
    
    return None

def find_keys_by_value(dictionary, target_value):
    return [key for key, value in dictionary.items() if value == target_value]

def move_and_repeate(array, initial_start_index, visited_indices=None) -> tuple[int, int]:

    new_index = None
    
    if visited_indices is None:
        visited_indices = []

    startingPosition = initial_start_index
    visited_indices.append(startingPosition)

    symbols_around_start = lookup_symbols_around(array, initial_start_index)
    non_none_items = [(key, value) for key, value in symbols_around_start.items() if value[0] is not None]
    
    while new_index != startingPosition:
        symbols_around_start = lookup_symbols_around(array, initial_start_index)
        non_none_items = [(key, value) for key, value in symbols_around_start.items() if value[0] is not None]

        if not non_none_items:
            print("\nNo more non-None values. Stopping.")
            break

        for first_non_key, first_non_value in non_none_items:
            if first_non_key in symbols_around_start:
                _, new_index = symbols_around_start[first_non_key]

                if new_index not in visited_indices:
                    
                    if new_index != (initial_start_index and startingPosition):
                        visited_indices.append(new_index)
                        initial_start_index = new_index
                        break
                    else:
                        break 

    area = polygonArea(visited_indices)
    return len(visited_indices) // 2 + len(visited_indices) % 2, pick_theorem(area, len(set(visited_indices)))

def main():
    with open("./input.txt") as f:
        data = f.read()

    data: list[str] = str.split(data, "\n")
    initial_start_index = find_index_of_start(data, 'S')

    partOne, partTwo = (move_and_repeate(data, initial_start_index))
    
    print("Longest Distance", partOne)
    print("Fillups", partTwo)
    
if __name__ == "__main__":
    main()