package main

func estPonct(r rune) bool {
	return r == '.' || r == ',' || r == '!' || r == '?' || r == ':' || r == ';'
}

func Chiffre(nombre rune) bool {
	return nombre >= '0' && nombre <= '9'
}

func ponctu(phrase string) string {
	runes := []rune(phrase)
	resultat := []rune{}
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if !estPonct(r) {
			resultat = append(resultat, r)
			continue
		}
		if i > 0 && i+1 < len(runes) && Chiffre(runes[i-1]) && Chiffre(runes[i+1]) {
			resultat = append(resultat, r)
			continue
		}
		for len(resultat) > 0 && resultat[len(resultat)-1] == ' ' {
			resultat = resultat[:len(resultat)-1]
		}
		for i < len(runes) && estPonct(runes[i]) {
			resultat = append(resultat, runes[i])
			i++
		}
		for i < len(runes) && runes[i] == ' ' {
			i++
		}
		if i < len(runes) {
			resultat = append(resultat, ' ')
		}
		i--
	}
	return string(resultat)
}
