package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Expression struct {
    Target int
    Terms []int
}

func readInput() []Expression {

    scanner := bufio.NewScanner(os.Stdin)

    exprList := make([]Expression, 0)

    for scanner.Scan() {

        txt := scanner.Text()

        targetTermsSlice := strings.Split(txt, ":")

        target, _ := strconv.Atoi(targetTermsSlice[0])
        strTerms := strings.Split(targetTermsSlice[1], " ")

        terms := make([]int, len(strTerms))

        for i, str := range strTerms {
            terms[i], _ = strconv.Atoi(str)
        }

        exprList = append(exprList, Expression{ target, terms })
    }

    return exprList
}

func orOperator(x int, y int) int {
    xStr := strconv.Itoa(x)
    yStr := strconv.Itoa(y)


    resStr := xStr + yStr

    res, _ := strconv.Atoi(resStr)
    return res
}

func findSolutionWithOr(acc int, terms []int, target int) (bool, int) {
    if acc > target {
        return false, 0
    }

    if(len(terms) == 1) {

        switch target {
        case acc * terms[0]:
            return true, acc * terms[0]
        case acc + terms[0]:
            return true, acc + terms[0]
        case orOperator(acc, terms[0]):
            return true, orOperator(acc, terms[0])
        }

        return false, 0

    }

    if solved, result := findSolutionWithOr(acc * terms[0], terms[1:], target); solved {
        return solved, result
    } else if solved, result := findSolutionWithOr(acc + terms[0], terms[1:], target); solved {
        return solved, result
    } else if solved, result := findSolutionWithOr(orOperator(acc, terms[0]), terms[1:], target); solved {
        return solved, result
    }

    return false, 0
}

func findSolutionWithoutOr(acc int, terms []int, target int) (bool, int) {
    if acc > target {
        return false, 0
    }

    if(len(terms) == 1) {

        switch target {
        case acc * terms[0]:
            return true, acc * terms[0]
        case acc + terms[0]:
            return true, acc + terms[0]
        }

        return false, 0

    }

    if solved, result := findSolutionWithoutOr(acc * terms[0], terms[1:], target); solved {
        return solved, result
    } else if solved, result := findSolutionWithoutOr(acc + terms[0], terms[1:], target); solved {
        return solved, result
    }

    return false, 0
}

func solve(expr Expression, allowOr bool) (bool, int){
    target := expr.Target
    terms := expr.Terms

    if allowOr {
        return findSolutionWithOr(terms[0], terms[1:], target)
    }  

    return findSolutionWithoutOr(terms[0], terms[1:], target)

}

func task1(exprList []Expression) {

    total := 0
    for _, expr := range exprList {
        _, result := solve(expr, false)

        // fmt.Printf("(%d) Solved: %v, Result: %d\n", expr.Target, solved, result)

        total += result

    }

    fmt.Printf("Total (only addition and multiplication): %d\n", total)
}

func task2(exprList []Expression) {
    total := 0
    for _, expr := range exprList {
        _, result := solve(expr, true)

        // fmt.Printf("(%d) Solved: %v, Result: %d\n", expr.Target, solved, result)

        total += result

    }

    fmt.Printf("Total (with \"or\" operator): %d\n", total)
}


func main() {

    exprList := readInput()

    task1(exprList)
    task2(exprList)
}
