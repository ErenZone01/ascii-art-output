package main

import (
	"fmt"
	"os"
	"strings"
)

func foundLine(texte string) []string {
	var compteur int = 0
	var phrase string
	var tabPhrase []string

	for i := 0; i < len(texte); i++ {
		if texte[i] != '\n' {
			phrase += string(texte[i])
		} else {
			compteur++
			if compteur == 9 {
				tabPhrase = append(tabPhrase, phrase)
				phrase = ""
				compteur = 0
				continue
			}
			if compteur != 9 {
				phrase += "\n"
			}

		}

	}
	return tabPhrase
}

func foundposition(position []int, tabPhrase []string) []string {
	var textePhrase []string
	for i := 0; i < len(position); i++ {
		for j := 0; j < len(tabPhrase); j++ {
			if position[i] == j {
				textePhrase = append(textePhrase, tabPhrase[j])
			}
		}
	}
	return textePhrase
}

func compraison(tabASCII []rune, arg1 string, position []int, tab []int) []int {
	var actif bool = false
	var special bool = true

	for i := 0; i < len(arg1); i++ {
		for j := 0; j < len(tabASCII); j++ {
			if i < len(arg1)-1 && (string(arg1[i]) == "\\" && string(arg1[i+1]) == "n") {
				actif = true
				break
			} else {
				if actif == false {
					if rune(arg1[i]) == tabASCII[j] {
						position = append(position, tab[j])
						special = true
						break
					} else {
						special = false
					}
				} else {
					actif = false
					break
				}

			}

		}
		if !special {
			error()
		}
	}
	return position
}

func HorzontalLine(x [9][]string) string {
	var texte2 string
	for i, ligne := range x {
		if i > 0 {
			for _, part := range ligne {
				if len(part) > 0 {
					part = part[:len(part)-1]
				}
				texte2 += part
				texte2 += " "

			}
			if len(ligne[0]) == 1 && ligne[0] != "\n" {
				continue
			} else {
				texte2 += "\n"
			}

		}
	}
	return texte2

}

func SplitTexte(textePhrase []string) [9][]string {
	x := [9][]string{}
	for _, text := range textePhrase {
		for i, y := range strings.Split(text, "\n") {
			x[i] = append(x[i], y)
		}
	}
	return x
}

func ajoutASCII() []rune {
	var tabASCII []rune
	for i := ' '; i <= '~'; i++ {
		tabASCII = append(tabASCII, i)
	}
	return tabASCII
}

func ajoutNmbr(arg1 string) []int {
	var tab []int
	for i := 0; i < 95; i++ {
		tab = append(tab, i)
	}
	return tab
}

func CaseResolved(arg1 string, arg3 string) bool {
	var actif bool = true
	var text string
	if len(os.Args) >= 2 && arg1 == "\\n" {
		if len(os.Args) == 4  || (IsFile(os.Args[1]) == true) {
			var file *os.File = writeFile(arg3)
			if file == nil {
				os.Exit(1)
			}
			file.WriteString("\n")
			os.Exit(0)
		}
		fmt.Println()
		return true
	}
	for i := 0; i < len(arg1); i = i + 2 {
		if (i+1 != len(arg1)) && (arg1[i] == '\\' && arg1[i+1] == 'n') {
			actif = false
		} else {
			actif = true
			break
		}

	}
	if actif == false {
		for i := 0; i < len(arg1); i = i + 2 {
			text += "\n"
		}
		if len(os.Args) == 4 || (IsFile(os.Args[1]) == true)  {
			var file *os.File = writeFile(arg3)
			if file == nil {
				os.Exit(0)
			}
			file.WriteString(text)
			os.Exit(0)
		}
		for i := 0; i < len(arg1); i = i + 2 {
			fmt.Println()
		}
		actif = true
		return true
	}

	if len(os.Args) >= 2 && arg1 == "" {
		if len(os.Args) == 4 || (IsFile(os.Args[1]) == true)  {
			var file *os.File = writeFile(arg3)
			if file == nil {
				os.Exit(1)
			}
			file.WriteString("")
		}
		return true
	}
	return false
}

func DefineArg() string {
	var arg2 string
	if len(os.Args) == 2 {
		arg2 = "standard"
	} else if len(os.Args) == 3 {
		if (os.Args[len(os.Args)-1] != "standard" && os.Args[len(os.Args)-1] != "shadow" && os.Args[len(os.Args)-1] != "thinkertoy") && (IsFile(os.Args[1]) == true) {
			arg2 = "standard"
			return arg2
		} else {
			arg2 = os.Args[2]
		}
		return arg2
	} else if len(os.Args) == 4 {
		arg2 = os.Args[3]
		return arg2
	}
	return arg2
}

func foundFile(arg2 string) string {
	file, err := os.ReadFile(arg2 + ".txt")
	if err != nil {
		error()
	}

	return string(file)
}
func error() {
	fmt.Println("Usage: go run . [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("EX: go run . something standard")
	os.Exit(0)
}

func IsFile(arg3 string) bool {
	var option = "--output="
	var fichier string
	if len(arg3) > 10 {
		if strings.Contains(option, arg3[0:9]) {
			for i := 9; i < len(arg3); i++ {
				fichier += string(arg3[i])
			}
			if len(fichier) > 4 && fichier[len(fichier)-1] == 't' && fichier[len(fichier)-2] == 'x' && fichier[len(fichier)-3] == 't' && fichier[len(fichier)-4] == '.' {
				return true
			} else {
				return false
			}

		} else {
			return false
		}
	} else {
		return false
	}
}

func writeFile(arg3 string) *os.File {
	var option = "--output="
	var fichier string
	var file *os.File
	if len(arg3) > 10 {
		if strings.Contains(option, arg3[0:9]) {
			for i := 9; i < len(arg3); i++ {
				fichier += string(arg3[i])
			}
			if len(fichier) > 4 && fichier[len(fichier)-1] == 't' && fichier[len(fichier)-2] == 'x' && fichier[len(fichier)-3] == 't' && fichier[len(fichier)-4] == '.' {
				file, _ = os.Create(fichier)
			} else {
				error()
			}

		} else {
			error()
		}
	} else {
		error()
	}

	return file
}

func asciiArt(arg1 string, arg3 string, texte string) {
	var arguments = strings.Split(arg1, "\\n")
	var texte2 string
	for i := 0; i < len(arguments); i++ {
		if len(arguments[i]) == 0 {
			texte2 += "\n"
			continue
		}
		var tab []int = ajoutNmbr(arg1)
		var tabPhrase []string
		var tabASCII []rune = ajoutASCII()
		var position []int
		var textePhrase []string

		x := [9][]string{}

		if arguments[i] != "\n" {
			arg1 = arguments[i]
			tab = ajoutNmbr(arg1)

			position = compraison(tabASCII, arg1, position, tab)
			tabPhrase = foundLine(texte)
			textePhrase = foundposition(position, tabPhrase)
			x = SplitTexte(textePhrase)
			texte2 += HorzontalLine(x)
		}

	}
	if len(os.Args) == 4 || len(os.Args) == 3 {
		if IsFile(arg3) {
			var file *os.File = writeFile(arg3)
			file.WriteString(texte2)
			os.Exit(0)
		} else {
			if len(os.Args) == 4 {
				var file *os.File = writeFile(arg3)
				file.WriteString(texte2)
				os.Exit(0)
			}

		}

	}
	fmt.Print(texte2)
}
