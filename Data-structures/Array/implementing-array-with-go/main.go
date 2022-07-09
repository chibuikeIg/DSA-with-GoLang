package main

type Array struct {
	Length int
	Data   []interface{}
}

func NewArray() *Array {
	length := 0
	data := []any{}
	return &Array{length, data}
}

func (array Array) Push(item any) {
	array.Data[array.Length] = item
	array.Length++
	return
}

func main() {

	// array := NewArray()

	// array.Push("John")

	// fmt.Println(array)

}
