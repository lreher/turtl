package main

import (
  "fmt"
  "os"
  "os/user"
  "turtl/repl"
)

func main() {
  user, err := user.Current()

  if err != nil {
    panic(err)
  }

  fmt.Printf("Hello %s! This is Turtl lang.\n", user.Username)

  repl.Start(os.Stdin, os.Stdout)
}
