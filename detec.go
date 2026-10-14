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
		if len(resultat) > 0 {
			dernier := resultat[len(resultat)-1]
			if mot == "(hex)" {
				if texte, ok := hexatodecimal(dernier); ok {
					resultat[len(resultat)-1] = texte
					continue
				}
			} else if mot == "(bin)" {
				if texte, ok := bintodecimal(dernier); ok {
					resultat[len(resultat)-1] = texte
					continue
				}
			} else if mot == "(up)" {
				if texte, ok := upper(dernier); ok {
					resultat[len(resultat)-1] = texte
					continue
				}
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
