package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
    "regexp"
)


type weather struct {
  day int
  maxTemperature, minTemperature   int
}

func (w *weather) spread() int {
    return w.maxTemperature - w.minTemperature
}

func main() {
  weather := ReadWeather("weather.dat")
  var smallest = weather[0]
  for _,element := range weather  {
    if element.spread() < smallest.spread() {
      smallest = element
    }
  }
  fmt.Printf("%+v\n", smallest)

}

func ReadWeather(filename string) []weather {
  var w []weather
  file, err := os.Open("weather.dat")
  if err != nil {
     log.Fatal(err)
  }
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    var line = scanner.Text()
    var validLine = regexp.MustCompile("^\\s+(\\d+)\\s+(\\d+)\\s+(\\d+)")
    var matched = validLine.FindAllString(line, -1)
    if len(matched) > 0  {
      w = append(w, CreateWeather(matched[0]))
    }
  }
  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }
  return w
}

func CreateWeather(line string) weather {
  words := strings.Fields(line)
  return weather{
                  day : ParseInt(words[0]),
                  maxTemperature : ParseInt(words[1]),
                  minTemperature : ParseInt(words[2]),
                 }
}

func ParseInt(value string) int {
  strvalue, err := strconv.Atoi(value)
  if  err != nil {
    log.Println(err)
  }
  return strvalue
}
