package main

import (
  "os"
  "bufio"
  "strings"
  "fmt"
)

var store = make(map[string]string)

func main() {

  scanner := bufio.NewScanner(os.Stdin)

  for {
    fmt.Print(">")
    scanner.Scan()
    input := scanner.Text()
    handleCommand(input)
  }

}

//input --> "SET name <name> / GET name"
func handleCommand(input string) {
  tokens := strings.Split(input," ")

  //fmt.Println(len(tokens))

  if len(tokens) < 2 {
    fmt.Println("Invalid command!")
    return
  }

  command := tokens[0]
  key := tokens[1]
  value := strings.Join(tokens[2:]," ")
  
  switch command {
  case "SET":
    store[key] = value
    fmt.Println("Added")
  case "GET":
    val, exists := store[key]
    if exists == true {
      fmt.Println(val)
    } else {
      fmt.Println("The key does not exist.")
    }
  default:
    fmt.Println("Command does not exist.")
  }
  
}
