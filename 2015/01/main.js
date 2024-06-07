import * as fs from 'fs';

/**
 * @param {string} input - Puzzle input
 * @returns {number}
 */
function main(input) {
    /** @type {number} */
    let result = 0;
    /** @type {boolean} */
    let basement = false;

    for (let i = 0; i < input.length; i++) {
        if (input.charAt(i) === '(') {
            result = result + 1
        } else if (input.charAt(i) === ')') {
            result = result - 1
        }

        if (basement === false && result === -1) {
            console.log("Entered the basement at: " + (i + 1))
            basement = true
        }
    }

    return(result)
}

fs.readFile('input.txt', (err, data) => {
    if (err) throw err;
    console.log(
        "Floor: " + main(data.toString())
    );
});

