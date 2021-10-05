package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	Sample("Command", Command)
	Sample("LookPath", LookPath)
}

func Command() {
	cmd := exec.Command("ls")
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
}

func LookPath() {
	result, err := exec.LookPath("ls")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Found:", result) // Found: /bin/ls
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}
