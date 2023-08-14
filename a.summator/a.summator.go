package main

import (
	"fmt"
)

func main() {
	var n int

	_, err := fmt.Scan(&n)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i := 0; i < n; i++ {
		var a, b int

		_, err := fmt.Scan(&a, &b)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println(a + b)
	}
}
