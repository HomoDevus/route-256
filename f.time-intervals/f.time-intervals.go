package main

import (
	"bufio"
	"fmt"
	"os"
)

type Time struct {
	hour int
	min  int
	sec  int
}

type TimeInterval struct {
	start Time
	end   Time
}

func main() {
	var setsAmount int
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	fmt.Fscan(in, &setsAmount)

	for setIndex := 0; setIndex < setsAmount; setIndex++ {
		var intervalsAmount int
		var isCorrect = true

		fmt.Fscan(in, &intervalsAmount)

		intervals := make([]TimeInterval, intervalsAmount)

		for intervalIndex := 0; intervalIndex < intervalsAmount; intervalIndex++ {
			var time TimeInterval

			fmt.Fscanf(in, "\n%d:%d:%d-%d:%d:%d", &time.start.hour, &time.start.min, &time.start.sec, &time.end.hour, &time.end.min, &time.end.sec)

			if !isCorrect {
				continue
			}

			isCorrect = checkTimeDigitsCorrectness(time.start) && checkTimeDigitsCorrectness(time.end) && checkIntervalCorrectness(time)

			intervals[intervalIndex] = time
		}

		if isCorrect {
			intervals = mergeSort(intervals)

			for i := 1; i < len(intervals) && isCorrect; i++ {
				isCorrect = checkNoIntersections(intervals[i - 1], intervals[i])
			}
		}

		if isCorrect {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}

	defer out.Flush()
}

func checkTimeDigitsCorrectness(time Time) bool {
	return time.hour >= 0 && time.hour <= 23 && time.min >= 0 && time.min <= 59 && time.sec >= 0 && time.sec <= 59
}

func checkIntervalCorrectness(time TimeInterval) bool {
	return time.start.hour < time.end.hour || (time.start.hour == time.end.hour && time.start.min < time.end.min) || (time.start.hour == time.end.hour && time.start.min == time.end.min && time.start.sec <= time.end.sec)
}

func checkNoIntersections(timeA, timeB TimeInterval) bool {
	return isTimeBefore(timeA.end, timeB.start) || isTimeBefore(timeB.end, timeA.start)
}

func isTimeBefore(timeA, timeB Time) bool {
	return timeA.hour < timeB.hour || (timeA.hour == timeB.hour && timeA.min < timeB.min) || (timeA.hour == timeB.hour && timeA.min == timeB.min && timeA.sec < timeB.sec)
}

func mergeSort(arr []TimeInterval) []TimeInterval {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []TimeInterval) []TimeInterval {
	result := make([]TimeInterval, 0, len(left)+len(right))
	lIdx, rIdx := 0, 0

	for lIdx < len(left) && rIdx < len(right) {
		if isTimeBefore(left[lIdx].start, right[rIdx].start) {
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
