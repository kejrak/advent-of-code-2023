import fs from "fs"

const bff = fs.readFileSync('./input.txt');
const lines = String(bff).split("\n");

function removeEmpty(s: string): boolean {
    return s !== ''
}

function isWinning(s: string, arr: string[]): boolean {
    return arr.indexOf(s) !== -1 ? true : false
}

function main() {
    let result = 0
    for (const line of lines) {
        const [_, numbers] = line.split(":")
        const [left, right] = numbers.split("|")

        const leftArr = left.split(" ").filter(removeEmpty)
        const rightArr = right.split(" ").filter(removeEmpty)


        const res = recursionSumRec(leftArr, rightArr)

        result += res

    }
    console.log(result)
}

function recursionSumRec(numbers: string[], reference: string[]): number {
    let total = 0


    function recursionSum(numbers: string[], reference: string[], sum: number) {
        if (numbers.length === 0) {
            total = sum
            return
        }


        const winning = isWinning(numbers[0], reference)

        if (winning) {
            if (sum === 0) {
                sum = 1
            } else {
                sum *= 2

            }
            recursionSum(numbers.slice(1, numbers.length), reference, sum)
            return
        }

        recursionSum(numbers.slice(1, numbers.length), reference, sum)
    }

    recursionSum(numbers, reference, 0)

    return total
}

// function recursionSumRec(numbers: string[], reference: string[]): number {
//     let total = 0


//     function recursionSum(numbers: string[], reference: string[], isStreak: boolean, sum: number) {
//         if (numbers.length === 0) {
//             total += sum
//             console.log(total)
//             return
//         }


//         const winning = isWinning(numbers[0], reference)

//         if (!isStreak && winning) {
//             sum = 1
//             recursionSum(numbers.slice(1, numbers.length), reference, true, sum)
//             return
//         }


//         if (isStreak && winning) {

//             sum *= 2
//             recursionSum(numbers.slice(1, numbers.length), reference, true, sum)
//             return
//         }

//         total += sum

//         recursionSum(numbers.slice(1, numbers.length), reference, false, 0)
//     }

//     recursionSum(numbers, reference, false, 0)

//     return total
// }


main()