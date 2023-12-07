package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("echo", "hello")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}

