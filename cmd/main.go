package main

import (
	"flag"
	"fmt"
	"github.com/Soemii/AdventOfCode/internal/util"
	"github.com/Soemii/AdventOfCode/internal/year2022/day1"
	"github.com/Soemii/AdventOfCode/internal/year2022/day2"
	"github.com/Soemii/AdventOfCode/internal/year2022/day3"
	"github.com/Soemii/AdventOfCode/internal/year2022/day4"
	"log"
)

func main() {
	year := flag.String("year", "2022", "set the year")
	day := flag.String("day", "1", "set the day")
	challenge := flag.String("challenge", "1", "set the challenge")
	example := flag.Bool("example", false, "load example data")
	flag.Parse()
	path := fmt.Sprintf("./data/%v/day%v.txt", *year, *day)
	if *example {
		path = fmt.Sprintf("./data/%v/example%v.txt", *year, *day)
	}
	challengeName := fmt.Sprintf("%v:%v:%v", *year, *day, *challenge)
	fileContent, err := util.ReadFile(path)
	if err != nil {
		log.Println("Cannot read file!")
		return
	}
	challenges := registerChallenges()
	ch, ok := challenges[challengeName]
	if !ok {
		log.Println("Cannot load challenge")
		return
	}
	ch.Solve(fileContent)
}

func registerChallenges() map[string]util.Challenge {
	challenges := make(map[string]util.Challenge)
	registerChallenge(day1.Day1_1{}, challenges)
	registerChallenge(day1.Day1_2{}, challenges)
	registerChallenge(day2.Day2_1{}, challenges)
	registerChallenge(day2.Day2_2{}, challenges)
	registerChallenge(day3.Day3_1{}, challenges)
	registerChallenge(day3.Day3_2{}, challenges)
	registerChallenge(day4.Day4_1{}, challenges)
	registerChallenge(day4.Day4_2{}, challenges)
	return challenges
}

func registerChallenge(challenge util.Challenge, challenges map[string]util.Challenge) {
	if challenge, ok := challenges[challenge.Name()]; ok {
		log.Printf("Challenge %v already registered!", challenge.Name())
		return
	}
	challenges[challenge.Name()] = challenge
}
