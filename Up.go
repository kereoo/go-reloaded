package main 

func upper(mot string) (string, bool) {
	if mot == "" {
		return "", false
	}
	motMaj := ""
	for _, lettre := range mot {
		if lettre >= 'a' && lettre <= 'z' {
			lettre -=  32
		}
		motMaj += string(lettre)

	}
	return motMaj, true
}