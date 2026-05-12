package main

import "fmt"

func main() {
	lrucache := NewLRUCache[string, string](3)
	lrucache.Put("1", "hassam")
	lrucache.Put("2", "waji")
	lrucache.Put("3", "wajahat")
	lrucache.Put("4", "waji")

	fmt.Println(lrucache.Get("3"))

	fmt.Println(lrucache.Get("1"))

	fmt.Println(lrucache.Get("4"))

	fmt.Println(lrucache.Size())
}
