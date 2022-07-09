package main

import (
	"fmt"
	"unicode/utf8"
)

type HashTableBucket struct {
	key string
	val any
}

type HashTable struct {
	data []any
}

func NewHashTable(size int) *HashTable {

	data := make([]any, size)

	return &HashTable{data}

}

func (ht HashTable) _hash(key string) int {

	hash := 0

	for i := 0; i < len(key); i++ {
		hash = (hash + int(charCode(key, i))*i) % len(ht.data)
	}

	return hash

}

func (ht HashTable) Set(key string, val any) {

	ht.data[ht._hash(key)] = HashTableBucket{key, val}

}

func (ht HashTable) Get(key string) any {

	if ht.data[ht._hash(key)] != nil {
		data := ht.data[ht._hash(key)].(HashTableBucket)
		return data.val
	}

	return nil

}

func charCode(str string, index int) rune {

	if indexSum := index + 1; indexSum <= len(str) {

		r, _ := utf8.DecodeRuneInString(str[index:indexSum])

		return r
	}

	return 0

}

func main() {

	myHashTable := NewHashTable(50)

	myHashTable.Set("grapes", 10000)

	myHashTable.Set("name", "peter")

	fmt.Println(myHashTable.Get("grapes"))

	fmt.Println(myHashTable._hash("grapes"))
}
