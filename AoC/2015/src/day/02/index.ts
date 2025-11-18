/**
 * Solution to AoC 2015 day 02:
 *  - https://adventofcode.com/2015/day/2
 */
export function solution(input: string) {
    return `\n\tPart One: ${partOne(input)}\n\tPart Two: ${partTwo(input)}`;
}

export function partOne(input: string): number {
    return wrappingPaperAlogrithm(input, { feetOfRibbon: false });
}

export function partTwo(input: string): number {
    return wrappingPaperAlogrithm(input, { feetOfRibbon: true });
}

function wrappingPaperAlogrithm(
    input: string,
    options: { feetOfRibbon: boolean }
): number {
    let result = 0
    const dimensions = input.split('\n')

    for (let i = 0; i < dimensions.length; i++) {
        const dimension = dimensions[i].split('x');

        if (dimension.length !== 3) continue;

        const length = parseInt(dimension[0], 10);
        const width = parseInt(dimension[1], 10);
        const height = parseInt(dimension[2], 10);

        if (options.feetOfRibbon) {
            const sides = [length, width, height]
            const volume = length * width * height;

            const [a, b] = sides.sort((x, y) => x - y).slice(0, 2);
            result += a * 2 + b * 2 + volume;
        } else {
            const side1 = length * width;
            const side2 = width * height;
            const side3 = height * length;

            result +=
                2 * (side1 + side2 + side3)
                + Math.min(side1, side2, side3);
        }
    }

    return result
}
