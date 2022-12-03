package day3

import (
	"fmt"
	"strings"
)

type Day3_2 struct {
	points int
}

type Group struct {
	lines []string
}

func (d Day3_2) Solve(content string) {
	lines := strings.Split(content, "\n")
	groups := make([]Group, 0)
	for i := 0; i < len(lines); i += 3 {
		group := Group{}
		group.lines = lines[i : i+3]
		groups = append(groups, group)
	}
	for _, group := range groups {
		first := group.lines[0]
		second := group.lines[1]
		third := group.lines[2]
		var contains rune
		for _, i := range first {
			if strings.Contains(second, string(i)) && strings.Contains(third, string(i)) {
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

func (Day3_2) Name() string {
	return "2022:3:2"
}
