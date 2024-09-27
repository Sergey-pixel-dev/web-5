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
	var str string
	fmt.Scan(&str)
	go removeDuplicates(in, out)
	go func() {
		defer close(in)
		for _, i := range str {
			in <- string(i)
		}
	}()
	for i := range out {
		fmt.Print(i)
	}

}
