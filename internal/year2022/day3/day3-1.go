package day3

import (
	"fmt"
	"strings"
)

type Day3_1 struct {
	points int
}

func (d Day3_1) Solve(content string) {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		length := len(line) / 2
		first := line[:length]
		second := line[length:]
		var contains rune
		for _, i := range first {
			if strings.Contains(second, string(i)) {
				contains = i
				continue
			}
		}
		var points int32
		if int32(contains) < 95 {
			//Großbuchstaben
			points = (int32(contains) - 65) + 27
		} else {
			//Kleinbuchstaben
			points = int32(contains) - 96
		}
		d.points += int(points)
	}
	fmt.Println(d.points)
}

func (Day3_1) Name() string {
	return "2022:3:1"
}
