package main

import (
    "fmt"
    "log"
    "bufio"
    "os"
    "unicode"
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

func getDigitFromRow(rowOfDigits string) int{

    listNums := strings.Split(rowOfDigits, "")

    /*
     * Assuming that on the input every row must have one or more digits
     */

    if len(listNums) == 1 {
        rowOfDigits = rowOfDigits + rowOfDigits
    } else {
        rowOfDigits = listNums[0] + listNums[len(listNums) - 1]
    }

    rowNumber, err := strconv.Atoi(rowOfDigits)

    if err != nil {
        log.Fatal(err)
    }

    return rowNumber
}

func calulateCalibration(data *bufio.Scanner) int{

    total := 0

    for data.Scan() {
        placeHolder := ""
        for index, char := range data.Text() {

            check := unicode.IsDigit(char)

            if check {
                placeHolder = placeHolder + string(char)
            }

            length := len(strings.Split(data.Text(), ""))

            if index == length - 1 {
                total = total + getDigitFromRow(placeHolder)
                placeHolder = ""
                continue
            }
        }
    }
    return total
}

func main() {
    data := getData()
    answer := calulateCalibration(data)
    fmt.Println("The answer for day 01 is:", answer)
}
