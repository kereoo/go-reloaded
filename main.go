package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Veuillez utiliser comme tel : go run . <fichier_a_traiter.txt>")
		os.Exit(1)
	}
	chemin := os.Args[1]

	contenu, err := os.ReadFile(chemin)
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
		lignes[i] = detect(ligne) + fin
	}
	err = os.WriteFile(chemin, []byte(strings.Join(lignes, "\n")), 0644)
	if err != nil {
		fmt.Println("Erreur d'écriture :", err)
		os.Exit(1)
	}
	fmt.Println("Fichier modifié :", chemin)
}
