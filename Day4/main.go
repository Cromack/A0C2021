package main

import (
	"fmt"
  "os"
  "strings"
  "strconv"
  "regexp"
)


type bingoBoard struct {
  solved bool
  boardValues [5][]string
}

var boards []bingoBoard

func main() {
  parsedData := parseData()
  incomingValuestemp := parsedData[0]
  incomingValues := strings.Split(string(incomingValuestemp), ",")
  bingoBoards := parsedData[1:]
  createBingoBoards(bingoBoards)

  for _, value := range incomingValues{
    checkboards(string(value))
  }
}

func parseData() []string{
  inputValues, _ := os.ReadFile("input.txt")
  separateValues := strings.Split(string([]byte(inputValues)),"\r\n\r\n")
  return separateValues
}


func createBingoBoards(values []string){

  for _, inputBoard := range values{
    var tempboard [5][]string
    valuesParsed := strings.Split(inputBoard,"\r\n")

    for i, row := range valuesParsed{
      space := regexp.MustCompile(`\s+`)
      tempRow := space.ReplaceAllString(row, " ")
      rowCleaned := strings.TrimLeft(tempRow, " ")
      rowParsed := strings.Split(rowCleaned," ")
      tempboard[i] = (rowParsed)
    }

    boards = append(boards, bingoBoard{solved:false, boardValues:tempboard})

  }
}

func checkboards(value string){
  for index, board := range boards{
    if board.solved == false{
    for i, row := range board.boardValues{
      for j, boardValue := range row{
        if boardValue == value{
          board.boardValues[i][j] = "X"
          winValue, _ := strconv.Atoi(value)
          checkWin(&board, row, j, winValue)
          boards[index] = board
        }
        }
      }
    }
  }
}

func checkWin(boardToCheck *bingoBoard, row []string, index int, winningValue int){

    rowWin := 0
    columnWin := 0
    for _, value := range row{
      if value == "X"{
        rowWin ++
      }
    }
    if rowWin == 5{
      //fmt.Println("Voitto")
      //fmt.Println(boardToCheck)
      boardToCheck.solved = true
      calculateWin(boardToCheck.boardValues, winningValue)
    }
    for _, row := range boardToCheck.boardValues{
      if row[index] == "X"{
        columnWin ++
      }
    }
    if columnWin == 5{
      //fmt.Println("Voitto")
      //fmt.Println(boardToCheck)
      boardToCheck.solved = true
      calculateWin(boardToCheck.boardValues, winningValue)
    }
}

func calculateWin(boardToCheck [5][]string, winningValue int){
  winSum := 0
  for _, row := range boardToCheck{
    for _, boardValue := range row{
      if boardValue != "X"{
        number,_ :=  strconv.Atoi(boardValue)
        winSum += number
      }
    }
  }
  fmt.Println(winSum * winningValue)
}
