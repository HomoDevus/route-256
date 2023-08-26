package main

import (
	"bufio"
	"fmt"
	"os"
)

type Process struct {
	value int
	time  int
}

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

type ProcessNode struct {
	Value Process
	Next  *ProcessNode
}

type ProcessLinkedList struct {
	Head *ProcessNode
	Tail *ProcessNode
}

func main() {
	var processesAmount, tasksAmount int
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	fmt.Fscan(in, &processesAmount, &tasksAmount)

	processesEnergyConsumption := make([]int, processesAmount)

	for processInedx := 0; processInedx < processesAmount; processInedx++ {
		fmt.Fscan(in, &processesEnergyConsumption[processInedx])
	}

	processesEnergyConsumption = mergeSort(processesEnergyConsumption)
	processesConsumptionStack := LinkedList{}

	for i := len(processesEnergyConsumption) - 1; i >= 0; i-- {
		processesConsumptionStack.Insert(processesEnergyConsumption[i])
	}

	// Use only together. If data changed in one, it must be change in other also.
	processesEnding := ProcessLinkedList{}
	energyConsumption := 0

	for taskIndex := 0; taskIndex < tasksAmount; taskIndex++ {
		var queuedTime, executionTime int

		fmt.Fscan(in, &queuedTime, &executionTime)

		// If there is processes that end on this step restore them
		for processesEnding.Last().time != -1 && queuedTime >= processesEnding.Last().time {
			processesConsumptionStack.Insert(processesEnding.Pop().value)
		}

		// Take lowest process calculate consumption and add process index to busy proceses map with a key of index when to stop
		mostEfficientProcess := processesConsumptionStack.Pop()

		if mostEfficientProcess != -1 {
			energyConsumption += mostEfficientProcess * executionTime

			processesEnding.Insert(Process{
				value: mostEfficientProcess,
				time:  executionTime + queuedTime,
			})
		}
	}

	fmt.Fprintln(out, energyConsumption)

	defer out.Flush()
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

func (ll *LinkedList) Insert(value int) {
	newNode := &Node{Value: value}

	if ll.Head == nil || ll.Head.Value >= value {
		newNode.Next = ll.Head
		ll.Head = newNode
		return
	}

	curr := ll.Head

	for curr.Next != nil && curr.Next.Value < value {
		curr = curr.Next
	}

	newNode.Next = curr.Next
	curr.Next = newNode
}

func (ll *LinkedList) Pop() int {
	if ll.Head == nil {
		return -1
	}

	value := ll.Head.Value
	ll.Head = ll.Head.Next
	return value
}

func (ll *LinkedList) Last() int {
	if ll.Head == nil {
		return -1
	}

	return ll.Head.Value
}

func (ll *ProcessLinkedList) Insert(value Process) {
	newNode := &ProcessNode{Value: value}

	if ll.Head == nil || ll.Head.Value.time >= value.time {
		newNode.Next = ll.Head
		ll.Head = newNode
		return
	}

	curr := ll.Head

	for curr.Next != nil && curr.Next.Value.time < value.time {
		curr = curr.Next
	}

	newNode.Next = curr.Next
	curr.Next = newNode
}

func (ll *ProcessLinkedList) Pop() Process {
	if ll.Head == nil {
		return Process{
			value: -1,
			time:  -1,
		}
	}

	value := ll.Head.Value
	ll.Head = ll.Head.Next
	return value
}

func (ll *ProcessLinkedList) Last() Process {
	if ll.Head == nil {
		return Process{
			value: -1,
			time:  -1,
		}
	}

	return ll.Head.Value
}
