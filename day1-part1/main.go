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
)

type dialManipulation struct {
	direction rotationDirection
	clicks    int
}

//go:embed input.txt
var input string

func (o dialManipulation) String() string {
	op := ""
	switch o.direction {
	case rotationSet:
		op = "set"
	case rotationRight:
		op = "right"
	case rotationLeft:
		op = "left"
	}

	return fmt.Sprintf("%s %d", op, o.clicks)
}

func main() {
	inputStream := startManipulationInputStream(50)
	zeroCountChan := startDial(inputStream)
	resultChan := startOutputter(zeroCountChan)

	result := <-resultChan

	fmt.Println(result)
}

func startDial(input <-chan dialManipulation) (zeroChan <-chan struct{}) {
	zeroChannel := make(chan struct{})
	go func(ch <-chan dialManipulation, zChan chan<- struct{}) {
		position := 0
		for {
			op, ok := <-ch
			if !ok {
				close(zeroChannel)
				return
			}
			switch op.direction {
			case rotationSet:
				position = op.clicks % 100
			case rotationRight:
				position = (position + op.clicks) % 100
			case rotationLeft:
				position = (position - op.clicks%100 + 100) % 100
			}

			if position == 0 {
				zChan <- struct{}{}
			}
		}
	}(input, zeroChannel)

	return zeroChannel
}

func startManipulationInputStream(startPoint int) <-chan dialManipulation {
	output := make(chan dialManipulation)

	lines := strings.Split(input, "\n")

	go func(ch chan<- dialManipulation) {
		defer close(ch)
		ch <- dialManipulation{direction: rotationSet, clicks: startPoint}

		for _, line := range lines {
			op, err := splitLine(line)
			if err != nil {
				panic(err)
			}
			ch <- op
		}

	}(output)

	return output
}

func startOutputter(zeroStream <-chan struct{}) (resultChan <-chan int) {
	rc := make(chan int)
	count := 0
	go func(stream <-chan struct{}) {
		defer func() {
			rc <- count
		}()
		for {
			_, ok := <-stream
			if !ok {
				return
			}
			count++
		}
	}(zeroStream)

	return rc
}

func splitLine(line string) (dialManipulation, error) {
	var ot rotationDirection
	switch line[0] {
	case 'R', 'r':
		ot = rotationRight
	case 'L', 'l':
		ot = rotationLeft
	default:
		return dialManipulation{}, errors.New("invalid dialManipulation type")
	}

	arg, err := strconv.Atoi(line[1:])
	if err != nil {
		return dialManipulation{}, err
	}

	return dialManipulation{direction: ot, clicks: arg}, nil
}
