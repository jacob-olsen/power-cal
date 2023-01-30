package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("power use for moth " + strconv.FormatFloat(float64(readMoth("power/data/2022/12"))/60000, 'f', 2, 64))
}

func readDay(taget string) int {
	data, err := ioutil.ReadFile(taget)
	if err != nil {
		panic(err)
	}
	temp := string(data)[1:]
	temp = temp[:len(temp)-1]

	tempData := strings.Split(temp, "}{")

	lastNumber := 0
	sumNumber := 0

	for i, taget := range tempData {
		number, err := strconv.Atoi(taget[strings.LastIndex(taget, ":")+1:])
		if err != nil {
			panic(err)
		}

		if i == 0 {
			lastNumber = number
		}

		if lastNumber > number {
			sumNumber += number
			lastNumber = number
		} else {
			sumNumber += (number - lastNumber)
			lastNumber = number
		}
	}
	return sumNumber
}

func readMoth(taget string) int {
	folder, err := ioutil.ReadDir(taget)
	if err != nil {
		panic(err)
	}

	sum := 0

	for _, days := range folder {
		number := readDay(taget + "/" + days.Name())
		fmt.Println(days.Name() + " - " + strconv.FormatInt(int64(number), 10))

		sum += number
	}

	return sum
}
