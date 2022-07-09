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

	address := ht._hash(key)

	if ht.data[address] == nil {
		ht.data[address] = []HashTableBucket{}
	}

	ht.data[address] = append(ht.data[address].([]HashTableBucket), HashTableBucket{key, val})

}

func (ht HashTable) Get(key string) any {

	address := ht._hash(key)

	if ht.data[address] != nil {

		currentBucket := ht.data[address].([]HashTableBucket)

		if len(currentBucket) == 1 {
			return currentBucket[0].val
		}

		for _, bucket := range currentBucket {

			if bucket.key == key {
				return bucket.val
			}
		}

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

}
