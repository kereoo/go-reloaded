package main 

import "strings"

func lireEntier(cible string) (int, bool) {
	if cible == "" {
		return 0, false
	}
	nombre := 0
	for _, nbr := range cible {
		if nbr < '0' || nbr > '9' {
			return 0, false
		}
		nombre = nombre*10 + int(nbr-'0')
	}
	return nombre, true
}

func lireCommande(mot string) (string, int, bool) {
	if !strings.HasPrefix(mot, "(") || !strings.HasSuffix(mot, ")") {
		return "", 0, false
	}
	interieur := mot[1 : len(mot)-1]
	parties := strings.Split(interieur, ",")
	nom := strings.TrimSpace(parties[0])
	if nom != "up" && nom != "low" && nom != "cap" {
		return "", 0, false
	}
	n := 1
	if len(parties) == 2 {
		v, ok := lireEntier(strings.TrimSpace(parties[1]))
		if !ok || v < 1 {
			return "", 0, false
		}
		n = v
	} else if len(parties) > 2 {
		return "", 0, false
	}
	return nom, n, true
}