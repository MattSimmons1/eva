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
      fmt.Println("𝐍𝐄𝐎𝐍")
      fmt.Println("𝐆𝐄𝐍𝐄𝐒𝐈𝐒")
      fmt.Println("𝐄 𝐕 𝐀 𝐍 𝐆 𝐄 𝐋 𝐈 𝐎 𝐍")
      fmt.Println("PROGRAM:")
      fmt.Println("                        𝑃𝑂𝑊𝐸𝑅 𝑆𝑈𝑃𝑃𝐿𝑌 𝑇𝐼𝑀𝐸𝑅")
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
