package main

import (
	"fmt"
  "os"
  "strings"
  "strconv"
)

var amountOfDays = 256
var newCycle = 8
var existingCycle = 6

func main() {
	fishStates := convertData(parseData())
  for days := 0; days < amountOfDays; days ++{
    fishStates = increaseCycle(fishStates)

  }
	fishSum := 0
	for _, v := range fishStates{
		fishSum = fishSum + v
	}
	fmt.Println(fishSum)
}

func parseData() []string{
  inputValues, _ := os.ReadFile("input.txt")
  separateValues := strings.Split(string([]byte(inputValues)),",")
  return separateValues
}

func convertData(rawData []string) map[int]int{
	converted := make(map[int]int)
	for _, value := range rawData{
		x,_ := strconv.Atoi(value)
		converted[x]++
	}
	return converted
}

func increaseCycle(lanternfish map[int]int)map[int]int {
	tempMap := make(map[int]int)
	for i, v := range lanternfish{
		if i > 0{
			tempMap[i-1] = tempMap[i-1] + v
		}else if i == 0{
			tempMap[existingCycle] = tempMap[existingCycle] + v
			tempMap[newCycle] = v
		}
	}
	return tempMap
}
