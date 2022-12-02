package day2

import (
	"fmt"
	"strings"
)

type Day2_1 struct {
	points int
}

func (d Day2_1) Solve(content string) {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		split := strings.Split(line, " ")
		me := SchereSteinPapier{}
		elb := SchereSteinPapier{}
		switch split[0] {
		case "A":
			elb.stein = true
		case "B":
			elb.papier = true
		case "C":
			elb.schere = true
		}
		switch split[1] {
		case "X":
			me.stein = true
			d.points += 1
		case "Y":
			me.papier = true
			d.points += 2
		case "Z":
			me.schere = true
			d.points += 3
		}
		if me.win(elb) {
			d.points += 6
		} else if me.draw(elb) {
			d.points += 3
		}
	}
	fmt.Println(d.points)
}

func (Day2_1) Name() string {
	return "2022:2:1"
}
