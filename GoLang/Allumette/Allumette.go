package main

import "fmt"

func main() {
	var NbAllumette = 0
	var Choix1 int
	var Choix2 int

	fmt.Println("Veuillez choisir le nombre d'allumettes.")
	fmt.Scanln(&NbAllumette)

	for NbAllumette > 0 {
		fmt.Println("Il reste", NbAllumette, "allumettes.")
		fmt.Println("Joueur 1, choisissez le nombre d'allumettes que vous voulez enlenver")
		fmt.Scanln(&Choix1)

		if Choix1 > 3 || Choix1 < 1 || Choix1 > NbAllumette {
			fmt.Println("VOtre choix est invalide.")
			continue
		}
		NbAllumette -= Choix1

		if NbAllumette == 0 {
			fmt.Println("Joueur 1, vous avez perdu")
			break
		}

		fmt.Println("Il reste", NbAllumette, "allumettes.")
		fmt.Println("Joueur 2, choisissez le nombre d'allumettes que vous voulez enlenver")
		fmt.Scanln(&Choix2)

		if Choix2 > 3 || Choix2 < 1 || Choix2 > NbAllumette {
			fmt.Println("VOtre choix est invalide.")
			continue
		}
		NbAllumette -= Choix2

		if NbAllumette == 0 {
			fmt.Println("Joueur 2, vous avez perdu")
			break
		}

	}
}
