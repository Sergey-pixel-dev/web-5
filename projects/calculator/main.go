package main

import "fmt"

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	answerChan := make(chan int, 1)
	select {
	case val := <-firstChan:
		answerChan <- val * val
		return answerChan
	case val := <-secondChan:
		answerChan <- val * val * val
		return answerChan
	case <-stopChan:
		answerChan <- 0
		return answerChan
	}
}
func main() {
	var firstChan = make(chan int, 1)
	var secondChan = make(chan int, 1)
	var stopChan = make(chan struct{}, 1)
	stopChan <- struct{}{}
	fmt.Println(<-calculator(firstChan, secondChan, stopChan))
}
