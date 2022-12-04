package day4

import (
	"fmt"
	"github.com/Soemii/AdventOfCode/internal/util"
	"strings"
)

type Group struct {
	section1Start uint
}

type Day4_1 struct {
	counter int
}

func (d Day4_1) Solve(content string) {
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
			for i := array[0]; i <= array[1]; i++ {
				if si == 0 {
					sectionOne = append(sectionOne, i)
				} else {
					sectionTwo = append(sectionTwo, i)
				}
			}
		}
		if sectionOne[0] >= sectionTwo[0] && sectionOne[len(sectionOne)-1] <= sectionTwo[len(sectionTwo)-1] {
			d.counter++
		} else if sectionOne[0] <= sectionTwo[0] && sectionOne[len(sectionOne)-1] >= sectionTwo[len(sectionTwo)-1] {
			d.counter++
		}
	}
	fmt.Println(d.counter)
}

func (d Day4_1) Name() string {
	return "2022:4:1"
}
