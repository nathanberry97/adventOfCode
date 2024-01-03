package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var total int

func getData() [][]string {
    file, err := os.Open("input.txt")

    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)

    var grid [][]string

    for scanner.Scan() {
        grid = append(grid, strings.Split(scanner.Text(), ""))
    }

    return grid
}

func getStartCoordinates(grid [][]string) []int {
    var x, y int

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == "S" {
                x = i
                y = j
            }
        }
    }

    return []int{x, y}
}

var NEIGHBOURS = map[string][][]int{
    "|": {{1, 0}, {-1, 0}},
    "-": {{0, 1}, {0, -1}},
    "L": {{-1, 0}, {0, 1}},
    "J": {{-1, 0}, {0, -1}},
    "7": {{1, 0}, {0, -1}},
    "F": {{1, 0}, {0, 1}},
    ".": {},
}

func main() {
    var grid [][]string
    var startCoordinates []int
    
    grid = getData()
    startCoordinates = getStartCoordinates(grid)

    fmt.Println("Start coordinates: ", grid[startCoordinates[0]][startCoordinates[1]])
}
