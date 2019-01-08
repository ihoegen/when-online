package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [command] [args]\n", os.Args[0])
		fmt.Printf("Example: %s git push\n", os.Args[0])
		os.Exit(2)
	}
	for !connected() {
		time.Sleep(2 * time.Second)
	}
	run()
	os.Exit(0)
}

func connected() bool {
	_, err := http.Get("http://clients3.google.com/generate_204")
	if err != nil {
		return false
	}
	return true
}

func run() {
	currentDir, _ := os.Getwd()
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Dir = currentDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()
}
