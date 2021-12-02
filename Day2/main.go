package main

import (
	"fmt"
  "os"
  "strings"
  "strconv"
)

func main() {
  parsedData := parseData()

  fmt.Println(calculateMovement(parsedData))

}

func parseData() []string{
  inputValues, _ := os.ReadFile("input.txt")

  separateValues := strings.Split(string([]byte(inputValues)),"\r\n")
  return separateValues
}


func calculateMovement(directions []string) int {
  depth := 0
  x := 0
  aim := 0
  for _, direction := range directions{
    parsedMovement:= strings.Fields(direction)
    change, _ := strconv.Atoi(parsedMovement[1])
    switch parsedMovement[0] {
    case "forward":
      x += change
      depth += (aim * change)
    case "down":
      aim += change
    case "up":
      aim -= change
    }
  }
  fmt.Println(depth)
  fmt.Println(x)
  return(depth * x)
}
