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
	// batteryCount == 2 for part 1
	// batteryCount == 12 for part 2
	batteryCount := 12
	bankStream, countChan := startBatteryBankStream()
	joltageStream := make(chan int)
	for i := 0; i < degreeOfFanOut; i++ {
		startBankComputer(bankStream, joltageStream, batteryCount)
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

func startBankComputer(bankChan <-chan string, outputChan chan<- int, batteryCount int) {
	go func(bc <-chan string, oc chan<- int) {
		for {
			bank, ok := <-bc
			if !ok {
				break
			}
			oc <- highestJoltageInBank(bank, batteryCount)
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

func highestJoltageInBank(bank string, batteryCount int) int {
	var source []int
	for _, dc := range bank {
		d, _ := strconv.Atoi(string(dc))
		source = append(source, d)
	}

	accum := highest(source, batteryCount, nil)
	s := ""
	for _, i := range accum {
		s += strconv.Itoa(i)
	}
	sum, _ := strconv.Atoi(s)
	return sum

	//highest, pos := 0, 0
	//
	//for i := 0; i < len(bank)-1; i++ {
	//	p, _ := strconv.Atoi(string(bank[i]))
	//	if p > highest {
	//		highest = p
	//		pos = i
	//	}
	//}
	//
	//nextHighest := 0
	//for i := pos + 1; i < len(bank); i++ {
	//	p, _ := strconv.Atoi(string(bank[i]))
	//	if p > nextHighest {
	//		nextHighest = p
	//	}
	//}
	//
	//return (highest * 10) + nextHighest
}

func highest(source []int, rem int, accum []int) []int {
	h, p := 0, 0
	lim := len(source) - rem
	for i := 0; i <= lim; i++ {
		if source[i] > h {
			h = source[i]
			p = i
		}
	}

	accum = append(accum, h)
	r := rem - 1
	if r == 0 {
		return accum
	}

	var remSource []int
	for i := p + 1; i < len(source); i++ {
		remSource = append(remSource, source[i])
	}

	return highest(remSource, r, accum)
}

//func highestJoltageInBankLength(bank string, length int) int {
//	var highest []int
//	pos := -1
//	for {
//		rem := len(bank) - pos - 1
//		h := 0
//		for i := pos + 1; i <= len(bank)-rem; i++ {
//			d, _ := strconv.Atoi(string(bank[i]))
//			if d > h {
//				h = d
//				pos = i
//			}
//		}
//		highest = append(highest, h)
//		if len(highest) == length {
//			break
//		}
//	}
//
//	joltage := 0
//	for i := 0; i < len(highest); i++ {
//		f := 1
//		for j := i + 1; j < (length - i); j++ {
//			f = f * 10
//		}
//		joltage += highest[i] * f
//	}
//	return joltage
//}
