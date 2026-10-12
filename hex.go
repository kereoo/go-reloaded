package main

func hexatodecimal(mot string) (string, bool) {
	if mot == "" {
		return "", false
	}
	nombre := 0
	for _, character := range mot {
		tempo := 0
		if character >= '0' && character <= '9' {
			tempo = int(character - '0')
		} else if character >= 'a' && character <= 'f' {
			tempo = int(character-'a') + 10
		} else if character >= 'A' && character <= 'F' {
			tempo = int(character-'A') + 10
		} else {
			return "", false
		}
		nombre = nombre*16 + tempo
	}
	if nombre == 0 {
		return "0", true
	}
	texte := ""
	for nombre > 0 {
		texte = string(rune('0'+nombre%10)) + texte
		nombre = nombre / 10
	}
	return texte, true
}