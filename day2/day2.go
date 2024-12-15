package main

import (
    "fmt"
    "bufio"
    "strconv"
    "os"
    "strings"
    "slices"
)

const maxDist = 3

func readInput() ([][]int) {

    out := make([][]int, 0)

    scanner := bufio.NewScanner(os.Stdin)

    i := 0

    for scanner.Scan() {
        //out[i] = make([]int, 0)
        out = append(out, make([]int, 0))
        string := strings.Split(scanner.Text(), " ")

        for _, s := range string {
            x, _ := strconv.Atoi(s)
            out[i] = append(out[i], x)
        }
        i++
    }

    return out

}

func abs(x int) (int) {
    if x < 0 {
        return -x
    }
    return x
}

func safeCompare(diff int, increasing bool) bool {
    return (diff >= 0) == increasing && diff != 0 && abs(diff) <= maxDist 
}

func determineSafe(report []int, allowCorrection bool) bool {

    increasing := (report[1] - report[0]) >= 0

    for i, level := range report[1:] {

        diff := level - report[i]
        if !safeCompare(diff, increasing) {
            if allowCorrection {
                for j := range report {
                    if determineSafe(slices.Concat(report[:j], report[j+1:]), false) {
                        return true
                    }
                }
            }

            return false
        }
    }

    return true
}

func task1(input [][]int) {
    accSafe := 0

    for _, report := range input {
        if determineSafe(report, false) {
            accSafe += 1
        }
    }

    fmt.Printf("Accumulated safe reports: %d\n", accSafe)

}

func task2(input [][]int) {
    accSafe := 0

    for _, report := range input {

        if determineSafe(report, true) {
            accSafe += 1
        } 

    }

    fmt.Printf("Accumulated safe reports with problem dampening: %d\n", accSafe)

}

func main() {
    input := readInput()

    task1(input)
    task2(input)

}
