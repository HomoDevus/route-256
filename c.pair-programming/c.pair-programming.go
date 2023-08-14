package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
)

const (
	Infinity = math.MaxInt64
)

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

func main() {
	var setsAmount int
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	fmt.Fscan(in, &setsAmount)

	for setIndex := 0; setIndex < setsAmount; setIndex++ {
		var devAmount int

		fmt.Fscan(in, &devAmount)

		devs := make([]int, devAmount)
		// Indexes on the right of the for loop curr position
		// that already have a pair on the left of loop current position
		deletedIndexes := make(map[int]bool)

		// Get list of developers skill values
		for devIndex := 0; devIndex < devAmount; devIndex++ {
			fmt.Fscan(in, &devs[devIndex])
		}
		
		for devIndex := 0; devIndex < devAmount - 1; devIndex++ {
			if deletedIndexes[devIndex] {continue}

			aSkill := devs[devIndex]
			minDiff := Infinity
			var minDiffIndex int

			for secDevIndex := devIndex + 1; secDevIndex < devAmount; secDevIndex++ {
				if deletedIndexes[secDevIndex] {continue}

				bSkill := devs[secDevIndex]
				diff := abs(aSkill - bSkill)

				if diff < minDiff {
					minDiff = diff
					minDiffIndex = secDevIndex
				}
			}

			deletedIndexes[minDiffIndex] = true

			fmt.Fprintln(out, devIndex + 1, minDiffIndex + 1)
		}

		// Print empty line
		if (setIndex + 1 < setsAmount) {
			fmt.Fprintln(out)
		}
	}

	defer out.Flush()
}