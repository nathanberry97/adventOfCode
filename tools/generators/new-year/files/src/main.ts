import fs from 'fs';
import path from 'path';

// Import all day modules
import * as day01 from './day/01/01';
import * as day02 from './day/02/02';
import * as day03 from './day/03/03';
import * as day04 from './day/04/04';
import * as day05 from './day/05/05';
import * as day06 from './day/06/06';
import * as day07 from './day/07/07';
import * as day08 from './day/08/08';
import * as day09 from './day/09/09';
import * as day10 from './day/10/10';
import * as day11 from './day/11/11';
import * as day12 from './day/12/12';
import * as day13 from './day/13/13';
import * as day14 from './day/14/14';
import * as day15 from './day/15/15';
import * as day16 from './day/16/16';
import * as day17 from './day/17/17';
import * as day18 from './day/18/18';
import * as day19 from './day/19/19';
import * as day20 from './day/20/20';
import * as day21 from './day/21/21';
import * as day22 from './day/22/22';
import * as day23 from './day/23/23';
import * as day24 from './day/24/24';
import * as day25 from './day/25/25';

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

for (let d = 1; d <= 25; d++) {
    runDay(d);
}
