package day4

import (
	"fmt"
	"github.com/Soemii/AdventOfCode/internal/util"
	"strings"
)

type Day4_2 struct {
	counter int
}

func (d Day4_2) Solve(content string) {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		sections := strings.Split(line, ",")
		sectionOne := make([]int, 0)
		sectionTwo := make([]int, 0)
		for si, section := range sections {
			array, err := util.MapStringArrayToIntArray(strings.Split(section, "-"))
			if err != nil {
				panic(err)
			}
			if si == 0 {
				sectionOne = array
			} else {
				sectionTwo = array
			}
		}

		fmt.Println(sectionOne)
		fmt.Println(sectionTwo)
		fmt.Println(sectionTwo[0] <= sectionOne[1])
		fmt.Println(sectionTwo[0] >= sectionOne[0])
		if sectionOne[0] >= sectionTwo[0] && sectionOne[1] <= sectionTwo[1] {
			d.counter++
			fmt.Println("COUNTER")
		} else if sectionOne[0] <= sectionTwo[0] && sectionOne[1] >= sectionTwo[1] {
			d.counter++
			fmt.Println("COUNTER2")
		} else if sectionOne[0] <= sectionTwo[1] && sectionOne[0] >= sectionTwo[0] {
			d.counter++
			fmt.Println("COUNTER3")
		} else if sectionTwo[0] <= sectionOne[1] && sectionTwo[0] >= sectionOne[0] {
			d.counter++
			fmt.Println("COUNTER4")
		}
	}
	fmt.Println(d.counter)
}

func (d Day4_2) Name() string {
	return "2022:4:2"
}
