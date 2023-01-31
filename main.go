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
	dayList := make([]int, len(folder))
	dayName := make([]string, len(folder))

	for i, days := range folder {
		dayList[i] = readDay(taget + "/" + days.Name())
		dayName[i] = days.Name()

		sum += dayList[i]
	}
	gemenSnit := (100 / float64(sum/len(dayList)))
	for i := 0; i < len(dayList); i++ {
		fmt.Println(dayName[i] + " - " + strconv.FormatInt(int64(dayList[i]), 10) + " ~ " + strconv.FormatFloat(float64(dayList[i])*gemenSnit, 'f', 2, 64) + "%")
	}

	return sum
}
