package main

import (
    "log"
    "bufio"
    "os"
    "fmt"
    "strings"
    "unicode/utf8"
)

type Instruction struct {
    move []string
    node map[string][]string
}

func getData() *bufio.Scanner {

    file, err := os.Open("./input.txt")

    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)

    return(scanner)
}

func trimFirstRune(s string) string {
    _, i := utf8.DecodeRuneInString(s)
    return s[i:]
}

func formatData(data *bufio.Scanner) *Instruction {
    firstLine := true
    formattedData := Instruction{node: make(map[string][]string)}

    for data.Scan() {
        line := data.Text()

        if firstLine {
            formattedData.move = strings.Split(line, "")
            firstLine = false
        } else {
            if line != "" {
                splitLine := strings.Split(line, " =")
                index := splitLine[0]
                value := strings.Split(splitLine[1], ",")
                valueLeft := strings.Split(value[0], "(")[1]
                valueRight := trimFirstRune(strings.Split(value[1], ")")[0])
                formattedData.node[index] = append(formattedData.node[index], valueLeft)
                formattedData.node[index] = append(formattedData.node[index], valueRight)
            }
        }
    }

    return &formattedData
}

func getAnswer(data *Instruction) int {
    movement := data.move
    currentNode := "AAA"
    endNode := "ZZZ"
    score := 0

    findMatch := false
    length := len(movement) - 1
    currentMovement := 0

    for !findMatch {
        if movement[currentMovement] == "R" {
            currentNode = data.node[currentNode][1]
        } else {
            currentNode = data.node[currentNode][0]
        }

        if currentMovement == length {
            currentMovement = 0
        } else {
            currentMovement++
        }

        if currentNode == endNode {
            findMatch = true
        }

        score++ 
    }


    return score
}

func main() {
    data := getData()
    formattedData := formatData(data)
    answer := getAnswer(formattedData)
    fmt.Println("The answer for day 08 is:", answer)
}
