package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)

	fmt.Fprintln(out, "I am sure that I will fill out the form by 23:59 on August 27, 2023.")

	defer out.Flush()
}
