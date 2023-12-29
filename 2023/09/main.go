package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var answer int;

func getData() *bufio.Scanner {
    file, err := os.Open("input.txt")

    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)

    return scanner
}

func stringToInt(data string) int {
    value, err := strconv.Atoi(data)

    if err != nil {
        log.Fatal(err)
    }

    return value
}   

func formatData(data *bufio.Scanner) map[int][]int {
    var formattedData = make(map[int][]int)
    var index int

    for data.Scan() {
        var line = strings.Split(data.Text(), " ")
        for i := 0; i < len(line); i++ {
            formattedData[index] = append(formattedData[index], stringToInt(line[i]))
        }
        index++
    }

    return formattedData
}

func getAnswer(data map[int][]int) {
    var previousValue int
    newIndex, index, intialLoop := len(data), len(data) - 1, true

    for _, value := range data[index]{
        if !intialLoop {
            data[newIndex] = append(data[newIndex], value - previousValue)
        }
        previousValue = value
        intialLoop = false
    }
    
    firstValue := data[newIndex][0]
    lastValue := data[newIndex][len(data[newIndex]) - 1]

    if firstValue == lastValue && lastValue == 0 {
        nextNum := 0
        for i := len(data) - 1; i >= 0; i-- {
            nextNum += data[i][len(data[i]) - 1]
        }
        answer += nextNum
    } else {
        getAnswer(data)
    }
}

func main() {
    data := getData()
    formattedData := formatData(data)

    for _, value := range formattedData {
        getAnswer(map[int][]int{0: value})
    }

    fmt.Println("The answer for day 09 is:", answer)
}
