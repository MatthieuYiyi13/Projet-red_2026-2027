package projet

import (
	"fmt"
	"os"
	"os/exec"
)

func (p *Character)AfficheMap() {
	dossier, _ := os.Getwd()
	fmt.Println("Je suis dans :", dossier)

	err := exec.Command("cmd", "/c", "start", "", "docs/mapV2.png").Run()
	if err != nil {
		panic(err)
	}
}
