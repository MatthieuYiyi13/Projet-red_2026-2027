package main

import (
	"os/exec"
)

func main() {
	err := exec.Command("cmd", "/c", "start", "", "maptest.png").Run()
	if err != nil {
		panic(err)
	}
}
