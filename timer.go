package main

import (
  "github.com/nsf/termbox-go"
  "math"
  "time"
)

func drawHorizontalLine(fromX int, fromY int, toX int, fg, bg termbox.Attribute) {
  for i := fromX; i <= toX; i++ {
    termbox.SetCell(i, fromY, borderVertical, fg, bg)
  }
}

const borderVertical rune = 0x2500
const BLOCK_STAIRS_1 rune = 0x259F // ▟
const BLOCK_STAIRS_2 rune = 0x2599 // ▙
const BLOCK_STAIRS_3 rune = 0x259B // ▛
const BLOCK_STAIRS_4 rune = 0x259C // ▜
const BLOCK rune = 0x2588          // █
const SQUARE rune = 0x25A0         // ■
const RIGHT_HALF_BLOCK = 0x2590    // ▐
const LEFT_HALF_BLOCK = 0x258C     // ▌
func drawLCDNumber(i int, x, y int, fg, bg termbox.Attribute) {

  // top
  if i == 0 || i == 2 || i == 3 || i == 5 || i == 6 || i == 7 || i == 8 || i == 9 {
    termbox.SetCell(x+1, y+0, BLOCK_STAIRS_4, fg, bg)
    termbox.SetCell(x+2, y+0, BLOCK, fg, bg)
    termbox.SetCell(x+3, y+0, BLOCK, fg, bg)
    termbox.SetCell(x+4, y+0, BLOCK, fg, bg)
    termbox.SetCell(x+5, y+0, BLOCK_STAIRS_3, fg, bg)
  }

  // top left
  if i == 0 || i == 4 || i == 5 || i == 6 || i == 8 || i == 9 {
    termbox.SetCell(x, y+1, BLOCK_STAIRS_2, fg, bg)
    termbox.SetCell(x, y+2, BLOCK, fg, bg)
    termbox.SetCell(x, y+3, BLOCK_STAIRS_3, fg, bg)
    termbox.SetCell(x-1, y+1, RIGHT_HALF_BLOCK, fg, bg)
    termbox.SetCell(x-1, y+2, RIGHT_HALF_BLOCK, fg, bg)
    termbox.SetCell(x-1, y+3, RIGHT_HALF_BLOCK, fg, bg)
  }

  // top right
  if i == 0 || i == 1 || i == 2 || i == 3 || i == 4 || i == 7 || i == 8 || i == 9 {
    termbox.SetCell(x+6, y+1, BLOCK_STAIRS_1, fg, bg)
    termbox.SetCell(x+6, y+2, BLOCK, fg, bg)
    termbox.SetCell(x+6, y+3, BLOCK_STAIRS_4, fg, bg)
    termbox.SetCell(x+7, y+1, LEFT_HALF_BLOCK, fg, bg)
    termbox.SetCell(x+7, y+2, LEFT_HALF_BLOCK, fg, bg)
    termbox.SetCell(x+7, y+3, LEFT_HALF_BLOCK, fg, bg)
  }

  // middle
  if i == 2 || i == 3 || i == 4 || i == 5 || i == 6 || i == 8 || i == 9 {
    termbox.SetCell(x+1, y+4, BLOCK_STAIRS_1, fg, bg)
    termbox.SetCell(x+2, y+4, BLOCK, fg, bg)
    termbox.SetCell(x+3, y+4, BLOCK, fg, bg)
    termbox.SetCell(x+4, y+4, BLOCK, fg, bg)
    termbox.SetCell(x+5, y+4, BLOCK_STAIRS_3, fg, bg)
  }

  // bottom left
  if i == 0 || i == 2 || i == 6 || i == 8 {
    termbox.SetCell(x, y+5, BLOCK_STAIRS_2, fg, bg)
    termbox.SetCell(x, y+6, BLOCK, fg, bg)
    termbox.SetCell(x, y+7, BLOCK_STAIRS_3, fg, bg)
    termbox.SetCell(x-1, y+5, RIGHT_HALF_BLOCK, fg, bg)
    termbox.SetCell(x-1, y+6, RIGHT_HALF_BLOCK, fg, bg)
    termbox.SetCell(x-1, y+7, RIGHT_HALF_BLOCK, fg, bg)
  }

  // bottom right
  if i == 0 || i == 1 || i == 3 || i == 4 || i == 5 || i == 6 || i == 7 || i == 8 || i == 9 {
    termbox.SetCell(x+6, y+5, BLOCK_STAIRS_1, fg, bg)
    termbox.SetCell(x+6, y+6, BLOCK, fg, bg)
    termbox.SetCell(x+6, y+7, BLOCK_STAIRS_4, fg, bg)
    termbox.SetCell(x+7, y+5, LEFT_HALF_BLOCK, fg, bg)
    termbox.SetCell(x+7, y+6, LEFT_HALF_BLOCK, fg, bg)
    termbox.SetCell(x+7, y+7, LEFT_HALF_BLOCK, fg, bg)
  }

  // bottom
  if i == 0 || i == 2 || i == 3 || i == 5 || i == 6 || i == 8 || i == 9 {
    termbox.SetCell(x+1, y+8, BLOCK_STAIRS_1, fg, bg)
    termbox.SetCell(x+2, y+8, BLOCK, fg, bg)
    termbox.SetCell(x+3, y+8, BLOCK, fg, bg)
    termbox.SetCell(x+4, y+8, BLOCK, fg, bg)
    termbox.SetCell(x+5, y+8, BLOCK_STAIRS_2, fg, bg)
  }
}

func drawSmallLCDNumber(i int, x, y int, fg, bg termbox.Attribute) {

  // top
  if i == 0 || i == 2 || i == 3 || i == 5 || i == 6 || i == 7 || i == 8 || i == 9 {
    termbox.SetCell(x, y+0, BLOCK_STAIRS_1, fg, bg)
    termbox.SetCell(x+1, y+0, BLOCK, fg, bg)
    termbox.SetCell(x+2, y+0, BLOCK, fg, bg)
    termbox.SetCell(x+3, y+0, BLOCK_STAIRS_2, fg, bg)
  }
  // top left
  if i == 0 || i == 4 || i == 5 || i == 6 || i == 8 || i == 9 {
    termbox.SetCell(x, y+0, BLOCK_STAIRS_1, fg, bg)
    termbox.SetCell(x, y+1, BLOCK, fg, bg)
    termbox.SetCell(x, y+2, BLOCK_STAIRS_3, fg, bg)
  }

  // top right
  if i == 0 || i == 1 || i == 2 || i == 3 || i == 4 || i == 7 || i == 8 || i == 9 {
    termbox.SetCell(x+3, y+0, BLOCK_STAIRS_2, fg, bg)
    termbox.SetCell(x+3, y+1, BLOCK, fg, bg)
    termbox.SetCell(x+3, y+2, BLOCK_STAIRS_3, fg, bg)
  }

  // middle
  if i == 2 || i == 3 || i == 4 || i == 5 || i == 6 || i == 8 || i == 9 {
    termbox.SetCell(x, y+2, BLOCK_STAIRS_1, fg, bg)
    termbox.SetCell(x+1, y+2, BLOCK, fg, bg)
    termbox.SetCell(x+2, y+2, BLOCK, fg, bg)
    termbox.SetCell(x+3, y+2, BLOCK_STAIRS_3, fg, bg)
  }

  // bottom left
  if i == 0 || i == 2 || i == 6 || i == 8 {
    termbox.SetCell(x, y+2, BLOCK_STAIRS_1, fg, bg)
    termbox.SetCell(x, y+3, BLOCK, fg, bg)
    termbox.SetCell(x, y+4, BLOCK_STAIRS_3, fg, bg)
  }

  // bottom right
  if i == 0 || i == 1 || i == 3 || i == 4 || i == 5 || i == 6 || i == 7 || i == 8 || i == 9 {
    termbox.SetCell(x+3, y+2, BLOCK_STAIRS_3, fg, bg)
    termbox.SetCell(x+3, y+3, BLOCK, fg, bg)
    termbox.SetCell(x+3, y+4, BLOCK_STAIRS_3, fg, bg)
  }

  // bottom
  if i == 0 || i == 2 || i == 3 || i == 5 || i == 6 || i == 8 || i == 9 {
    termbox.SetCell(x, y+4, BLOCK_STAIRS_4, fg, bg)
    termbox.SetCell(x+1, y+4, BLOCK, fg, bg)
    termbox.SetCell(x+2, y+4, BLOCK, fg, bg)
    termbox.SetCell(x+3, y+4, BLOCK_STAIRS_3, fg, bg)
  }
}

// draw a box of the provided character
func drawBox(x0, x1, y0, y1 int, character rune, fg, bg termbox.Attribute) {
  for x := 0; x < x1-x0; x++ {
    for y := 0; y < y1-y0; y++ {
      termbox.SetCell(x0+x, y0+y, character, fg, bg)
    }
  }
}

func writeWord(x, y int, word []rune, fg, bg termbox.Attribute) {
  for i, r := range word {
    termbox.SetCell(x+i, y, r, fg, bg)

  }
}

/**
  ▜███▛
▙       ▟
█       █  ▀
▛       ▜
  ▟███▛
▙       ▟
█       █  ▀
▛       ▜
  ▟███▙
*/
var tTime = time.Now() // value to count down from
func drawTimer(fromSeconds int64) {

  fg := termbox.ColorWhite
  bg := termbox.ColorDefault

  milliseconds := fromSeconds*1000 - time.Since(tTime).Milliseconds()
  if milliseconds < 0 {
    milliseconds = 0
    bg = termbox.ColorRed

    drawBox(0, 80, 0, 20, ' ', fg, bg)
  }

  seconds := math.Floor(float64(milliseconds) / 1000) // 11:11.000
  milliseconds = milliseconds - int64(seconds)*1000   // 00:00.111
  minutes := math.Floor(seconds / 60)                 // 11:00.000
  seconds = seconds - 60*minutes                      // 00:11.000
  decaminutes := math.Floor(minutes / 10)             // 10:00.000
  minutes = minutes - decaminutes*10                  // 01:00.000

  decaseconds := math.Floor(seconds / 10) // 00:10.000
  seconds = seconds - decaseconds*10      // 00:01.000

  centiseconds := math.Floor(float64(milliseconds) / 100)      // 00:00.100
  milliseconds = milliseconds - int64(centiseconds*100)        // 00:00.011
  milliseconds = int64(math.Floor(float64(milliseconds) / 10)) // 00:00.010

  drawHorizontalLine(3, 15, 59, fg, bg)

  writeWord(4, 2, []rune("活動限界まで"), fg, bg)
  writeWord(4, 3, []rune("ACTIVE TIME REMAINING:"), fg, bg)
  drawLCDNumber(int(decaminutes), 5, 5, fg, bg)
  drawLCDNumber(int(minutes), 15, 5, fg, bg)

  termbox.SetCell(24, 7, SQUARE, fg, bg)
  termbox.SetCell(24, 12, SQUARE, fg, bg)
  termbox.SetCell(46, 10, SQUARE, fg, bg)
  termbox.SetCell(46, 12, SQUARE, fg, bg)

  drawLCDNumber(int(decaseconds), 27, 5, fg, bg)
  drawLCDNumber(int(seconds), 37, 5, fg, bg)

  writeWord(48, 2, []rune{'内'}, fg, bg)
  writeWord(50, 2, []rune{'部'}, fg, bg)
  writeWord(48, 3, []rune("INTERNAL"), fg, bg)

  writeWord(48, 6, []rune("主エネルギー供給システム"), fg, bg)
  writeWord(48, 7, []rune("MAIN ENERGY SUPPLY SYSTEM"), fg, bg)
  drawSmallLCDNumber(int(centiseconds), 48, 9, fg, bg)
  drawSmallLCDNumber(int(milliseconds), 54, 9, fg, bg)

  writeWord(5, 17, []rune("STOP"), fg, bg)
  writeWord(5, 18, []rune("やめる"), fg, bg)

  writeWord(17, 17, []rune("SLOW"), fg, bg)
  writeWord(17, 18, []rune("スロー"), fg, bg)

  writeWord(29, 17, []rune("NORMAL"), fg, bg)
  writeWord(29, 18, []rune{'正'}, fg, bg)
  writeWord(31, 18, []rune{'常'}, fg, bg)

  writeWord(42, 17, []rune("RACING"), fg, bg)
  writeWord(42, 18, []rune("レース"), fg, bg)
}
