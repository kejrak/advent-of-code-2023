import numpy as np
import math

def bruteForce(time: list[str], dist: list[str]) -> None: 
    result = []
    for i, _ in enumerate(time):

        t = int(time[i])
        d = int(dist[i])

        milliseconds = 0
        counter = 0

        for _ in range(t+1):
            milimeter = (t - milliseconds) * milliseconds
            milliseconds += 1

            if milimeter > d:
                counter +=1

        result.append(counter)

    t = int(''.join(time))
    d = int(''.join(dist))

    num = 0
    for _ in range(t+1):
        milimeter = (t - milliseconds) * milliseconds
        milliseconds += 1

        if milimeter > d:
            num +=1

    print("First part:", np.prod(result))
    print("Second part:", num)

def useMath(time: list[str], dist: list[str]) -> None:


    t = int(''.join(time))
    d = int(''.join(dist))

    D = math.sqrt(t**2 - 4*d)
    x1 = (-t + D) / 2
    x2 = (-t - D) / 2

    diff = (int(math.floor(x1)) - int(math.ceil(x2))+1)
    
    print("Second part with clever math:", diff)

def main():
    with open("./input.txt") as f:
        data = f.read()
    time: list[str] = data.split("\n")[0].split(":")[1].split()
    dist: list[str] = data.split("\n")[1].split(":")[1].split()

    
    bruteForce(time, dist)
    useMath(time, dist)

if __name__ == "__main__":
    main()