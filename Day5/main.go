package main

import (
	"fmt"
  "os"
  "strings"
  "strconv"
)

var lineMap = make(map[string]int)
type point struct{
  x,y int
}

func main() {
  parsedData := parseData()
  for _, input := range parsedData{
    mapLines(strings.Split(input," "))
  }
  fmt.Println(lineMap)
  getDangerPoints()
}

func parseData() []string{
  inputValues, _ := os.ReadFile("input.txt")
  separateValues := strings.Split(string([]byte(inputValues)),"\r\n")
  return separateValues
}

func mapLines(input []string){
  startingpoint := createPoint(input[0])
  endingpoint := createPoint(input[2])
  if (startingpoint.x == endingpoint.x) || (startingpoint.y == endingpoint.y){
    checkCoordinateOrder(&startingpoint, &endingpoint)
    //fmt.Println(startingpoint, endingpoint)
    for x := startingpoint.x; x <= endingpoint.x; x++{
      for y := startingpoint.y; y <= endingpoint.y; y ++{
        coordinatePoint := strconv.Itoa(x) + "," + strconv.Itoa(y)
        _, ok := lineMap[coordinatePoint]
        if ok{
          lineMap[coordinatePoint] = lineMap[coordinatePoint] + 1
        }else{
          lineMap[coordinatePoint] = 1
        }
      }
    }
  }else{
    //Toteutus diagonal lineille
    //Normicase menee alkuperäisellä functiolla ja coordinaattien järjestyksen vaihdolla, mutta on myös case jossa joko X tai y pienenee
  }
}

func createPoint(coordinates string) point{
    points := strings.Split(coordinates,",")
    x, _ := strconv.Atoi(points[0])
    y, _ := strconv.Atoi(points[1])

    return point{x:x, y:y}
}

func getDangerPoints(){
  total := 0
  for _, lineAmount := range lineMap{
    if lineAmount >= 2{
      total++
    }
  }
  fmt.Println(total)
}

func checkCoordinateOrder(start *point, end *point){
  if start.x > end.x{
    temp := *start
    *start = *end
    *end = temp
  }else if start.y > end.y{
    temp := *start
    *start = *end
    *end = temp
  }
  //fmt.Println(*start, *end)
}
