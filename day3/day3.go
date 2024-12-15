package main

import (
	"fmt"
    "strconv"
    "io"
    "os"
)

type DoType int
const (
    INVALID DoType = iota
    DO
    DONT
)

func readInput() (string) {
    bytes, _ := io.ReadAll(os.Stdin)
    return string(bytes)

}

func isNumber(x byte) bool {
    return x >= 48 && x <= 57
}

func parseDoDont(input string) (success bool, doType DoType) {

    success = false
    doType = INVALID

    i := 0

    if(i+7 < len(input) && input[:i+7] == "don't()") {
        doType = DONT
        success = true
        return
    }

    if(i+4 < len(input) && input[:i+4] == "do()") {
        doType = DO
        success = true
        return
    }

    return

}

func parseMulExpr(input string) (success bool, x int, y int) {

    // Initialize return values
    x = 0
    y = 0
    success = false

    i := 0 
    j := 0

    xStr := ""
    yStr := ""


    if !isNumber(input[i]) {
        return
    }
    i++
    for ;;i++ {
        if !isNumber(input[i]) {
            break
        }
    } 

    xStr = input[:i]

    if !(input[i] == ',') {
        return
    }
    i++

    j = i
    if !isNumber(input[i]) {
        return
    }
    i++
    for ;;i++ {
        if !isNumber(input[i]) {
            break
        }
    } 
    yStr = input[j:i]

    if!(input[i] == ')') {
        return
    }

    success = true
    x, _ = strconv.Atoi(xStr)
    y, _ = strconv.Atoi(yStr)

    return
}

func task2(input string) {

    sum := 0

    do := true

    for i, c := range input {
        if(c == 'm') {
            if(i+4 < len(input) && input[i:i+4] == "mul(") {
                success, x, y := parseMulExpr(input[i+4:])
                if success && do {
                    fmt.Printf("Sum += (%d * %d)\n", x, y)
                    sum += x * y
                }           
            }
        }

        if(c == 'd') {
            success, doType := parseDoDont(input[i:])

            if !success {
                continue
            }

            switch(doType) {
            case DO:
                fmt.Println("DO!")
                do = true
            case DONT:
                fmt.Println("DONT!")
                do = false
            }
        }

    }

    fmt.Printf("Total Sum with DO/DONT: %d\n", sum)

}

func task1(input string) {

    sum := 0

    for i, c := range input {
        if(c == 'm') {
            if(i+4 < len(input) && input[i:i+4] == "mul(") {
                success, x, y := parseMulExpr(input[i+4:])
                if success {
                    fmt.Printf("Sum += (%d * %d)\n", x, y)
                    sum += x * y
                }           
            }
        }

    }

    fmt.Printf("Total Sum: %d\n", sum)

}

func main() {
    //input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
    input := readInput()
    fmt.Println("Day 3!")

    task1(input)
    task2(input)
}
