package main

import (
    "fmt"
    "log"
    "bufio"
    "os"
    "strings"
)

func getData() *bufio.Scanner {

    file, err := os.Open("./input.txt")

    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)

    return(scanner)
}

func winningTotal (data *bufio.Scanner) int {

    total := 0

    for data.Scan() {
        values := strings.Split(data.Text(), ":")[1]

        winningNums := strings.Split(values, "|")[0]
        cardNums := strings.Split(data.Text(), "|")[1]

        value := 0
        for i, check := range strings.Split(cardNums, " ") {
            match := false
            for _, winningValue := range strings.Split(winningNums, " ") {
                if check == winningValue && check != "" {
                       match = true 
                }
            }

            if match == true && check != "" {
                if value >= 2 {
                    value = value * 2
                } else {
                    value = value + 1
                }
            }

            if i == len(strings.Split(cardNums, " ")) - 1 {
                total = total + value
            }
        }
    }

    return total
}

func main() {
    data := getData()
    answer := winningTotal(data)
    fmt.Println("The answer for day 04 is:", answer)
}
