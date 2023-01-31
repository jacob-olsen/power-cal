package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	sizeArgs := len(os.Args)
	if sizeArgs == 1 {
		fmt.Println("no path given")
	} else {
		totalSum := 0.0
		for i := 1; i < sizeArgs; i++ {
			number := float64(readMoth(os.Args[i])) / 60000
			totalSum += number
			fmt.Println("power use for moth " + strconv.FormatFloat(number, 'f', 2, 64) + ":KWh")
		}
		fmt.Println()
		fmt.Println("tortal Power " + strconv.FormatFloat(totalSum, 'f', 2, 64) + ":KWh")
	}
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
