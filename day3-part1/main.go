package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

const degreeOfFanOut = 10

func main() {
	bankStream, countChan := startBatteryBankStream()
	joltageStream := make(chan int)
	for i := 0; i < degreeOfFanOut; i++ {
		startBankComputer(bankStream, joltageStream)
	}
	resultChan := startResultComputer(joltageStream, countChan)

	result := <-resultChan

	fmt.Println(result)
}

func startBatteryBankStream() (bankStream <-chan string, countChan <-chan int) {
	bStream := make(chan string)
	cChan := make(chan int)

	go func(ch chan<- string, countCh chan<- int) {
		count := 0
		defer close(ch)
		defer func() {
			countCh <- count
		}()
		for _, line := range strings.Split(input, "\n") {
			ch <- line
			count++
		}

	}(bStream, cChan)

	return bStream, cChan
}

func startBankComputer(bankChan <-chan string, outputChan chan<- int) {
	go func(bc <-chan string, oc chan<- int) {
		for {
			bank, ok := <-bc
			if !ok {
				break
			}
			oc <- highestJoltageInBank(bank)
		}
	}(bankChan, outputChan)
}

func startResultComputer(joltageChan <-chan int, countChan <-chan int) <-chan int {
	resultChan := make(chan int)

	go func(jc <-chan int, cc <-chan int, rc chan<- int) {
		sum, count, goalCount := 0, 0, 0
		defer func() {
			rc <- sum
		}()
		for {
			select {
			case c := <-cc:
				goalCount += c
			case j := <-jc:
				sum += j
				count++
				if goalCount > 0 && count >= goalCount {
					return
				}
			}
		}

	}(joltageChan, countChan, resultChan)

	return resultChan
}

func highestJoltageInBank(bank string) int {
	highest, pos := 0, 0

	for i := 0; i < len(bank)-1; i++ {
		p, _ := strconv.Atoi(string(bank[i]))
		if p > highest {
			highest = p
			pos = i
		}
	}

	nextHighest := 0
	for i := pos + 1; i < len(bank); i++ {
		p, _ := strconv.Atoi(string(bank[i]))
		if p > nextHighest {
			nextHighest = p
		}
	}

	return (highest * 10) + nextHighest
}
