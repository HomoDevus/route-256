package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	SaleFloor = 3
)

func main() {
	var setsAmount int
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	fmt.Fscan(in, &setsAmount)

	for i := 0; i < setsAmount; i++ {
		var goodsAmount int
		priceSum := 0
		goods := make(map[int]int)

		fmt.Fscan(in, &goodsAmount)

		for i := 0; i < goodsAmount; i++ {
			var goodPrice int

			fmt.Fscan(in, &goodPrice)

			if _, exists := goods[goodPrice]; !exists {
				goods[goodPrice] = 0
			}

			goods[goodPrice] += 1
		}

		for price, amount := range goods {
			amount -= amount / SaleFloor
			priceSum += amount * price
		}

		fmt.Fprintln(out, priceSum)
	}

	defer out.Flush()
}
