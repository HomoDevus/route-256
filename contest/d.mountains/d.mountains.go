package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	var dataSetsAmount int

	fmt.Fscan(in, &dataSetsAmount)

	for dataSetIndex := 0; dataSetIndex < dataSetsAmount; dataSetIndex++ {
		var reliefsAmount, cols, rows int
		var composedReliefMatrix [][]byte

		fmt.Fscan(in, &reliefsAmount, &cols, &rows)

		for reliefIndex := 0; reliefIndex < reliefsAmount; reliefIndex++ {
			reliefMatrix := matrixInput(cols, rows, in)

			if reliefIndex == 0 {
				composedReliefMatrix = reliefMatrix
			} else {
				mergeReliefs(cols, rows, reliefMatrix, composedReliefMatrix)
			}
		}

		matrixOutput(cols, rows, composedReliefMatrix, out)
		fmt.Fprintln(out)
	}

	defer out.Flush()
}

func mergeReliefs(cols, rows int, currRelief, composedRelief [][]byte) {
	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			currCell := currRelief[col][row]
			composedCell := composedRelief[col][row]

			if composedCell == '.' && currCell != '.' {
				composedRelief[col][row] = currCell
			}
		}
	}
}

func matrixInput(cols int, rows int, in io.Reader) [][]byte {
	matrix := make([][]byte, cols)

	for col := 0; col < cols; col++ {
		matrix[col] = make([]byte, rows)

		for row := 0; row < rows; row++ {
			var cell byte
			cell = 0

			if col == 0 && row == 0 {
				for cell != '.' && cell != '/' {
					fmt.Fscanf(in, "%c", &cell)
				}
			} else {
				fmt.Fscanf(in, "%c", &cell)
			}

			matrix[col][row] = cell
		}

		fmt.Fscanf(in, "\n%c") // Skip new line character
	}

	return matrix
}

func matrixOutput(cols int, rows int, matrix [][]byte, out io.Writer) {
	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			fmt.Fprintf(out, "%c", matrix[col][row])
		}

		fmt.Fprintln(out)
	}
}
