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
		var cols, rows, clicksAmount int

		_, err := fmt.Fscan(in, &cols, &rows)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		table := matrixInput(cols, rows, in)

		fmt.Fscan(in, &clicksAmount)

		for clickIndex := 0; clickIndex < clicksAmount; clickIndex++ {
			var sortCol int

			fmt.Fscan(in, &sortCol)
			sortCol--

			// Bubble sorting algorithm
			for i := 0; i < cols; i++ {
				swapped := false

				for col := 0; col < cols-i-1; col++ {
					if table[col][sortCol] > table[col+1][sortCol] {
						table[col], table[col+1] = table[col+1], table[col]
						swapped = true
					}
				}

				if !swapped {
					break
				}
			}
		}

		matrixOutput(cols, rows, table, out)

		if setIndex+1 < setsAmount {
			fmt.Fprintln(out)
		}
	}

	defer out.Flush()
}

func matrixInput(cols int, rows int, in io.Reader) [][]int {
	matrix := make([][]int, cols)

	for col := 0; col < cols; col++ {
		matrix[col] = make([]int, rows)

		for row := 0; row < rows; row++ {
			fmt.Fscan(in, &matrix[col][row])
		}
	}

	return matrix
}

func matrixOutput(cols int, rows int, matrix [][]int, out io.Writer) {
	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			fmt.Fprint(out, matrix[col][row])

			if row+1 < rows {
				fmt.Fprint(out, " ")
			}
		}

		fmt.Fprintln(out)
	}
}
