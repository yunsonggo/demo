package demo

import "github.com/yunsonggo/demo/pr"

func NewPrDemo(name string) *pr.Demo {
	return pr.New(name)
}
