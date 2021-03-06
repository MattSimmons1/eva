package main

import (
  "fmt"
  "github.com/nsf/termbox-go"
  "os"
  "strconv"
  "time"
)

var ctrlXPressed = false

func eventLoop() {
  for !ctrlXPressed {
    switch ev := termbox.PollEvent(); ev.Type {
    case termbox.EventKey:
      if ev.Key == termbox.KeyCtrlX || ev.Key == termbox.KeyCtrlC {
        ctrlXPressed = true
      }
    case termbox.EventError:
      panic(ev.Err)
    }
  }
}

func main() {
  args := os.Args[1:]
  var fromSeconds int64 = 300

  if len(args) > 0 {
    if args[0] == "help" || args[0] == "-h" || args[0] == "--help" {
      fmt.Println("")
      fmt.Println("ππππ")
      fmt.Println("πππππππ")
      fmt.Println("π π π π π π π π π π")
      fmt.Println("PROGRAM:")
      fmt.Println("                        ππππΈπ πππππΏπ ππΌππΈπ")
      fmt.Println("")
      fmt.Println("SOURCE https://GITHUB.COM/MATTSIMMONS1/EVA")
      fmt.Println("TAKES THE NUMBER OF SECONDS AS AN ARGUMENT")
      fmt.Println("LOOKS BEST IN ITERM2: 20 ROWS x 80 COLUMNS")
      fmt.Println("")
      return
    } else if s, err := strconv.ParseInt(args[0], 10, 64); err == nil {
      fromSeconds = s
    }
  }

  err := termbox.Init()
  if err != nil {
    panic(err)
  }
  defer termbox.Close()

  termbox.SetOutputMode(termbox.Output256)

  termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
  termbox.Flush()

  drawTimer(fromSeconds)

  go eventLoop()

  tick := time.Tick(10 * time.Millisecond)

  for !ctrlXPressed {
    select {
    case <-tick:
      termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
      drawTimer(fromSeconds)
      termbox.Flush()
    }
  }
}
