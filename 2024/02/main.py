#!/usr/bin/env python3


def main():
    data = getInput()
    print('Safe reports (part one)', partOne(data))
    print('Safe reports (part two)', partTwo(data))


def partOne(data: list[list[int]]) -> int:
    count = 0

    for row in data:

        if not verifyRow(row):
            continue

        if verifyNums(row):
            count += 1

    return count


def partTwo(data: list[list[int]]) -> int:
    count = 0

    for row in data:
        isSafe = False

        for index in range(len(row)):
            tmp = row[:index] + row[index + 1:]
            if verifyRow(tmp) and verifyNums(tmp):
                isSafe = True

        if isSafe:
            count += 1

    return count


def verifyRow(row: list[int]) -> bool:
    rowSorted, rowSortedReversed = sorted(row), sorted(row, reverse=True)
    if row == rowSorted or row == rowSortedReversed:
        return True
    return False


def verifyNums(row: list[int]) -> bool:
    isSafe = True

    for index in range(len(row) - 1):
        if not 1 <= abs(row[index] - row[index + 1]) <= 3:
            isSafe = False
            break

    return isSafe


def getInput() -> list[list[int]]:
    input = open('./input.txt', 'r').read().split('\n')
    data = []

    for val in input:
        if val == '':
            continue
        data.append(list(map(int, val.split(' '))))

    return  data


if __name__ == "__main__":
    main()
