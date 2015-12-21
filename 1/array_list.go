package base

// ArrayList implements something close to go slice
type ArrayList struct {
	arr     []int
	size    int
	arrSize int
}

// NewArrayList create an ArrayList of specific size
func NewArrayList(size int) ArrayList {
	return ArrayList{make([]int, size), size, size}
}

// Get returns a value from specific index
func (a *ArrayList) Get(index int) int {
	a.checkRange(index)

	return a.arr[index]
}

// Set sets the value under specific index
func (a *ArrayList) Set(index int, val int) {
	a.checkRange(index)

	a.arr[index] = val
}

// Len returns a length of the array
func (a *ArrayList) Len() int {
	return a.size
}

// Append appends a new value to the end of the array
func (a *ArrayList) Append(val int) {
	if a.size == a.arrSize {
		a.arrSize = a.arrSize*2 + 1
		old := a.arr
		a.arr = make([]int, a.arrSize)
		copy(a.arr, old)
	}

	a.size++
	a.arr[a.size-1] = val
}

func (a *ArrayList) checkRange(index int) {
	if index < 0 || index >= a.size {
		panic("out of range")
	}
}
