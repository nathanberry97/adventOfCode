/**
 * Solution to AoC 2015 day 01:
 *  - https://adventofcode.com/2015/day/1
 */
export function solution(input: string) {
    return `\n\tPart One: ${partOne(input)}\n\tPart Two: ${partTwo(input)}`;
}

export function partOne(input: string): number {
    return floorAlgorithm(input, { stopAtBasement: false })
}

export function partTwo(input: string): number {
    return floorAlgorithm(input, { stopAtBasement: true })
}

function floorAlgorithm(
    input: string,
    options: { stopAtBasement: boolean }
): number {
    let floor = 0;

    for (let i = 0; i <= input.length; i++) {
        const char = input[i];

        if (char === '(') floor++;
        else if (char === ')') floor--;

        if (options.stopAtBasement && floor === -1) {
            return i + 1;
        };
    }

    return floor;
}
