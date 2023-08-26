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
	var friendsAmount, cardsAmount int
	var response []int

	fmt.Fscan(in, &friendsAmount, &cardsAmount)

	friendsMaxCard := make([]int, friendsAmount)
	cards := make([]int, cardsAmount)
	hasAnswer := true

	for cardIndex := 0; cardIndex < cardsAmount; cardIndex++ {
		cards[cardIndex] = cardIndex + 1
	}

	for friendIndex := 0; friendIndex < friendsAmount; friendIndex++ {
		fmt.Fscan(in, &friendsMaxCard[friendIndex])
	}

	for friendIndex := 0; friendIndex < friendsAmount && hasAnswer; friendIndex++ {
		friendMax := friendsMaxCard[friendIndex]

		closestCardIndex := binarySearch(cards, friendMax)
		closestCard := cards[closestCardIndex]

		if (closestCard == friendMax && closestCardIndex == len(cards)-1) || closestCard < friendMax {
			hasAnswer = false
		} else if closestCard == friendMax {
			response = append(response, cards[closestCardIndex+1])
			cards = removeItem(cards, closestCardIndex+1)
		} else if closestCard > friendMax {
			response = append(response, closestCard)
			cards = removeItem(cards, closestCardIndex)
		}
	}

	if hasAnswer {
		printArray(response, out)
	} else {
		fmt.Fprintln(out, -1)
	}

	defer out.Flush()
}

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left < right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	for left < len(arr)-1 && arr[left] < target {
		left++
	}

	return left
}

func printArray(arr []int, out io.Writer) {
	for i := 0; i < len(arr); i++ {
		fmt.Fprint(out, arr[i])

		if i != len(arr)-1 {
			fmt.Fprint(out, " ")
		}
	}
}

func removeItem(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}
