package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {

	for {
		// at the start of each loop get the hour, minute and second
		hour := time.Now().Hour()
		minute := time.Now().Minute()
		second := time.Now().Second()

		// display time then sleep for 1 second
		fmt.Printf("%02d:%02d:%02d\n", hour, minute, second)
		time.Sleep(1 * time.Second)

		// clear the console
		clearScreen()
	}
}

func clearScreen() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
