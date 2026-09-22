package projet 

import (
	"fmt"
)

func (p *Character) Combat_start(){
	monstre := "Loup"
	fmt.Println()
	fmt.Printf("Vous rencontrez un %s sauvage !",monstre)
	fmt.Println()
	fmt.Println("Que voulez vous faire ?")
		fmt.Printf("\n")
	fmt.Printf("\t+------------------------------------------------+\n")
	fmt.Printf("\t|                                                |\n")
	fmt.Printf("\t|            Que voulez vous faire ?             |\n")
	fmt.Printf("\t|                                                |\n")
	fmt.Printf("\t|     [O] Objet       [F] Fuir       [A] Attaque |\n")
	fmt.Printf("\t|                                                |\n")
	fmt.Printf("\t+------------------------------------------------+\n")
	fmt.Println()
	var moove string 
	fmt.Scanln(&moove)
		if choice == "O" || "o" {
			
		}
		if choice == "2" {
}