package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type ConstraintList []int
type ConstraintMap map[int]ConstraintList

// Add a constraint: beforeValue must be printed before afterValue
func addConstraint(beforeConstraints ConstraintMap, beforeValue int, afterValue int) {
    beforeConstraints[beforeValue] = append(beforeConstraints[beforeValue], afterValue)
}

func readRules(text string) (before int, after int) {
    vals := strings.Split(text, "|")

    before, _= strconv.Atoi(vals[0])
    after, _= strconv.Atoi(vals[1])

    return
}

func readPrintOrder(text string) []int {

    vals := strings.Split(text, ",")

    printOrderSlice := make([]int, len(vals))

    for i, v := range vals {
        printOrderSlice[i], _= strconv.Atoi(v)
    }

    return printOrderSlice
}

func readInput() (ConstraintMap, [][]int){
    beforeConstraints := make(ConstraintMap)
    printOrder := make([][]int, 0)

    scanner := bufio.NewScanner(os.Stdin)

    scanningRules := true

    for scanner.Scan() {
        if(scanner.Text() == "") {
            scanningRules = false
            continue
        }

        switch scanningRules {
        case true:
            before, after := readRules(scanner.Text())
            addConstraint(beforeConstraints, before, after)
        case false:
            printOrder = append(printOrder, readPrintOrder(scanner.Text()))
        }
    }

    return beforeConstraints, printOrder
}

func checkConstraintsFulfilled(beforeConstraints ConstraintMap, print []int) (bool) {
    for i, num := range print {

        for _, prevNum := range print[:i] {

            if slices.Contains(beforeConstraints[num], prevNum) {
                return false

            }
        }
    }

    return true
}

func fixBrokenConstraints(beforeConstraints ConstraintMap, print []int) []int {
    newPrint := make([]int, len(print))
    copy(newPrint, print)

    i := 0
    for i < len(newPrint) {
        beforeList := make([]int, 0)
        beforeList = append(beforeList, newPrint[:i]...)

        afterList := make([]int, 0)

        outer := newPrint[i]

        for _, inner := range newPrint[i+1:] {

            if slices.Contains(beforeConstraints[inner], outer) {
                // Inner elements that should go before outer element according to
                // constraints are moved in front of it
                beforeList = append(beforeList, inner)
            } else {
                // Inner elements that do not have a constraint on outer element go 
                // after it in the same order
                afterList = append(afterList, inner)
            }
        }

        if len(beforeList) > i {
            newPrint = append(beforeList, outer)
            newPrint = append(newPrint, afterList...)
        } else {
            // No elements moved due to constraint, check next index
            i++
        }

    }

    if !checkConstraintsFulfilled(beforeConstraints, newPrint) {
        panic("CONSTRAINTS STILL NOT FULFILLED")
    }



    return newPrint
}

func task1(beforeConstraints ConstraintMap, printOrder [][]int) {

    accValidMiddles := 0

    for _, print := range printOrder {
        if checkConstraintsFulfilled(beforeConstraints, print) {
            accValidMiddles += print[len(print) / 2]
        }
    }

    fmt.Printf("Accumulated valid middle prints: %d\n", accValidMiddles)
}

func task2(beforeConstraints ConstraintMap, printOrder [][]int)  {

    accFixedMiddles := 0

    for _, print := range printOrder {
        if !checkConstraintsFulfilled(beforeConstraints, print) {
            fixedPrint := fixBrokenConstraints(beforeConstraints, print)
            accFixedMiddles += fixedPrint[len(fixedPrint) / 2]
        }
    }

    fmt.Printf("Accumulated fixed middle prints: %d\n", accFixedMiddles)

}

func main() {
    beforeConstraints, printOrder := readInput()

    task1(beforeConstraints, printOrder)
    task2(beforeConstraints, printOrder)

}
