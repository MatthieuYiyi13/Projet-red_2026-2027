package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	dossier, _ := os.Getwd()
	fmt.Println("Je suis dans :", dossier)

	err := exec.Command("cmd", "/c", "start", "", "docs/maptest.png").Run()
	if err != nil {
		panic(err)
	}
}
