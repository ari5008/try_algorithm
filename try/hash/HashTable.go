package try

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

type HashTable struct {
	size  int
	table [][][2]string
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		size:  size,
		table: make([][][2]string, size),
	}
}

func (ht *HashTable) hash(key string) int {
	hasher := sha256.New()
	hasher.Write([]byte(key))
	hashBytes := hasher.Sum(nil)
	hashValue := binary.BigEndian.Uint32(hashBytes[:4])
	return int((hashValue % 10) + 1)
}

func (ht *HashTable) Add(key string, value string) {
	index := ht.hash(key)
	for _, data := range ht.table[index] {
		if data[0] == key {
			data[1] = value
			return
		}
	}
	ht.table[index] = append(ht.table[index], [2]string{key, value})
}

func (ht *HashTable) Print() {
	for index := 0; index < ht.size; index++ {
		fmt.Printf("%d ", index)
		for _, data := range ht.table[index] {
			fmt.Print("--> ")
			fmt.Print(data)
		}
		fmt.Println()
	}
}

func (ht *HashTable) Get(key string) string {
	index := ht.hash(key)
	for _, data := range ht.table[index] {
		if data[0] == key {
			return data[1]
		}
	}
	return ""
}
