package main

import "fmt"

const ARRAY_SIZE = 10

type bucketNode struct {
	key  string
	next *bucketNode
}

func (bn *bucketNode) display() {
	fmt.Printf("%v, ", bn.key)
}

type bucket struct {
	head *bucketNode
}

func (b *bucket) insert(key string) {
	if !b.search(key) {
		newNode := &bucketNode{key: key}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("Already exists in the bucket")
	}
}

func (b *bucket) search(key string) bool {
	currNode := b.head
	for currNode != nil {
		if currNode.key == key {
			return true
		}
		currNode = currNode.next
	}
	return false
}

func (b *bucket) delete(key string) {
	if b.head.key == key {
		b.head = b.head.next
		return
	}
	prevNode := b.head
	for prevNode != nil {
		if prevNode.next.key == key {
			prevNode.next = prevNode.next.next
		}
		prevNode = prevNode.next
	}
}

func (b *bucket) display() {
	currNode := b.head
	for currNode != nil {
		currNode.display()
		currNode = currNode.next
	}
}

type HashTable struct {
	array [ARRAY_SIZE]*bucket
}

func NewHashTable() *HashTable {
	hashTable := &HashTable{}
	for i := range ARRAY_SIZE {
		hashTable.array[i] = &bucket{}
	}
	return hashTable
}

func (ht *HashTable) Insert(key string) {
	idx := hash(key)
	ht.array[idx].insert(key)
}

func (ht *HashTable) Search(key string) bool {
	idx := hash(key)
	return ht.array[idx].search(key)
}

func (ht *HashTable) Delete(key string) {
	idx := hash(key)
	ht.array[idx].delete(key)
}

func (ht *HashTable) Display() {
	for i := range ARRAY_SIZE {
		fmt.Printf("{ %d: ", i)
		ht.array[i].display()
		fmt.Printf("}\n")
	}
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ARRAY_SIZE
}

func main() {
	hashTable := NewHashTable()
	list := []string{
		"hassam",
		"wajahat",
		"kiani",
		"rusty",
		"billi",
	}
	for _, v := range list {
		hashTable.Insert(v)
	}

	hashTable.Display()
}
