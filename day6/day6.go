package main

import (
	"bufio"
	"fmt"
	"os"
)

type Puzzle struct {
    s string
    Width int
    Height int

    ExtraObstacle Pos
}

// Position in 2D space
type Pos struct {
    X, Y int
}

// A transition between one position to another
type Edge struct {
    From, To Pos
}

// Set datastructure. Non-generic, only supports Pos keys
type Set struct {
    s map[Pos]bool
}

func (s Set) Add(p Pos) {
    _, ok := s.s[p]
    if !ok {
        s.s[p] = true
    }
}

func (s Set) Length() int {
    return len(s.s)
}

func (s Set) GetKeys() []Pos {
    keys := make([]Pos, len(s.s))
    i := 0
    for k := range s.s {
        keys[i] = k
        i++
    }
    return keys
}

func makeSet() Set {
    return Set{ make(map[Pos]bool) }
}

func (p Puzzle) GetAt (pos Pos) (byte) {
    if  (pos.X < 0 || pos.X >= p.Width) || (pos.Y < 0 || pos.Y >= p.Height) {
        return ' '
    }

    if pos == p.ExtraObstacle {
        return 'O'
    }

    return p.s[p.Width * pos.Y + pos.X]
}

func (p Puzzle) FindGuard() (Pos) {
    for y := 0; y < p.Height; y++ {
        for x := 0; x < p.Width; x++ {
            if p.GetAt(Pos{x, y}) == '^' {
                return Pos{x, y}
            }
        }

    }
    return Pos{-1, -1}
}

func (p Puzzle) PrintPuzzle(center Pos, radius int) {
    for y := -radius; y < radius; y++ {
        for x := -radius; x < radius; x++ {
            if x == 0 && y == 0 {
                fmt.Printf("^")
                continue
            }
            square := p.GetAt(Pos{x + center.X, y + center.Y})
            fmt.Printf("%c", square)
        }
        fmt.Printf("\n")
    }

}

func newPuzzle (input string, width int, height int) Puzzle {
    return Puzzle { input, width, height, Pos {-1, -1}}
}


func rotateClockwise(dirX int, dirY int) (int, int) {
    return -dirY, dirX
}

func translate(pos Pos, dirX int, dirY int) (Pos) {
    return Pos{pos.X + dirX, pos.Y + dirY}
}

func runSimulation(puzzle Puzzle) (Set, bool) {
    gPos := puzzle.FindGuard()
    gDirX := 0
    gDirY := -1

    set := makeSet()
    set.Add(gPos)

    hareIndex := 0
    tortoiseIndex := 0
    tortoiseMove := 0

    visitedEdges := make([]Edge, 0)

    var last *Pos = nil
    
    for {

        frontPos := translate(gPos, gDirX, gDirY)

        switch tile := puzzle.GetAt(frontPos); {
        case tile == '#' || tile == 'O' :
            if last != nil {
                visitedEdges = append(visitedEdges, Edge{ *last, frontPos })
            }
            last = &frontPos

            hareIndex = len(visitedEdges) - 1
            if len(visitedEdges) > 0 {
                if tortoiseMove == 1 {
                    if visitedEdges[tortoiseIndex] == visitedEdges[hareIndex] {
                        // Loop found (hare caught up to tortoise)
                        return set, true
                    }
                    tortoiseIndex++
                    tortoiseMove = 0
                } else {
                    tortoiseMove++
                }
            }
            gDirX, gDirY = rotateClockwise(gDirX, gDirY)
            continue

        case tile == ' ':
            // Outside map
            goto EndLoop
        }

        gPos = translate(gPos, gDirX, gDirY)
        set.Add(gPos)

    }
    EndLoop:


    return set, false
}

func task1(puzzle Puzzle) {
    uniquePoses, _ := runSimulation(puzzle)
    fmt.Printf("Unique positions visited, task 1: %d\n", uniquePoses.Length())
}

func task2(puzzle Puzzle) {
    // Get initial set of positions the guard can visit
    guardLocations, _ := runSimulation(puzzle)

    gPos := puzzle.FindGuard()

    accTimeloops := 0
    accIterations := 0

    for _, pos := range guardLocations.GetKeys() {
        if pos == gPos {
            continue
        }

        // Add an extra obstacle at one of the positions the guard can visit
        puzzle.ExtraObstacle = pos

        _, timeloop := runSimulation(puzzle)

        if timeloop {
            accTimeloops++
        }
        accIterations++
    }
    fmt.Printf("Number of timeloops found: %d out of total %d\n", accTimeloops, accIterations)
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

func main() {

    input, width, height := readInput()
    puzzle := newPuzzle(input, width, height)

    task1(puzzle)
    task2(puzzle)

}
