import fs from 'fs';
import path from 'path';

// Import all day modules
import * as day01 from './day/01/index';
import * as day02 from './day/02/index';
import * as day03 from './day/03/index';
import * as day04 from './day/04/index';
import * as day05 from './day/05/index';
import * as day06 from './day/06/index';
import * as day07 from './day/07/index';
import * as day08 from './day/08/index';
import * as day09 from './day/09/index';
import * as day10 from './day/10/index';
import * as day11 from './day/11/index';
import * as day12 from './day/12/index';
import * as day13 from './day/13/index';
import * as day14 from './day/14/index';
import * as day15 from './day/15/index';
import * as day16 from './day/16/index';
import * as day17 from './day/17/index';
import * as day18 from './day/18/index';
import * as day19 from './day/19/index';
import * as day20 from './day/20/index';
import * as day21 from './day/21/index';
import * as day22 from './day/22/index';
import * as day23 from './day/23/index';
import * as day24 from './day/24/index';
import * as day25 from './day/25/index';

const days = [
    day01, day02, day03, day04, day05, day06, day07, day08, day09, day10,
    day11, day12, day13, day14, day15, day16, day17, day18, day19, day20,
    day21, day22, day23, day24, day25,
];

function runDay(day: number) {
    const dayStr = day.toString().padStart(2, '0');
    const inputPath = path.resolve(__dirname, `day/${dayStr}/input.txt`);
    const input = fs.readFileSync(inputPath, 'utf-8');

    const dayModule = days[day - 1];
    console.log(`Day ${dayStr} answer:`, dayModule.solution(input));
}

// Parse command line arguments
const args = process.argv.slice(2);
const dayArg = args.find(arg => arg.startsWith('--day='));

if (!dayArg) {
    console.error('Error: --day parameter is required');
    console.error('Usage: node main.cjs --day=<1-25>');
    process.exit(1);
}

const specificDay = parseInt(dayArg.split('=')[1]);

if (isNaN(specificDay) || specificDay < 1 || specificDay > 25) {
    console.error('Error: Day must be a number between 1 and 25');
    process.exit(1);
}

runDay(specificDay);
