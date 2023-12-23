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

func formatData(data *bufio.Scanner) ([]string, []string) {
    time, distance := []string {}, []string {}
    count := 0

    for data.Scan() {
        test := strings.Fields(data.Text())
        for _, v := range test {
            if count == 0 && v != "Time:" {
               time = append(time, v) 
            } else if v != "Time:" && v != "Distance:"{
               distance = append(distance, v) 
            }
        }
        count ++
    }

    return time, distance
}

func stringToInt(data string) int {
    length, err := strconv.Atoi(data)

    if err != nil {
        log.Fatal(err)
    }

    return length
}

func getAnswer(time []string, distance []string) int{
    index, total := len(time), 0
    var winAmount []int;

    for i := 0; i < index; i++ {
        possibleWin := 0
        for currentSec := 0; currentSec < stringToInt(time[i]); currentSec++ {
            remainingSec := stringToInt(time[i]) - currentSec
            moveSec := stringToInt(time[i]) - remainingSec
            if moveSec * remainingSec > stringToInt(distance[i]) {
                possibleWin ++
            }
        }
        winAmount = append(winAmount, possibleWin)
        possibleWin = 0
    }

    for _, v := range winAmount {
        if total == 0 {
            total = v
        } else {
            total = total * v
        }
    }

    return total
}

func main() {
    data := getData()
    time, distance := formatData(data)
    answer := getAnswer(time, distance)

    fmt.Println("The answer for day 06 is:", answer)
}
