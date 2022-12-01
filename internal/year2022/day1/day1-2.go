package day1

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Day1_2 struct{}

func (Day1_2) Solve(content string) {
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
	sort.Slice(elben, func(i, j int) bool {
		return elben[i].sum > elben[j].sum
	})
	sum := elben[0].sum + elben[1].sum + elben[2].sum
	fmt.Printf("%v + %v + %v = %v\n", elben[0].sum, elben[1].sum, elben[2].sum, sum)
}

func (Day1_2) Name() string {
	return "2022:1:2"
}
