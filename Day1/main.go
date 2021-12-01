package main

import (
	"fmt"
  "os"
  "strings"
  "strconv"
)

func main() {
  parsedData := parseData()
  fmt.Println(caclulateDepths(parsedData))

}

func parseData() []int{
  inputValues, _ := os.ReadFile("input.txt")
  numberData := []int{}

  separateValues := strings.Split(string([]byte(inputValues)),"\r\n")
  for _, i := range separateValues {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		numberData = append(numberData, j)
	}
  return numberData
}

func caclulateDepths(depthList []int) int{
  increases := 0
  previousDepth := depthList[0]
  for i, _ := range depthList[0:len(depthList)-2]{
    windowDepth := calculateSlidingWindow(depthList[i:i+3])
      fmt.Println(windowDepth)
    if windowDepth > previousDepth{
      increases ++
    }
    previousDepth = windowDepth
  }
  return increases
}

func calculateSlidingWindow( windowValues []int) int{
  windowDepth := 0
  for _, depth := range windowValues{
      windowDepth += depth
  }
  return windowDepth
}
