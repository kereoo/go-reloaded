package main

func estGui(runes []rune, compteur int) bool {
	if runes[compteur] != '\'' {
		return false
	}
	pre := compteur == 0 || runes[compteur-1] == ' '
	suf := compteur+1 == len(runes) || runes[compteur+1] == ' '
	return pre || suf
}

func Guillemets(phrase string) string {
	runes := []rune(phrase)

	total := 0
	for compteur := range runes {
		if estGui(runes, compteur) {
			total++
		}
	}
	limite := total - total%2

	resultat := []rune{}
	vus := 0
	ouvert := false
	for i := 0; i < len(runes); i++ {
		if vus >= limite || !estGui(runes, i) {
			resultat = append(resultat, runes[i])
			continue
		}
		vus++
		if !ouvert {
			resultat = append(resultat, '\'')
			for i+1 < len(runes) && runes[i+1] == ' ' {
				i++
			}
			ouvert = true
		} else {
			for len(resultat) > 0 && resultat[len(resultat)-1] == ' ' {
				resultat = resultat[:len(resultat)-1]
			}
			resultat = append(resultat, '\'')
			ouvert = false
		}
	}
	return string(resultat)
}
