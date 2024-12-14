package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const debug = false

func readInput() ([][]int) {

    out := make([][]int, 2)
    out[0] = make([]int, 0)
    out[1] = make([]int, 0)

    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        string := strings.Split(scanner.Text(), "   ")
        x, _ := strconv.Atoi(string[0])
        y, _ := strconv.Atoi(string[1])
        out[0] = append(out[0], x)
        out[1] = append(out[1], y)
    }

    return out

}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}

func task1(input [][]int) {
    sortX := input[0]
    sortY := input[1]

    start := time.Now()
    sort.Ints(sortX)
    sort.Ints(sortY)
    elapsed := time.Since(start)

    fmt.Printf("Sorting time: %s\n", elapsed)

    accumulatedDistance := 0

    for i := 0; i < len(sortX); i++ {
       accumulatedDistance += abs(sortX[i] - sortY[i])
       if debug {
           fmt.Printf("acc: prev + (%d - %d) = %d\n", sortX[i], sortY[i], accumulatedDistance)
       }
    }

    fmt.Printf("Accumulated distance: %d\n", accumulatedDistance)
}

func task2(input [][]int) {
    numMap := make(map[int]int)

    for _, y := range input[1] {
        numMap[y] += 1
    }
    
    accSimilarityScore := 0
    for _, x := range input[0] {
        accSimilarityScore += x * numMap[x]
    }

    fmt.Printf("Accumulated similarity score: %d\n", accSimilarityScore)
}

func main() {

    input := readInput()

    if ( len(input[0]) != len(input[1])) {
        panic("Input lists does not have a matching number of items!")
    }

    task1(input)
    task2(input)

}
