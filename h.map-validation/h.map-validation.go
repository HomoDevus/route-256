package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var setsAmount int
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	fmt.Fscan(in, &setsAmount)

	for setIndex := 0; setIndex < setsAmount; setIndex++ {
		var columns, rows int

		fmt.Fscan(in, &columns, &rows)

		gameMap := matrixInput(columns, rows, in)
		checkedCells := make([]map[int]bool, columns*rows)
		checkedColors := make(map[rune]bool)

		for i := range checkedCells {
			checkedCells[i] = make(map[int]bool)
		}

		result := true

		for col := 0; col < columns && result; col++ {
			for row := 0; row < rows && result; row++ {
				color := gameMap[col][row]

				if !checkedCells[col][row] {
					if checkedColors[color] {
						result = false
					} else {
						result = checkCell(col, row, gameMap, checkedCells, checkedColors)
						checkedColors[color] = true
					}
				}
			}
		}

		if result {
			fmt.Fprint(out, "YES")
		} else {
			fmt.Fprint(out, "NO")
		}

		if setIndex != setsAmount-1 {
			fmt.Fprintln(out)
		}
	}

	defer out.Flush()
}

func checkCell(col, row int, gameMap [][]rune, checkedCells []map[int]bool, checkedColors map[rune]bool) bool {
	result := true
	colsAmount := len(gameMap)
	rowsAmount := len(gameMap[col])
	color := gameMap[col][row]
	grid := [][]int{
		{-1, -1},
		{-1, 1},
		{0, 2},
		{1, 1},
		{1, -1},
		{0, -2},
	}

	if checkedCells[col][row] {
		return result
	}

	if checkedColors[color] {
		return false
	}

	checkedCells[col][row] = true

	for i := 0; i < len(grid) && result; i++ {
		currCol := col + grid[i][0]
		currRow := row + grid[i][1]

		if currCol < 0 || currCol >= colsAmount || currRow < 0 || currRow >= rowsAmount || checkedCells[currCol][currRow] {
			continue
		}

		currColor := gameMap[currCol][currRow]

		if color == currColor {
			result = checkCell(currCol, currRow, gameMap, checkedCells, checkedColors)
		}
	}

	return result
}

func matrixInput(cols int, rows int, in io.Reader) [][]rune {
	matrix := make([][]rune, cols)

	for col := 0; col < cols; col++ {
		matrix[col] = make([]rune, rows)

		for row := 0; row < rows; row++ {
			var pattern string

			if row == 0 {
				pattern = "\n%c"
			} else {
				pattern = "%c"
			}

			fmt.Fscanf(in, pattern, &matrix[col][row])
		}
	}

	return matrix
}
