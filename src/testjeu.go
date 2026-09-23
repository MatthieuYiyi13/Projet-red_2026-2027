package projet
var nom string
fmt.Scanln(&nom)

if len(nom) > 0 && nom[0] >= 'a' && nom[0] <= 'z' {
    nom = string(nom[0]-'a'+'A') + nom[1:]
}

p.name = nom