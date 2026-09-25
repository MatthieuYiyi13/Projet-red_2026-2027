package projet

func (p *Character) AfficheCredit() {
	stopSound()
	PlaySoundAsynccredit()
	afficherTexte100("Inspiré par un royaume oublié, né d’un équilibre rompu entre les biomes et les destins.")
	afficherTexte50("Ce jeu a été créé dans l’univers d’Eldoria, un royaume façonné par des siècles de paix, de mystères et de conflits oubliés.")
	afficherTexte50("Chaque biome, de la Forêt d’Émeraude aux Montagnes de Glaces, a été pensé comme un fragment vivant de ce monde.")
	afficherTexte50("Que votre chemin à travers Eldoria soit rempli de choix, de dangers… et de gloire.")
	afficherTexte100("Projet RED fait en 35h")
	afficherTexte100("Langage utilisé : GOLANG")
	afficherTexte100("Réaliser en groupe de 3 par :")
	afficherTexte100("Davy FILOMENO")
	afficherTexte100("Matteo SOLAZ")
	afficherTexte100("Matthieu TOLISANO")
	afficherTexte10000("Merci à Ange et Line pour la résurrection")
}
