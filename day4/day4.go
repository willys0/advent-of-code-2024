package main

import (
    "fmt"
    "bufio"
    "os"
)

type Puzzle struct {
    s string
    Width int
    Height int
}

func (p Puzzle) GetAt (x int, y int) (byte) {
    if  (x < 0 || x >= p.Width) || (y < 0 || y >= p.Height) {
        return ' '
    }

    return p.s[p.Width * y + x]
}

func newPuzzle (input string, width int, height int) Puzzle {
    return Puzzle { input, width, height }
}

func readInput() (text string, width int, height int) {

    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        text += scanner.Text()
        width = len(scanner.Text())
        height += 1
    }

    return 

}

func searchOneDirForMas(puzzle Puzzle, x int, y int, xDir int, yDir int) bool {
    // fmt.Printf("%c %c %c\n", 
    //         puzzle.GetAt(x + xDir, y + yDir), 
    //         puzzle.GetAt(x + xDir * 2, y + yDir * 2), 
    //         puzzle.GetAt(x + xDir * 3, y + yDir * 3))

    return  puzzle.GetAt(x + xDir, y + yDir) == 'M' &&
            puzzle.GetAt(x + xDir * 2, y + yDir * 2) == 'A' &&
            puzzle.GetAt(x + xDir * 3, y + yDir * 3) == 'S'
}

func searchAllDirsForXMAS (puzzle Puzzle, x int, y int) int {
    acc := 0
    
    for i := -1; i <= 1; i++ {
        for j := -1; j <= 1; j++ {

            if i == 0 && j == 0 {
                continue
            }

            if searchOneDirForMas(puzzle, x, y, i, j) {
                acc += 1
            }
        }
    }

    return acc
}

func searchAllDirsForMasCross (puzzle Puzzle, x int, y int) bool {
    acc := 0
    
    for i := -1; i <= 1; i++ {
        for j := -1; j <= 1; j++ {

            if i == 0 || j == 0 {
                continue
            }

            if searchOneDirForMas(puzzle, x - i * 2, y - j * 2, i, j) {
                acc += 1
            }
        }
    }

    return acc >= 2
}

func task1(puzzle Puzzle) {

    accNumXmases := 0

    for y := 0; y < puzzle.Height; y++ {
        for x := 0; x < puzzle.Width; x++ {
            if puzzle.GetAt(x, y) == 'X' {
                accNumXmases += searchAllDirsForXMAS(puzzle, x, y)
            }
        }
    }

    fmt.Printf("Found in %d XMASes in total\n", accNumXmases)
}

func task2(puzzle Puzzle) {
    accNumMases := 0

    for y := 0; y < puzzle.Height; y++ {
        for x := 0; x < puzzle.Width; x++ {
            if puzzle.GetAt(x, y) == 'A' {
                if(searchAllDirsForMasCross(puzzle, x, y)) {
                    accNumMases += 1
                }
            }
        }
    }

    fmt.Printf("Found %d Mas:es shaped as cross in total\n", accNumMases)

}


func main() {
    // input := "MMMSXXMASM" +
    //          "MSAMXMSMSA" +
    //          "AMXSXMAAMM" +
    //          "MSAMASMSMX" +
    //          "XMASAMXAMM" +
    //          "XXAMMXXAMA" +
    //          "SMSMSASXSS" +
    //          "SAXAMASAAA" +
    //          "MAMMMXMMMM" +
    //          "MXMXAXMASX"
    // width := 10
    // height := 10

    input, width, height := readInput()

    puzzle := newPuzzle(input, width, height)

    task1(puzzle)
    task2(puzzle)
}

