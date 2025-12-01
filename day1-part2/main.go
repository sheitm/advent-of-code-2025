package main

import (
	_ "embed"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type rotationDirection int

const (
	rotationSet rotationDirection = iota
	rotationRight
	rotationLeft

	startingPoint = 50
	dialSize      = 100
)

//go:embed input.txt
var input string

type dialManipulation struct {
	direction rotationDirection
	clicks    int
}

func main() {
	inputStream := startManipulationInputStream(startingPoint)
	wrapsCountChan := startDial(inputStream, dialSize)
	resultChan := startOutputter(wrapsCountChan)

	result := <-resultChan

	fmt.Println(result)
}

func startManipulationInputStream(sPoint int) <-chan dialManipulation {
	output := make(chan dialManipulation)

	go func(ch chan<- dialManipulation) {
		defer close(ch)
		ch <- dialManipulation{direction: rotationSet, clicks: sPoint}

		lines := strings.Split(input, "\n")
		for _, line := range lines {
			manipulation, err := splitLine(line)
			if err != nil {
				panic(err)
			}
			ch <- manipulation
		}

	}(output)

	return output
}

func startDial(input <-chan dialManipulation, base int) (zeroChan <-chan int) {
	zeroChannel := make(chan int)

	go func(zc chan<- int) {
		position := 0
		for {
			op, ok := <-input
			if !ok {
				close(zeroChannel)
				return
			}
			wraps := 0
			switch op.direction {
			case rotationSet:
				position = op.clicks % base
				continue
			case rotationRight:
				for i := 0; i < op.clicks; i++ {
					position++
					if position == base {
						position = 0
						wraps++
					}
				}
			case rotationLeft:
				for i := 0; i < op.clicks; i++ {
					position--
					if position == 0 {
						wraps++
						continue
					}
					if position < 0 {
						position = base - 1
					}
				}
			}
			zc <- wraps
		}
	}(zeroChannel)

	return zeroChannel
}

func startOutputter(zeroStream <-chan int) (resultChan <-chan int) {
	rc := make(chan int)
	sum := 0
	go func(stream <-chan int) {
		defer func() {
			rc <- sum
		}()
		for {
			wraps, ok := <-stream
			if !ok {
				return
			}
			sum += wraps
		}
	}(zeroStream)

	return rc
}

func splitLine(line string) (dialManipulation, error) {
	var direction rotationDirection
	switch line[0] {
	case 'R', 'r':
		direction = rotationRight
	case 'L', 'l':
		direction = rotationLeft
	default:
		return dialManipulation{}, errors.New("invalid rotationDirection type")
	}

	arg, err := strconv.Atoi(line[1:])
	if err != nil {
		return dialManipulation{}, err
	}

	return dialManipulation{direction: direction, clicks: arg}, nil
}
