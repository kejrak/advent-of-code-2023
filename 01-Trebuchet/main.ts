import fs from "fs"
import { off } from "process"

function sumOfVisibleNumbers(): number {
    const bff = fs.readFileSync('./input.txt')
    const data = String(bff)
        .split("\n")

    let sum = 0
    for (const line of data) {
        const chars = line.split("")
            .map(char => isNaN(parseInt(char)) ? 0 : parseInt(char))
            .filter(dig => dig !== 0)
            .map(dig => String(dig))

        const dig = parseInt(chars[0] + chars[chars.length - 1])
        if (!isNaN(dig)) {
            sum += dig
        }
    }
    return sum
}

const nums = [
    { label: "one", value: 1 },
    { label: "two", value: 2 },
    { label: "three", value: 3 },
    { label: "four", value: 4 },
    { label: "five", value: 5 },
    { label: "six", value: 6 },
    { label: "seven", value: 7 },
    { label: "eight", value: 8 },
    { label: "nine", value: 9 },
    { label: "1", value: 1 },
    { label: "2", value: 2 },
    { label: "3", value: 3 },
    { label: "4", value: 4 },
    { label: "5", value: 5 },
    { label: "6", value: 6 },
    { label: "7", value: 7 },
    { label: "8", value: 8 },
    { label: "9", value: 9 },
]

type IndexValue = {
    index: number,
    value: string
}

function numberOfOcc(line: string, sub: string): number {
    return line.split(sub).length - 1
}

function getIndexes(line: string, sub: string): number[] {
    const numOfOcc = numberOfOcc(line, sub)
    const result: number[] = []
    let prev = 0
    for (let i = 0; i < numOfOcc; i++) {
        const i = line.indexOf(sub, prev)
        prev = i + sub.length
        result.push(i)
    }
    return result
}

function sumOfAllNumbers(): number {
    const bff = fs.readFileSync("./input.txt")
    const data = String(bff).split("\n")

    let sum = 0
    let i = 1
    for (const line of data) {
        const arr: IndexValue[] = []
        for (const num of nums) {
            const indexes = getIndexes(line, num.label)
            for (const index of indexes) {
                arr.push({
                    index,
                    value: String(num.value)
                })
            }
        }

        arr.sort((a, b) => a.index - b.index)

        if (arr.length != 0) {
            const dig = parseInt(arr[0].value + arr[arr.length - 1].value)
            if (!isNaN(dig)) {
                sum += dig
            }
        }
    }
    return sum
}


const sumVisible = sumOfVisibleNumbers()
const sumNumbers = sumOfAllNumbers()

console.log(`Sum of all visible numbers ${sumVisible}`)
console.log(`Sum of all visible numbers ${sumNumbers}`)