package main

import (
	"fmt"
	"lottery/data"
	"lottery/lottery"
	"lottery/util"
	"strconv"
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
	"F": 5,
	"G": 95,
}
var sample10 = lottery.NewSample3(possibility10)

// 百连抽概率
var possibility100 = lottery.Possibility{
	"B": 2,
	"C": 8,
	"D": 30,
	"E": 960,
}
var sample100 = lottery.NewSample3(possibility100)

var x2Possibility = lottery.Possibility{
	"A2":   190,
	"A3":   185,
	"A4":   75,
	"A4.5": 40,
	"A5":   10,
	"B1.5": 68,
	"B2":   100,
	"B3":   32,
	"C1.5": 102,
	"C2":   150,
	"C3":   48,
}
var samplex2 = lottery.NewSample3(x2Possibility)

var score = 0

var jieguo map[string]int

func main() {
	jieguo = make(map[string]int)
	plans := data.GetPlan()
	for _, plan := range plans {
		if plan == -1 {
			teshu = false
			continue
		}
		gots := lot(plan)
		for _, item := range gots {
			jieguo[item]++
		}
	}
	fmt.Println(jieguo)
}

var teshu_sample *lottery.Sample3

var teshu = false
var baodi = true

func lot(times int) []string {
	gots := make([]string, 0, times)
	if baodi {
		if teshu {
			for i := 0; i < times; i++ {
				got := teshu_sample.Lot()
				gots = append(gots, got)
			}
		} else {
			for i := 0; i < times; i++ {
				got := sample.Lot()
				gots = append(gots, got)
			}
		}
		if times == 10 {
			if !util.StringContains(util.Keys(possibility10), gots) {
				got := sample10.Lot()
				gots[9] = got
			}
		} else if times == 100 {
			if !util.StringContains(util.Keys(possibility100), gots) {
				got := sample100.Lot()
				gots[99] = got
			}
		}
	} else {
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
		if teshu {
			for i := 0; i < times; i++ {
				got := teshu_sample.Lot()
				gots = append(gots, got)
			}
		} else {
			for i := 0; i < times; i++ {
				got := sample.Lot()
				gots = append(gots, got)
			}
		}
	}

	if score > 400 {
		score = score % 400
		temp := samplex2.Lot()
		item := temp[0:1]
		number := temp[1:]
		times, _ := strconv.ParseFloat(number, 64)
		new_possibility, _ := util.NewStringInt(possibility)
		// todo: gailv 需要改成正确的英文单词
		new_gailv := int(float64(new_possibility[item]) * times)
		new_possibility[item] = new_gailv
		teshu_sample = lottery.NewSample3(possibility)
		teshu = true
	}

	return gots
}
