package main 

func cap(mot string) (string, bool) {
	if mot == "" {
		return "", false
	}
	premierelettre := mot[0]
	if premierelettre >= 'a' && premierelettre <= 'z' {
		premierelettre -=  32
	}
	motCap := string(premierelettre)
	for i := 1 ; i < len(mot) ; i++ {
		motCap += string(mot[i])
	}
	return motCap, true
}