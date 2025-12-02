package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

const degreeOfFanOut = 10

//go:embed input.txt
var input string

func main() {
	checkedIDs := make(chan int)
	inputStream, countChan := startInputStream()
	//inputStream, countChan := startTestInputStream(1188511880, 1188511890)
	for i := 0; i < degreeOfFanOut; i++ {
		startChecker(inputStream, checkedIDs)
	}

	resultChan := startResultStream(checkedIDs, countChan)

	res := <-resultChan

	fmt.Printf("Manipulation complete! Result: %v\n", res)

	// part1 -> 28146997880
}

func startTestInputStream(start, end int) (idStream, countStream <-chan int) {
	output := make(chan int, 1000)
	cntChan := make(chan int)

	go func(ch, cch chan<- int) {
		defer close(ch)
		count := 0
		defer func() {
			fmt.Println(count)
			cch <- count
		}()
		for i := start; i <= end; i++ {
			ch <- i
			count++
		}
	}(output, cntChan)

	return output, cntChan
}

func startInputStream() (idStream, countChan <-chan int) {
	output := make(chan int, 1000)
	cntChan := make(chan int)

	go func(ch, cch chan<- int) {
		defer close(ch)
		count := 0
		defer func() {
			fmt.Println(count)
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
			outChan <- checkId(id)
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

func checkId(id int) int {
	ids := fmt.Sprintf("%d", id)
	lim := len(ids) / 2
	for i := 1; i <= lim; i++ {
		if isInvalidStep(ids, i) {
			return id
		}
	}

	return 0
}

func isInvalidStep(ids string, step int) bool {
	l := len(ids)
	if l%step != 0 {
		return false
	}

	check := ids[0:step]

	i := step
	for {
		if i+step > l {
			break
		}

		candidate := ids[i : i+step]
		if candidate != check {
			return false
		}
		i += step
	}

	return true
}
