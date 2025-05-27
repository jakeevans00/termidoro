package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const width = 100
const (
	BgColor    = "\x1b[41m" 
	ResetColor = "\x1b[0m"
)
const bar = "█"
const emptyBar = "░"
const pomoIcon = "🍊"

func main() {
	defer showCursor()
	hideCursor()

	fmt.Println(pomoIcon + " Termidoro")
	fmt.Println()
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [work|break|help]")
		return
	}
	command := os.Args[1]
	timeInSeconds := 25 * 60 // 25 Minutes
	if len(os.Args) > 2 {
		if customTime, err := strconv.Atoi(os.Args[2]); err == nil {
			timeInSeconds = customTime
		}
	}

	runProgram(command, timeInSeconds)
}

func runProgram(command string, timeInSeconds int) {
	switch command {
	case "work":
		focus(timeInSeconds)
	case "focus":
		focus(timeInSeconds)
	case "break":
		breakSession()
	case "help":
		printUsage()
	default:
		printUsage()
	}
}

func focus(timeInSeconds int) {
	for i := range timeInSeconds {
		printProgressBar(i, timeInSeconds)
		time.Sleep(1 * time.Second)
	}

	printProgressBar(timeInSeconds, timeInSeconds)
	time.Sleep(1 * time.Second)

	fmt.Printf("\n\n%s Focus session complete!\n", pomoIcon)
	sendNotification("Termidoro", pomoIcon + " Focus session complete!")
}

func breakSession() {
	printBreatCat()
	for i := range 300 {
		printProgressBar(i, 300)
		time.Sleep(1 * time.Second)
	}
}

func printUsage() {
	fmt.Println("Commands:")
	fmt.Println("  focus | work [time] - Start a focus session")
	fmt.Println("  break        - Start a break")
	fmt.Println("  help         - Show this help")
} 

func printBreatCat() {
	fmt.Println("  /\\_/\\   Zzz...")
	fmt.Println(" ( -.- )")
	fmt.Println("  ( U )")
	fmt.Println()
}

func printProgressBar(timeElapsed int, timeTotal int) {
	progress := float64(timeElapsed) / float64(timeTotal)
	barLength := int(progress * float64(width))

	filled := BgColor + strings.Repeat(bar, barLength) + ResetColor
	empty := strings.Repeat(emptyBar, width-barLength)

	elapsedFormatted := fmt.Sprintf("%02d:%02d", timeElapsed/60, timeElapsed%60)
	totalFormatted := fmt.Sprintf("%02d:%02d", timeTotal/60, timeTotal%60)

	fmt.Printf("\r[%s%s] %3d%% %s / %s%s", filled, empty, int(progress*100), elapsedFormatted, totalFormatted, strings.Repeat(" ", 10))
}

func hideCursor() {
	fmt.Print("\033[?25l")
}

func showCursor() {
	fmt.Print("\033[?25h")
}

func sendNotification(title, message string) {
	cmd := exec.Command("terminal-notifier", "-title", title, "-message", message)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error sending notification:", err)
	}
}