package base

// Element represents a single element in array
type Element interface{}

// ArrayList implements something close to go slice
type ArrayList struct {
	arr     []Element
	size    int
	arrSize int
}

// NewArrayList create an ArrayList of specific size
func NewArrayList(size int) ArrayList {
	return ArrayList{make([]Element, size), size, size}
}

// Get returns a value from specific index
func (a *ArrayList) Get(index int) Element {
	a.checkRange(index)

	return a.arr[index]
}

// Set sets the value under specific index
func (a *ArrayList) Set(index int, val interface{}) {
	a.checkRange(index)

	a.arr[index] = Element(val)
}

// Len returns a length of the array
func (a *ArrayList) Len() int {
	return a.size
}

// Append appends a new value to the end of the array
func (a *ArrayList) Append(val interface{}) {
	if a.size == a.arrSize {
		a.arrSize = a.arrSize*2 + 1
		old := a.arr
		a.arr = make([]Element, a.arrSize)
		copy(a.arr, old)
	}

	a.size++
	a.arr[a.size-1] = Element(val)
}

func (a *ArrayList) checkRange(index int) {
	if index < 0 || index >= a.size {
		panic("out of range")
	}
}
