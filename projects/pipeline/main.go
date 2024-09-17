package main

import "fmt"

// реализовать removeDuplicates(in, out chan string)

func removeDuplicates(in, out chan string) {
	defer close(out)
	var lastString string
	for i := range in {
		if i != lastString {
			out <- i
		}
		lastString = i
	}
}

func main() {
	var in = make(chan string)
	var out = make(chan string)
	go removeDuplicates(in, out)
	go func() {
		defer close(in)
		for _, i := range "1112222223333333444444" {
			in <- string(i)
		}
	}()
	for i := range out {
		fmt.Print(i)
	}

}
