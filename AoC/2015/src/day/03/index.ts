/**
 * Solution to AoC 2015 day 03:
 *  - https://adventofcode.com/2015/day/3
 */
export function solution(input: string) {
    return `\n\tPart One: ${partOne(input)}\n\tPart Two: ${partTwo(input)}`;
}

export function partOne(input: string): number {
    return houseAlgorithm(input, { roboSanta: false });
}

export function partTwo(input: string): number {
    return houseAlgorithm(input, { roboSanta: true });
}

function houseAlgorithm(
    input: string,
    options: { roboSanta: boolean }
): number {
    let postions;

    if (options.roboSanta) {
        postions = [
            { x: 0, y: 0 }, // Santa
            { x: 0, y: 0 }  // roboSanta
        ];
    } else {
        postions = [{ x: 0, y: 0 }];
    }

    const visited = new Set<string>();
    visited.add(`0,0`);

    let index = 0;
    for (const char of input) {
        const who = index % 2;
        const p = postions[who]

        switch (char) {
            case '^': p.y++; break;
            case 'v': p.y--; break;
            case '<': p.x--; break;
            case '>': p.x++; break;
        }

        visited.add(`${p.x},${p.y}`);

        if (options.roboSanta) index++;
    }

    return visited.size;
};
