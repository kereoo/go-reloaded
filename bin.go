package main

func bintodecimal(mot string) (string, bool) {
	if mot == "" {
		return "", false
	}
	nombre := 0
	for _, character := range mot {
		if character != '0' && character != '1' {
			return "", false
		}
		nombre = nombre*2 + int(character-'0')
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
