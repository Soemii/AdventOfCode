package util

type Challenge interface {
	Solve(content string)
	Name() string
}
