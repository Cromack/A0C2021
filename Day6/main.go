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
type point struct{
  x,y int
}

func main() {
  parsedData := parseData()
	fishStates := convertData(parsedData)
  for days := 0; days < amountOfDays; days ++{
    increaseCycle(&fishStates)
		//fmt.Println((fishStates))

  }
	fmt.Println(len(fishStates))
}

func parseData() []string{
  inputValues, _ := os.ReadFile("testInput.txt")
  separateValues := strings.Split(string([]byte(inputValues)),",")
  return separateValues
}

func convertData(rawData []string) []int{
	converted := make([]int, len(rawData))
	for i, value := range rawData{
		x,_ := strconv.Atoi(value)
		converted[i] = x
	}
	return converted
}

func increaseCycle(lanternfish *[]int){
	for i, fish := range *lanternfish{
		if fish == 0{
			(*lanternfish)[i] = existingCycle
			*lanternfish = append(*lanternfish, newCycle)
		}else{
			(*lanternfish)[i]--
		}
	}

}
