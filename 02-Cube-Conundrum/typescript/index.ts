import fs from "fs"
import { off } from "process";

function getGames() {
    const bff = fs.readFileSync('./input.txt');
    const lines = String(bff).split("\n");

    let sum = 0
    let count = 1
    const totals: { [color: string]: number[] } = {};
    for (const line of lines) {
        const values = line
            .split(': ')[1]
            .split("; ")


        const results: boolean[] = [];
        for (const round of values) {

            const game = round
                .split(' ')

            let redCount = 0
            let greenCount = 0
            let blueCount = 0

            for (let i = 0; i < game.length; i += 2) {
                const quantity = parseInt(game[i])
                const color = game[i + 1].replace(',', '')

                if (color === 'red' && quantity !== 0) {
                    redCount = quantity
                } else if (color === 'green' && quantity !== 0) {
                    greenCount = quantity
                } else if (color === 'blue' && quantity !== 0) {
                    blueCount = quantity
                }

                (redCount <= 12) ? results.push(true) : results.push(false);
                (greenCount <= 13) ? results.push(true) : results.push(false);
                (blueCount <= 14) ? results.push(true) : results.push(false);


            }
            console.log(`Round - Red: ${redCount}, Green: ${greenCount}, Blue: ${blueCount}`)
        }
        console.log(results)
        const allTrue = results.every(value => value === true);

        if (allTrue) {
            sum += count
        }
        count++
    }
    console.log(sum)
}

getGames()

