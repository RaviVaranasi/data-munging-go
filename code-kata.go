package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)


type weather struct {
  day int
  maxTemperature, minTemperature   int
}

func (w *weather) spread() int {
    return w.maxTemperature - w.minTemperature
}

func main() {
  file, err := os.Open("weather.dat")
  if err != nil {
     log.Fatal(err)
 }
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
      fmt.Println(scanner.Text())
  }

  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }
}

func 
