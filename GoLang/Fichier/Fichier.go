package main

import (
	"fmt"
	"os"
)

func UserEntry() string {
	var user string
	fmt.Scanln(&user)
	return user
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read(file string) string {
	data, err := os.ReadFile(file)
	check(err)
	return string(data)
	fmt.Println("Contenu du fichier :")
	fmt.Println(data)
}

func write(text string, file *os.File) {
	if _, err := file.WriteString(text); err != nil {
		panic(err)
	}
	fmt.Println("Entrez le texte à écrire dans le fichier :")
	text = UserEntry()
	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	check(err)
	defer file.Close()
	write(text, file)
}

func Delete(text string, file *os.File) {

}

func main() {
	var i string
	fmt.Scan(&i)
	switch i {
	case "1":
		{
			read()
		}
	case "2":
		{
			write()
		}
	case "3":
		{
			Delete()
		}
	case "4":
		{

		}
	default:
		{
			fmt.Println("Choisissez parmi les 4 propositions")
		}

	}
}
