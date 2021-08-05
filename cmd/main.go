package main

import (
	"fmt"
	"lottery/lottery"
)

// 单抽概率
var possibility = lottery.Possibility{
	"A": 9,
	"B": 15,
	"C": 40,
	"D": 87,
	"E": 211,
	"F": 1230,
	"G": 23303,
	"H": 29332,
	"I": 45773,
}
var sample = lottery.NewSample3(possibility)

// 十连抽概率
var possibility10 = lottery.Possibility{
	"B": 2,
	"C": 8,
	"D": 30,
	"E": 960,
}
var sample10 = lottery.NewSample3(possibility10)

// 百连抽概率
var possibility100 = lottery.Possibility{
	"F": 5,
	"G": 95,
}
var sample100 = lottery.NewSample3(possibility100)

var score = 0

func main() {
	gots := lot(10)
	fmt.Println(gots)
	fmt.Println(score)
}

func lot(times int) []string {
	gots := make([]string, 0, times)
	if times == 10 {
		got := sample10.Lot()
		gots = append(gots, got)
		times--
		score = score + 11
	} else if times == 100 {
		got := sample100.Lot()
		gots = append(gots, got)
		times--
		score = score + 115
	} else {
		score = score + 1
	}
	for i := 0; i < times; i++ {
		got := sample.Lot()
		gots = append(gots, got)
	}

	if score > 400 {
		score = score % 400
	}

	return gots
}
