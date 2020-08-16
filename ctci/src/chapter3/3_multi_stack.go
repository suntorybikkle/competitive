package chapter3

import (
	"fmt"
)

type stackInfo struct {
	start       int
	size        int
	capacity    int
	valueLength int
}

func newStackInfo(start, capacity, valueLength int) *stackInfo {
	return &stackInfo{start: start, capacity: capacity, valueLength: valueLength}
}

func (info *stackInfo) isWithinStackCapacity(index int) bool {
	valueLength := info.valueLength
	if index < 0 || index >= valueLength {
		return false
	}
	start := info.start
	var contiguousIndex int
	if start > index {
		contiguousIndex = index + valueLength
	} else {
		contiguousIndex = index
	}
	end := start + info.capacity
	return start <= contiguousIndex && contiguousIndex < end
}

func (info *stackInfo) lastCapacityIndex() int {
	return info.start + info.capacity - 1
}

func (info *stackInfo) lastElementIndex() int {
	return info.start + info.size - 1
}

func (info *stackInfo) isFull() bool {
	return info.size == info.capacity
}

func (info *stackInfo) isEmpty() bool {
	return info.size == 0
}

type MultiStack struct {
	info   []*stackInfo
	values []int
}

func NewMultiStack(numberOfStacks, defaultSize int) *MultiStack {
	values := make([]int, numberOfStacks*defaultSize)
	info := make([]*stackInfo, numberOfStacks)
	for i := range info {
		info[i] = newStackInfo(i*defaultSize, defaultSize, len(values))
	}
	return &MultiStack{info: info, values: values}
}

func (multiStack *MultiStack) Push(stackNum, value int) (err error) {
	if multiStack.AllStacksAreFull() {
		return new(FullStackError)
	}
	stack := multiStack.info[stackNum]
	if stack.isFull() {
		multiStack.expand(stackNum)
	}
	stack.size++
	adjustedIndex := multiStack.adjustIndex(stack.lastElementIndex())
	multiStack.values[adjustedIndex] = value
	return
}

func (multiStack *MultiStack) Pop(stackNum int) (int, error) {
	stack := multiStack.info[stackNum]
	if stack.isEmpty() {
		return 0, new(EmptyStackError)
	}
	adjustedIndex := multiStack.adjustIndex(stack.lastElementIndex())
	value := multiStack.values[adjustedIndex]
	multiStack.values[adjustedIndex] = 0
	stack.size--
	return value, nil
}

func (multiStack *MultiStack) Peek(stackNum int) int {
	stack := multiStack.info[stackNum]
	adjustedIndex := multiStack.adjustIndex(stack.lastElementIndex())
	return multiStack.values[adjustedIndex]
}

func (multiStack *MultiStack) shift(stackNum int) {
	fmt.Println("/// Shifting ", stackNum)
	stack := multiStack.info[stackNum]

	if stack.size >= stack.capacity {
		nextStack := (stackNum + 1) % len(multiStack.info)
		multiStack.shift(nextStack)
		stack.capacity++
	}
	index := multiStack.adjustIndex(stack.lastCapacityIndex())
	values := multiStack.values
	for stack.isWithinStackCapacity(index) {
		values[index] = values[multiStack.previousIndex(index)]
		index = multiStack.previousIndex(index)
	}
	values[stack.start] = 0
	stack.start = multiStack.nextIndex(stack.start)
	stack.capacity--
}

func (multiStack *MultiStack) expand(stackNum int) {
	multiStack.shift((stackNum + 1) % len(multiStack.info))
	multiStack.info[stackNum].capacity++
}

func (multiStack *MultiStack) NumberOfElements() int {
	size := 0
	for _, sd := range multiStack.info {
		size += sd.size
	}
	return size
}

func (multiStack *MultiStack) AllStacksAreFull() bool {
	return len(multiStack.values) == multiStack.NumberOfElements()
}

func (multiStack *MultiStack) adjustIndex(index int) int {
	max := len(multiStack.values)
	return ((index % max) + max) % max
}

func (multiStack *MultiStack) nextIndex(index int) int {
	return multiStack.adjustIndex(index + 1)
}

func (multiStack *MultiStack) previousIndex(index int) int {
	return multiStack.adjustIndex(index - 1)
}
