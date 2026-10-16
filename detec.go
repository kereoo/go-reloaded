package main

import "strings"

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

	fusion := []string{}
	for i := 0; i < len(liste_mots); i++ {
		mot := liste_mots[i]
		if strings.HasPrefix(mot, "(") && !strings.HasSuffix(mot, ")") && i+1 < len(liste_mots) {
			mot = mot + " " + liste_mots[i+1]
			i++
		}
		fusion = append(fusion, mot)
	}
	liste_mots = fusion

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
			} else if mot == "(low)" {
				if texte, ok := lower(dernier); ok {
					resultat[len(resultat)-1] = texte
					continue
				}
			} else if mot == "(cap)" {
				if texte, ok := capitalize(dernier); ok {
					resultat[len(resultat)-1] = texte
					continue
				}
			} else if nom, n, ok := lireCommande(mot); ok {
				debut := len(resultat) - n
				if debut < 0 {
					debut = 0
				}
				for i := debut; i < len(resultat); i++ {
					switch nom {
					case "up":
						resultat[i], _ = upper(resultat[i])
					case "low":
						resultat[i], _ = lower(resultat[i])
					case "cap":
						resultat[i], _ = capitalize(resultat[i])
					}
				}
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
