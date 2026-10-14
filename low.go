package main 

func lower(mot string) (string, bool) {
	if mot == "" {
		return "", false
	}
	motMin := ""
	for _, lettre := range mot {
		if lettre >= 'A' && lettre <= 'Z' {
			lettre +=  32
		}
		motMin += string(lettre)

	}
	return motMin, true
}