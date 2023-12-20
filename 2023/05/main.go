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

func formatData(data *bufio.Scanner) ([]string, map[int][][]string, []int) {

    almanac := [][]string {}
    almanacMap := map[int][][]string {}
    keys := []int {}
    index := 0
    var seeds []string

    for data.Scan() {
        if strings.Split(data.Text(), " ")[0] == "seeds:" {
            seeds = strings.Split(data.Text(), " ")[1:]
        } else if strings.Contains(data.Text(), ":") && almanac != nil {
            if index != 0 {
                almanacMap[index] = almanac
                almanac = [][]string {}
                keys = append(keys, index)
            }
            index += 1
        } else if data.Text() != ""{
            almanac = append(almanac, strings.Split(data.Text(), " "))
        } else if index == 7 {
            almanacMap[index] = almanac
            keys = append(keys, index)
        }
    }

    return seeds, almanacMap, keys
}

func stringToInt(data string) int {

    length, err := strconv.Atoi(data)

    if err != nil {
        log.Fatal(err)
    }

    return length
}

func getLocationValue(data *bufio.Scanner) []string {
    seeds, almanacMap, keys := formatData(data)

    for _, value := range keys {
        match := map[int]int {}
        for _, row := range almanacMap[value] {
            destinationRange := stringToInt(row[0])
            sourceRange := stringToInt(row[1])
            rangeLength := stringToInt(row[2])

            for index, seed := range seeds {
                seedInt := stringToInt(seed)
                _, ok := match[index]
                if seedInt >= sourceRange && seedInt <= sourceRange + rangeLength && !ok {
                   match[index] = destinationRange + (seedInt - sourceRange)
                }
            }
        }

        for i, v := range match {
            seeds[i] = strconv.Itoa(v)
        }
    }

    return seeds
}

func getAnswer(seeds []string) int{
    counter := 0;
    var value int;

    for _, v := range seeds {
        if counter == 0 {
            value = stringToInt(v)
            counter ++
        }

        if value > stringToInt(v) {
            value = stringToInt(v)
        }
    } 

    return value
}

func main() {
    data := getData()
    seeds := getLocationValue(data)
    value := getAnswer(seeds)

    fmt.Println("The answer for day 05 is:", value)
}
