#!/usr/bin/env python3.11


import re as regEx


def main():
    input = open('./input.txt', 'r').read()
    digitsPattern = r"mul\(\d+,\d+\)"

    print('Part one:', partOne(regEx.findall(digitsPattern, input)))
    print('Part two:', partTwo(input, digitsPattern))


def partOne(data: list[str]) -> int:
    answer = 0

    for val in data:
        digits = val.replace('mul(', '').replace(')', '').split(',')
        answer += int(digits[0]) * int(digits[1])

    return answer


def partTwo(data: str, digitsPattern: str) -> int:
    enabled, validMul = True, []
    dontPattern = r"don't\(\)"
    doPattern = r"do\(\)"

    sections = regEx.split(f"({doPattern}|{dontPattern})", data)

    for section in sections:
        if regEx.fullmatch(doPattern, section):
            enabled = True
        elif regEx.fullmatch(dontPattern, section):
            enabled = False
        elif enabled:
            validMul.extend(regEx.findall(digitsPattern, section))

    return(partOne(validMul))


if '__main__' == __name__:
    main()
