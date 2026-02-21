package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// ASCII Digits V2 START
var digits2 = map[rune]string{
	'0': `
█████
█   █
█   █
█   █
█████
`,
	'1': `
  █  
  █  
  █  
  █  
  █  
`,
	'2': `
█████
    █
█████
█    
█████
`,
	'3': `
█████
    █
█████
    █
█████
`,
	'4': `
█   █
█   █
█████
    █
    █
`,
	'5': `
█████
█    
█████
    █
█████
`,
	'6': `
█████
█    
█████
█   █
█████
`,
	'7': `
█████
    █
    █
    █
    █
`,
	'8': `
█████
█   █
█████
█   █
█████
`,
	'9': `
█████
█   █
█████
    █
█████
`,
	':': `
     
  █  
     
  █  
     
`,
}

//END

func clear() {
	fmt.Print("\033[H\033[2J")
}

var linelength int

func printNumber2(s string, replacement string, r, g, b int) {
	lines := make([][]string, len(s))
	maxLines := 0

	for i, c := range s {
		digitStr := strings.ReplaceAll(digits2[c], "█", replacement)
		lines[i] = strings.Split(strings.Trim(digitStr, "\n"), "\n")

		if len(lines[i]) > maxLines {
			maxLines = len(lines[i])
		}
	}
	linelength = len(lines[1])
	for i := 0; i < maxLines; i++ {
		for j := 0; j < len(lines); j++ {
			if i < len(lines[j]) {
				fmt.Printf("\033[38;2;%d;%d;%dm%s   \033[0m", r, g, b, lines[j][i])
			} else {
				fmt.Printf("\033[38;2;%d;%d;%dm     \033[0m", r, g, b)
			}
		}
		fmt.Println()
	}
}

func main() {
	symbol := "█"
	t := 0.0
	for {
		clear()
		now := time.Now()
		Time := now.Format("15:04:05")
		day := now.Day()
		month := now.Month()
		year := now.Year()
		r := int((math.Sin(t*0.3) + 1) * 127)
		g := int((math.Sin(t*0.3+2) + 1) * 127)
		b := int((math.Sin(t*0.3+4) + 1) * 127)
		printNumber2(Time, symbol, r, g, b)
		dateStr := fmt.Sprintf("%02d.%02d.%d", day, month, year)
		fmt.Printf("%27s\033[38;2;%d;%d;%dm%s\033[0m\n", "", r, g, b, dateStr)
		t += 0.125
		time.Sleep(time.Millisecond * 50)
	}
}
