package day1

import (
	"log"
	"strconv"
	"strings"
)

type Day1_1 struct{}

func (Day1_1) Solve(content string) {
	lines := strings.Split(content, "\n")
	elben := make([]Elbe, 0)
	lastElbe := Elbe{}
	for _, line := range lines {
		if strings.Compare(line, "") == 0 {
			sum := 0
			for _, number := range lastElbe.numbers {
				sum += number
			}
			lastElbe.sum = sum
			elben = append(elben, lastElbe)
			lastElbe = Elbe{}
			continue
		}
		atoi, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		lastElbe.numbers = append(lastElbe.numbers, atoi)
	}
	best := Elbe{}
	for _, elbe := range elben {
		if elbe.sum > best.sum {
			best = elbe
		}
	}
	log.Println(best.sum)
}

func (Day1_1) Name() string {
	return "2022:1:1"
}
