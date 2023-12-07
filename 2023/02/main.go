package main

import (
    "fmt"
    "log"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func getData() *bufio.Scanner {

    file, err := os.Open("./input.txt")

    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)

    return(scanner)
}

func possibleGame(redRoll int, blueRoll int, greenRoll int) bool {
    possible := true

    if redRoll > 12 || blueRoll > 14 || greenRoll > 13 {
        possible = false
    }

    return possible

}

func checkGame(data []string) bool {

    possible := true

    for _, hand := range data {
        roll := strings.Split(hand, ",")
        blue, red, green := 0, 0, 0

        for i, valueOfDice := range roll {
            content := strings.Split(valueOfDice, " ")

            value := convertStringToInt(content[1])
            colour := content[2]

            if colour == "blue" {
                blue = blue + value
            } else if colour == "red" {
                red = red + value
            } else {
                green = green + value
            }

            if i == len(roll) - 1 {
                if !possibleGame(red, blue, green) {
                    possible = false
                }
            }
        }
    }

    return possible
}

func convertStringToInt(data string) int {
    num, err := strconv.Atoi(data)

    if err != nil {
        log.Fatal(err)
    }

    return num
}

func cubeConundrum(data *bufio.Scanner) int {
    total := 0

    for data.Scan() {
        line := data.Text()
        gameContent := strings.Split(line, ":")
        gameNumber := convertStringToInt(strings.Split(gameContent[0], " ")[1])
        dice := strings.Split(gameContent[1], ";")

        if checkGame(dice) {
            total = total + gameNumber
        }
    }

    return total
}

func main() {
    data := getData()
    answer := cubeConundrum(data)
    fmt.Println("The answer for day 02 is:", answer)
}
