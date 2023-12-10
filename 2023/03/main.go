package main

import (
    "fmt"
    "log"
    "bufio"
    "os"
    "unicode"
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

func checkMatch(index int, number map[string]string ) int {
    left, right := string(index - 1), string(index + 1)
    total := 0
    for key, value := range number {
        for _, match := range key {
            if string(index) == string(match) || left == string(match) || right == string(match) {
                num, err := strconv.Atoi(value)
                if err != nil {
                    log.Fatal(err)
                }
                total += num
                break
            }
        } 
    }
    return total
}

func isNumber(char rune) bool {
    return unicode.IsDigit(char)
}

func isChar(char rune) bool {
    return char != '.' && !unicode.IsDigit(char)
}

func gearRatios (data *bufio.Scanner) int {

    total := 0
    prevLineNums, prevLineChars := map[string]string {}, []int {}

    for data.Scan() {
        number, numberIndex := "", ""
        lineNums, lineChars := map[string]string {}, []int {}

        for index, char := range data.Text() {

            if isNumber(char) {
                number = number + string(char)
                numberIndex = numberIndex + string(index)
            } else if number != "" {
                lineNums[numberIndex] = number
                number, numberIndex = "", ""
            }

            if isChar(char) {
                lineChars = append(lineChars, index)
            }

            if index == len(data.Text()) - 1 && number != "" {
                lineNums[numberIndex] = number
            }

        }

        for _, value := range prevLineChars {
            total += checkMatch(value, lineNums)
        }
        for _, value := range lineChars {
            total += checkMatch(value, lineNums)
            total += checkMatch(value, prevLineNums)
        }

        prevLineNums = lineNums
        prevLineChars = lineChars
    }

    return total
}

func main() {
    data := getData()
    answer := gearRatios(data)
    fmt.Println("The answer for day 03 is:", answer)
}
