package projet

import ( 
	"fmt"
	"os/exec"
	"strings"
)

var x_position int = 6
var y_position int = 4

func (p *Character)AfficheMap() {
	fmt.Println()
	fmt.Printf("Vous êtes actuellement en %d, %d\n", x_position, y_position)
	fmt.Println()
	err := exec.Command("cmd", "/c", "start", "", "docs/mapV2.png").Run()
	if err != nil {
		panic(err)
	}
}

func (p *Character) AfficheMapAscii() {
	fmt.Println()
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║                          ELDORIA                         ║")
	fmt.Println("╠══════════════════════════════════════════════════════════╣")
	fmt.Println("║ 🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵║🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋║")
	fmt.Println("║ 🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵║🌋🌋🌋🌋🌋🌋🌋🌋🌋●Astapor🌋║")
	fmt.Println("║ 🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵║🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋║")
	fmt.Println("║ 🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵║🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋║")
	fmt.Println("║ 🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵║🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋║")
	fmt.Println("║ 🌵🌵🌵🌵🌵🌵 ● Port-Réal🌵🌵║🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋║")
	fmt.Println("║ 🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵║🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋║")
	fmt.Println("║ 🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵║🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋║")
	fmt.Println("║ 🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵║🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋║")
	fmt.Println("║ 🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵║🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋║")
	fmt.Println("║ 🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵🌵║🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋🌋║")
	fmt.Println("╠─────────────────────────────╬────────────────────────────╣")
	fmt.Println("║ 🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲║🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊║")
	fmt.Println("║ 🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲║🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊║")
	fmt.Println("║ 🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲║🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊║")
	fmt.Println("║ 🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲║🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊║")
	fmt.Println("║ 🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲║🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊║")
	fmt.Println("║ 🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲║🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊║")
	fmt.Println("║ 🌲🌲🌲🌲🌲🌲 ● QARTH🌲🌲🌲🌲║🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊║")
	fmt.Println("║ 🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲║🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊║")
	fmt.Println("║ 🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲║🧊🧊🧊🧊🧊🧊 ●WINTERFELL🧊🧊║")
	fmt.Println("║ 🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲║🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊║")
	fmt.Println("║ 🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲║🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊🧊║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
}

const (
	mapWidth  = 28
	mapHeight = 22
)

func (p *Character) AfficheMapCo() {

	if x_position > 22 || y_position >28 || x_position<0 || y_position<0 {
		x_position = 0
		y_position = 0
		fmt.Println()
		fmt.Println("Vous ne pouvez pas sortir de la map !")
		fmt.Println("Vous êtes renvoyé au point de départ.")
		fmt.Println()
	}

	grille := make([][]string, mapHeight)
	for construire := 0; construire < mapHeight; construire++ {
		grille[construire] = make([]string, mapWidth)
		for col := 0; col < mapWidth; col++ {
			switch {
			case construire < mapHeight/2 && col < mapWidth/2:
				grille[construire][col] = "🌵" 
			case construire < mapHeight/2:
				grille[construire][col] = "🌋" 
			case col < mapWidth/2:
				grille[construire][col] = "🌲" 
			default:
				grille[construire][col] = "🧊" 
			}
		}
	}
	type ville struct {
		row, col int
		name    string
		occupe   int 
	}
	villes := []ville{
		{row: 1, col: 20, name: "● Astapor ", occupe: 5},
		{row: 5, col: 6, name: "● Port-Réal ", occupe: 6},
		{row: 16, col: 6, name: "●  QARTH", occupe: 4},
		{row: 18, col: 20, name: "● WINTERFELL", occupe: 6},
	}
	for _, v := range villes {
		grille[v.row][v.col] = v.name
		for i := 1; i < v.occupe && v.col+i < mapWidth; i++ {
			grille[v.row][v.col+i] = ""
		}
	}
	playerRow := (mapHeight - 1) - y_position
	playerCol := x_position
	if playerRow >= 0 && playerRow < mapHeight && playerCol >= 0 && playerCol < mapWidth {
		grille[playerRow][playerCol] = "📍"
	} 

	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║                          ELDORIA                         ║")
	fmt.Println("╠══════════════════════════════════════════════════════════╣")
 
	for row := 0; row < mapHeight; row++ {
		var sb strings.Builder
		sb.WriteString("║ ")
		for col := 0; col < mapWidth; col++ {
			sb.WriteString(grille[row][col])
			if col == mapWidth/2-1 {
				sb.WriteString("║")
			}
		}
		sb.WriteString("║")
		fmt.Println(sb.String())
 
		if row == mapHeight/2-1 {
			fmt.Println("╠─────────────────────────────╬────────────────────────────╣")
		}
	}
 
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Printf("\n📍 Position actuelle : %d, %d\n", x_position,y_position)
}

