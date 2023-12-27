package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getData() *bufio.Scanner {
    file, err := os.Open("./input.txt")

    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)
    return scanner
}

func formatHand(data *bufio.Scanner) map[int][]string {
    cardValues := map[int][]string {}
    i := 0

    for data.Scan() {
        cardValues[i] = append(cardValues[i], strings.Split(data.Text(), " ")[0])
        cardValues[i] = append(cardValues[i], strings.Split(data.Text(), " ")[1])
        i ++
    }

    return cardValues
}

func handType(hand string, types []string) string {
    cardMatches := map[rune] int {}
    returnType := ""

    for _, v := range hand {
        _, match := cardMatches[v]
        if match {
            cardMatches[v] ++
        } else {
            cardMatches[v] = 1
        }
    }

    if len(cardMatches) == 1 {
        returnType = types[0]
    } else if len(cardMatches) == 2 {
        currentType := types[1]
        for _, v := range cardMatches {
            if v == 2 {
                currentType = types[2]
            }
        }
        returnType = currentType
    } else if len(cardMatches) == 3 {
        currentType := types[4]
        for _, v := range cardMatches {
            if v == 3 {
                currentType = types[3]
            }
        }
        returnType = currentType
    } else if len(cardMatches) == 4 {
        returnType = types[5]
    } else {
        returnType = types[6]
    }

    return returnType
}

func stringToInt(data string) int {
    length, err := strconv.Atoi(data)

    if err != nil {
        log.Fatal(err)
    }

    return length
}

func winningHand(hand []string) []string {
    sortedHand := make([]string, len(hand))
    sengthOfChar := []string {"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}

    for i, v := range hand {
        sortedHand[i] = v
    }

    handOne := strings.Split(sortedHand[0], "")
    handTwo := strings.Split(sortedHand[1], "")

    for i := 0; i < len(handOne); i ++ {
        valOne := handOne[i]
        valTwo := handTwo[i]
        winner := false

        if valOne == valTwo {
            continue
        }

        for _, v := range sengthOfChar {
            if valOne == v && valTwo != v {
                winner = true
                break
            } else if valOne != v && valTwo == v {
                sortedHand[0], sortedHand[1] = sortedHand[1], sortedHand[0]
                winner = true
                break
            }
        }

        if winner {
            break
        }
    }

    return sortedHand
}

func sortHand(cardValues []string) []string{
    for i := 0; i < len(cardValues); i ++ {
        for j := i + 1; j < len(cardValues); j ++ {
            winner := winningHand([]string {cardValues[i], cardValues[j]})
            if winner[0] != cardValues[i] {
                cardValues[i], cardValues[j] = cardValues[j], cardValues[i]
            }
        }
    }

    return cardValues
}

func handRankings(cardValues map[int][]string, types []string) int{
    handTypesIndex := make(map[string][]int)
    rankings := len(cardValues)
    score := 0

    for i, v := range cardValues {
        handTypesIndex[v[2]] = append(handTypesIndex[v[2]], i)
    }

    for _, v := range types {
        handVal, match := handTypesIndex[v]

        if match {
            if len(handVal) == 1 {
                score = score + stringToInt(cardValues[handVal[0]][1]) * rankings
                rankings --
            } else {
                handsInPlay := []string {}
                handsIndex := map[string][]string {}
                for _, v := range handVal {
                    handsInPlay = append(handsInPlay, cardValues[v][0])
                    handsIndex[cardValues[v][0]] = append(handsIndex[cardValues[v][0]], cardValues[v][1])
                }
                winningOrder := sortHand(handsInPlay)
                for _, v := range winningOrder {
                    score = score + (stringToInt(handsIndex[v][0]) * rankings)
                    rankings --
                }
            }
        }
    }
    return score
}

func main() {
    types := []string {
        "fiveKind", "fourKind", "fullHouse", "threeKind", "twoPair", "onePair", "high",
    }
    data := getData()
    cardValues := formatHand(data)

    for i := 0; i < len(cardValues); i ++ {
        cardValues[i] = append(cardValues[i], handType(cardValues[i][0], types))
    }

    answer := handRankings(cardValues, types)

    fmt.Println("The answer for day 07 is:", answer)
}
