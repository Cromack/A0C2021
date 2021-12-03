package main

import (
	"fmt"
  "os"
  "strings"
  "strconv"
)

func main() {
  parsedData := parseData()
  gammaBit, deltaBit := mostCommonBit(parsedData)
  gamma, _ :=  strconv.ParseInt(gammaBit, 2, 64)
  delta, _ := strconv.ParseInt(deltaBit, 2, 64)
  fmt.Println("Decimal gamma rate", gamma)
  fmt.Println("Decimal delta rate", delta)
  fmt.Println("Power consumption", delta*gamma)

  oxygenBit := findOxygenRating(parsedData, 0)
  oxygen, _ := strconv.ParseInt(oxygenBit, 2, 64)
  fmt.Println("Oxygen", oxygen)

  C02Bit := findC02Rating(parsedData, 0)
  C02, _ := strconv.ParseInt(C02Bit, 2, 64)
  fmt.Println("C02", C02)

  fmt.Println("Life Support", C02*oxygen)

}

func parseData() []string{
  inputValues, _ := os.ReadFile("input.txt")
  separateValues := strings.Split(string([]byte(inputValues)),"\r\n")
  return separateValues
}

func mostCommonBit(diagnostics []string) (string, string){
  length := len(diagnostics[0])
  mostCommon := ""
  leastCommon := ""
  for i := 0; i < length; i++{
    ones := 0
    zeros := 0
    for _, row := range diagnostics{
      if string(row[i]) == "0"{
        zeros ++
      }else if string(row[i]) == "1"{
        ones ++
      }
    }
    if zeros < ones{
      mostCommon += "1"
      leastCommon += "0"
    }else{
      mostCommon += "0"
      leastCommon += "1"
    }
  }

  return mostCommon, leastCommon
}

func findOxygenRating(diagnostics []string, i int) string{
  var lastValue = diagnostics[0]
  if len(diagnostics) > 1{
    var oneSlice []string
    var zeroSlice []string
      ones := 0
      zeros := 0
      for _, row := range diagnostics{
        if string(row[i]) == "0"{
          zeros ++
          zeroSlice = append(zeroSlice, row)
        }else if string(row[i]) == "1"{
          ones ++
          oneSlice = append(oneSlice, row)
        }

      }

      if zeros < ones{
        lastValue = findOxygenRating(oneSlice, i+1)
      }else if zeros > ones{
        lastValue = findOxygenRating(zeroSlice, i+1)
      }else if zeros == ones{
        lastValue = findOxygenRating(oneSlice, i+1)
      }

  }
  return lastValue
}

func findC02Rating(diagnostics []string, i int) string{
  //length := len(diagnostics[0])
  var lastValue = diagnostics[0]
  if len(diagnostics) > 1{
    var oneSlice []string
    var zeroSlice []string
      ones := 0
      zeros := 0
      for _, row := range diagnostics{
        if string(row[i]) == "0"{
          zeros ++
          zeroSlice = append(zeroSlice, row)
        }else if string(row[i]) == "1"{
          ones ++
          oneSlice = append(oneSlice, row)
        }

      }

      if zeros > ones{
        lastValue = findC02Rating(oneSlice, i+1)
      }else if zeros < ones{
        lastValue = findC02Rating(zeroSlice, i+1)
      }else if zeros == ones{
        lastValue = findC02Rating(zeroSlice, i+1)
      }

  }
  return lastValue
}
