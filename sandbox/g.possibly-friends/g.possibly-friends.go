package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	var usersAmount, pairsAmount int

	fmt.Fscan(in, &usersAmount, &pairsAmount)

	users := make([]map[int]bool, usersAmount)

	for i := range users {
		users[i] = make(map[int]bool)
	}

	for pairIndex := 0; pairIndex < pairsAmount; pairIndex++ {
		var userA, userB int

		fmt.Fscan(in, &userA, &userB)

		users[userA-1][userB-1] = true
		users[userB-1][userA-1] = true
	}

	for userIndex := 0; userIndex < usersAmount; userIndex++ {
		maxIntersectionsLen := 0
		var commonFirends []int
		secondLevelFriends := make(map[int]bool)

		for friend := range users[userIndex] {
			for friendOfFriend := range users[friend] {
				if !secondLevelFriends[friendOfFriend] && userIndex != friendOfFriend && !users[friendOfFriend][userIndex] {
					secondLevelFriends[friendOfFriend] = true
				}
			}
		}

		for friendOfFriend := range secondLevelFriends {
			intersectionLen := countIntersectionLen(users[userIndex], users[friendOfFriend])

			if intersectionLen > maxIntersectionsLen {
				maxIntersectionsLen = intersectionLen
				commonFirends = []int{friendOfFriend}
			} else if intersectionLen > 0 && intersectionLen == maxIntersectionsLen {
				commonFirends = append(commonFirends, friendOfFriend)
			}
		}

		if len(commonFirends) == 0 {
			fmt.Fprint(out, 0)
		} else {
			for _, num := range mergeSort(commonFirends) {
				fmt.Fprint(out, num+1)

				fmt.Fprint(out, " ")
			}
		}

		fmt.Fprintln(out)
	}

	defer out.Flush()
}

func countIntersectionLen(set1, set2 map[int]bool) int {
	intersectionLen := 0

	for key1 := range set1 {
		if set2[key1] {
			intersectionLen++
		}
	}

	return intersectionLen
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	lIdx, rIdx := 0, 0

	for lIdx < len(left) && rIdx < len(right) {
		if left[lIdx] < right[rIdx] {
			result = append(result, left[lIdx])
			lIdx++
		} else {
			result = append(result, right[rIdx])
			rIdx++
		}
	}

	result = append(result, left[lIdx:]...)
	result = append(result, right[rIdx:]...)

	return result
}
