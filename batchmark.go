package gotest

func MakeSliceWithoutAlloc(size int) []int {
	var newSlice []int
	for i := 0; i < size; i++ {
		newSlice = append(newSlice, i)
	}
	return newSlice
}

func MakeSliceWitPrevAlloc(size int) []int {
	newSlice := make([]int, 0, size)
	for i := 0; i < size; i++ {
		newSlice = append(newSlice, i)
	}
	return newSlice
}
