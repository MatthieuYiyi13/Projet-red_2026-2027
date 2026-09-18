package projet

import (
	"os/exec"
)

func (p *Character)AfficheMap() {
	err := exec.Command("cmd", "/c", "start", "", "docs/mapV2.png").Run()
	if err != nil {
		panic(err)
	}
}
