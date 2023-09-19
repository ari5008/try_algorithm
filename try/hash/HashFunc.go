package try

import (
    "crypto/md5"
    "fmt"
)

type HashTable struct {
    size  int
    table [][][2]interface{}
}

func NewHashTable(size int) *HashTable {
    return &HashTable{
        size:  size,
        table: make([][][2]interface{}, size),
    }
}

func (ht *HashTable) Hash(key string) int {
    hashBytes := md5.Sum([]byte(key))
    hashInt := 0
    for _, b := range hashBytes {
        hashInt = hashInt<<8 | int(b)
    }
    return hashInt % ht.size
}

func (ht *HashTable) Add(key string, value interface{}) {
	index := ht.Hash(key)
	if ht.table[index] == nil {
			ht.table[index] = make([][2]interface{}, 0)
	}
	for i, data := range ht.table[index] {
			if data[0].(string) == key {
					ht.table[index][i][1] = value
					return
			}
	}
	ht.table[index] = append(ht.table[index], [2]interface{}{key, value})
}

func (ht *HashTable) Get(key string) interface{} {
    index := ht.Hash(key)
    for _, data := range ht.table[index] {
        if data[0].(string) == key {
            return data[1]
        }
    }
    return nil
}

func (ht *HashTable) Print() {
	for index := 0; index < ht.size; index++ {
			fmt.Print(index, " ")
			for _, data := range ht.table[index] {
					fmt.Printf("--> %v ", data[1])
			}
			fmt.Println()
	}
}
