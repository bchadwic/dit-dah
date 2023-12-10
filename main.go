package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	cmd := exec.Command("stty", "-g", "-F", "/dev/tty")
	settings, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	defer resetTty(settings)
	cmd = exec.Command("stty", "-F", "/dev/tty", "cbreak", "-echo")
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
	input := make(chan rune, 20)
	go readRune(input)
	select {
	case i := <-input:
		fmt.Printf("you typed %s\n", string(i))
	case <-time.NewTimer(3 * time.Second).C:
		fmt.Println("time is up")
	}
}

func readRune(input chan rune) {
	reader := bufio.NewReader(os.Stdin)
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			panic(err)
		}
		input <- char
	}
}

func resetTty(settings []byte) {
	cmd := exec.Command("stty", "-F", "/dev/tty", strings.TrimRight(string(settings), "\n\r"))
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
