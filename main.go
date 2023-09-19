package main

import (
	try "backend/try/hash"
	"fmt"
)

func main() {
	hashTable := try.NewHashTable(10)
	hashTable.Add("key1", "value1")
	hashTable.Add("key2", "value2")
	hashTable.Add("key3", "value3")

	hashTable.Print()

	fmt.Println(hashTable.Get("key1"))
	fmt.Println(hashTable.Get("key2"))
	fmt.Println(hashTable.Get("key3"))
	fmt.Println(hashTable.Get("key4"))
}
