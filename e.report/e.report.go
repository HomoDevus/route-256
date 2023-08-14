package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var setsAmount int
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	fmt.Fscan(in, &setsAmount)

	for setIndex := 0; setIndex < setsAmount; setIndex++ {
		var reportsAmount, prevTask, currentTask int
		var isSuccessful = true
		switchedTasks := make(map[int]bool)

		fmt.Fscan(in, &reportsAmount)
		fmt.Fscan(in, &prevTask)

		switchedTasks[prevTask] = true

		for reportIndex := 0; reportIndex < reportsAmount - 1; reportIndex++ {
			fmt.Fscan(in, &currentTask)

			if !isSuccessful {
				continue
			}

			if currentTask != prevTask {
				if switchedTasks[currentTask] {
					isSuccessful = false
				} else {
					switchedTasks[currentTask] = true
				}
			}

			prevTask = currentTask
		}

		if isSuccessful {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}

	defer out.Flush()
}
