#!/usr/bin/env python3


def main():
    input = open('./input.txt', 'r').read()
    arrayLeft, arrayRight = getInput(input.split('\n'))

    distance = getDistance(arrayLeft, arrayRight)
    similarity = similarityScore(arrayLeft, arrayRight)

    print('Distance:', distance)
    print('Similarity:', similarity)


def similarityScore(l: list[int], r: list[int]) -> int:
    checkedLeft, score = {}, 0

    for val in l:
        if not val in checkedLeft:
            checkedLeft[val] = 1
        else:
            checkedLeft[val] += 1

    for val in r:
        if val in checkedLeft:
            score += val * checkedLeft[val]

    return score


def getDistance(l: list[int], r: list[int]) -> int:
    total = 0

    for index in range(len(l)):
        if l[index] > r[index]:
            total += l[index] - r[index]
        else:
            total += r[index] - l[index]

    return total


def getInput(input: list[str]) -> tuple[list[int], list[int]]:
    arrayLeft, arrayRight = [], []
    for val in input:
        if val != '':
            arrayLeft.append(int(val.split('-')[0]))
            arrayRight.append(int(val.split('-')[1]))

    return mergeSort(arrayLeft), mergeSort(arrayRight)


def mergeSort(input: list[int]) -> list[int]:
    if len(input) <= 1:
        return input

    mid = len(input) // 2
    left = mergeSort(input[:mid])
    right = mergeSort(input[mid:])

    return merge(left, right)


def merge(left: list[int], right: list[int]) -> list[int]:
    result = []
    leftIndex = rightIndex = 0

    while leftIndex < len(left) and rightIndex < len(right):
        if left[leftIndex] < right[rightIndex]:
            result.append(left[leftIndex])
            leftIndex += 1
        else:
            result.append(right[rightIndex])
            rightIndex += 1

    result.extend(left[leftIndex:])
    result.extend(right[rightIndex:])
    return result


if __name__ == "__main__":
    main()
