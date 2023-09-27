package try

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
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
	hasher := md5.New()
	hasher.Write([]byte(key))
	hashValue := hex.EncodeToString(hasher.Sum(nil))
	valueInt, _ := strconv.ParseInt(hashValue, 16, 32)
	return int(valueInt) % ht.size
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
