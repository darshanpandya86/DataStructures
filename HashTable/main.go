package main

import "fmt"

const ArraySize = 7

//HashTable structure
type HashTable struct {
	array [ArraySize]*bucket
}

//Insert into Hashtable
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

//Search in hashtable
/*func (h *HashTable) Search(key string) bool {
	index := hash(key)
}

//Delete
func (h *HashTable) Delete(key string) {
	index := hash(key)
}*/

//bucket structure
type bucket struct {
	head *bucketNode
}
type bucketNode struct {
	key  string
	next *bucketNode
}

//insert in bucket
func (b *bucket) insert(k string) {
	newNode := &bucketNode{key: k}
	newNode.next = b.head
	b.head = newNode
}

//search
//delete

//hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

//Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	testHashtable := Init()
	fmt.Println(testHashtable)
	fmt.Println(hash("RANDY"))

	testBucket := &bucket{}
	testBucket.insert("RANDY")
}
