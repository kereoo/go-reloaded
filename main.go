package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	entree := "texte.txt"
	sortie := "resultat.txt"
	
	if len(os.Args) >= 2 {
		entree = os.Args[1]
	}
	if len(os.Args) >= 3 {
		sortie = os.Args[2]
	}
	contenu, err := os.ReadFile(entree)
	if err != nil {
		fmt.Println("Erreur de lecture :", err)
		os.Exit(1)
	}

	lignes := strings.Split(string(contenu), "\n")
	for i, ligne := range lignes {
		fin := ""
		if strings.HasSuffix(ligne, "\r") {
			ligne = strings.TrimSuffix(ligne, "\r")
			fin = "\r"
		}
		lignes[i] = Guillemets(ponctu(detect(ligne))) + fin
	}
	err = os.WriteFile(sortie, []byte(strings.Join(lignes, "\n")), 0644)
	if err != nil {
		fmt.Println("Erreur d'écriture :", err)
		os.Exit(1)
	}
	fmt.Println("Fichier modifié :", sortie)
}
