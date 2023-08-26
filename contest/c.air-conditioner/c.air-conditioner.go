package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	var dataSetsAmount int

	fmt.Fscan(in, &dataSetsAmount)

	for dataSetIndex := 0; dataSetIndex < dataSetsAmount; dataSetIndex++ {
		var peopleAmount int
		minTemp := 15
		maxTemp := 30

		fmt.Fscan(in, &peopleAmount)

		for peopleIndex := 0; peopleIndex < peopleAmount; peopleIndex++ {
			var symbol string
			var temp int

			fmt.Fscanf(in, "\n%s ", &symbol)
			fmt.Fscan(in, &temp)

			if symbol[0] == '>' && minTemp < temp {
				minTemp = temp
			} else if symbol[0] == '<' && maxTemp > temp {
				maxTemp = temp
			}

			if minTemp > maxTemp {
				fmt.Fprintln(out, -1)
			} else {
				fmt.Fprintln(out, maxTemp)
			}
		}

		fmt.Fprintln(out)
	}

	defer out.Flush()
}