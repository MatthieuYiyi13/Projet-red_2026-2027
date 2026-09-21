package projet

import ( 
	"fmt"
	"os/exec"
)

var x_position int = 0
var y_position int = 0

func (p *Character)AfficheMap() {
	fmt.Println()
	fmt.Printf("Vous êtes actuellement en %d, %d\n", x_position, y_position)
	fmt.Println()
	err := exec.Command("cmd", "/c", "start", "", "docs/mapV2.png").Run()
	if err != nil {
		panic(err)
	}
}
