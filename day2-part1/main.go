package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

const countIDCheckers = 10

//go:embed input.txt
var input string

func main() {
	checkedIDs := make(chan int)
	inputStream, countChan := startInputStream()
	for i := 0; i < countIDCheckers; i++ {
		startChecker(inputStream, checkedIDs)
	}

	resultChan := startResultStream(checkedIDs, countChan)

	result := <-resultChan

	fmt.Println(result)
}

func startInputStream() (idStream, countChan <-chan int) {
	output := make(chan int)
	cntChan := make(chan int)

	go func(ch, cch chan<- int) {
		defer close(ch)
		count := 0
		defer func() {
			cch <- count
		}()
		lines := strings.Split(input, ",")
		for _, line := range lines {
			arr := strings.Split(line, "-")
			f, _ := strconv.Atoi(arr[0])
			t, _ := strconv.Atoi(arr[1])
			for i := f; i <= t; i++ {
				ch <- i
				count++
			}
		}
	}(output, cntChan)

	return output, cntChan
}

func startChecker(incoming <-chan int, checkedIDs chan<- int) {
	go func(inChan <-chan int, outChan chan<- int) {
		for {
			id, ok := <-inChan
			if !ok {
				break
			}
			outChan <- checkID(id)
		}
	}(incoming, checkedIDs)
}

func startResultStream(checkedIDs, countChan <-chan int) <-chan int {
	output := make(chan int)

	go func(idChan, cChan <-chan int, rChan chan<- int) {
		sum, count, goalCount := 0, 0, 0
		defer func() {
			rChan <- sum
		}()

		for {
			select {
			case id := <-idChan:
				sum += id
				count++
				if goalCount > 0 && count >= goalCount {
					return
				}
			case gc := <-countChan:
				goalCount = gc
			}
		}

	}(checkedIDs, countChan, output)

	return output
}

func checkID(id int) int {
	s := fmt.Sprintf("%d", id)
	if len(s)%2 != 0 {
		return 0
	}
	d := len(s) / 2
	if s[0:d] == s[d:] {
		return id
	}

	return 0
}
