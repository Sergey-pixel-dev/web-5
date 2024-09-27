package main

import "fmt"

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	answerChan := make(chan int)
	go func() {
		defer close(answerChan)
		select {
		case val := <-firstChan:
			answerChan <- val * 2
		case val := <-secondChan:
			answerChan <- val * 3
		case <-stopChan:
		}
	}()
	return answerChan
}
func main() {
	var firstChan = make(chan int, 1)
	var secondChan = make(chan int, 1)
	var stopChan = make(chan struct{}, 1)
	firstChan <- 5
	fmt.Println(<-calculator(firstChan, secondChan, stopChan))
}
