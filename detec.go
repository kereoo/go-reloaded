package main

func detect(phrase string) string {
	liste_mots := []string{}
	conca := ""
	for _, c := range phrase {
		if c == ' ' {
			if conca != "" {
				liste_mots = append(liste_mots, conca)
				conca = ""
			}
		} else {
			conca += string(c)
		}
	}
	if conca != "" {
		liste_mots = append(liste_mots, conca)
	}
	resultat := []string{}
	for _, mot := range liste_mots {
		if mot == "(hex)" && len(resultat) > 0 {
			if texte, ok := hexatodecimal(resultat[len(resultat)-1]); ok {
				resultat[len(resultat)-1] = texte
				continue
			}
		}
		resultat = append(resultat, mot)
	}
	output := ""
	for i, w := range resultat {
		if i > 0 {
			output += " "
		}
		output += w
	}
	return output
}
