package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Cor struct {
	x int
	y int
}

var corX, corY, score, length, width int
var snake []Cor

func addPoint(state [][]int) [][]int {
	var x, y int = rand.Int() % length, rand.Int() % width
	for {
		if state[x][y] == 0 {
			state[x][y] = 8
			break
		} else {
			x, y = rand.Int()%length, rand.Int()%width
		}
	}
	return state
}

func printGame(state [][]int) {
	fmt.Println()
	fmt.Println("Game State:")
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			fmt.Print(state[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println("Controls: w=up, d=right, s=down, a=left")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for length <= 0 || width <= 0 {
		fmt.Print("Enter length and width of field: ")
		fmt.Scan(&length, &width)
		if length <= 0 || width <= 0 {
			fmt.Println("Field dimensions cannot be 0 or negative.")
			fmt.Println()
		}
	}
	var state [][]int
	for i := 0; i < length; i++ {
		state = append(state, make([]int, width))
	}
	state[0][0] = 2
	snake = append(snake,
		Cor{
			x: 0,
			y: 0,
		},
	)
	state = addPoint(state)
	printGame(state)
	for {
		var move string
		fmt.Print("Enter move: ")
		fmt.Scan(&move)
		corX = snake[len(snake)-1].x
		corY = snake[len(snake)-1].y
		switch move {
		case "w", "W":
			corX--
		case "s", "S":
			corX++
		case "d", "D":
			corY++
		case "a", "A":
			corY--
		default:
			fmt.Println("Invalid input, type again")
			continue
		}
		if corX >= length || corX < 0 || corY >= width || corY < 0 {
			break
		}
		if state[corX][corY] == 8 {
			state = addPoint(state)
			score++
			if len(snake) > 0 {
				state[snake[0].x][snake[0].y] = 6
			}
		} else {
			if state[corX][corY] == 1 {
				break
			}
			state[snake[0].x][snake[0].y] = 0
			snake = snake[1:]
			if len(snake) > 0 {
				state[snake[0].x][snake[0].y] = 6
			}
		}
		state[corX][corY] = 2
		snake = append(snake,
			Cor{
				x: corX,
				y: corY,
			},
		)
		if len(snake) > 1 && state[snake[len(snake)-2].x][snake[len(snake)-2].y] != 6 {
			state[snake[len(snake)-2].x][snake[len(snake)-2].y] = 1
		}
		printGame(state)
	}
	fmt.Println()
	fmt.Println("Game Over.")
	fmt.Println("Final Score: ", score)
}
