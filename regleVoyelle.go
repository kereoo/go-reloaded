package main

func testVoyelle(mot string) bool {
	for _, r := range mot {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
			continue
		}
		if r >= 'A' && r <= 'Z' {
			r += 32
		}
		return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'h'
	}
	return false
}

func article(phrase string) string {
	mots := []string{}
	conca := ""
	for _, chara := range phrase {
		if chara == ' ' {
			if conca != "" {
				mots = append(mots, conca)
				conca = ""
			}
		} else {
			conca += string(chara)
		}
	}
	if conca != "" {
		mots = append(mots, conca)
	}

	for i := 0; i+1 < len(mots); i++ {
		if mots[i] == "a" && testVoyelle(mots[i+1]) {
			mots[i] = "an"
		} else if mots[i] == "A" && testVoyelle(mots[i+1]) {
			mots[i] = "An"
		}
	}

	output := ""
	for i, w := range mots {
		if i > 0 {
			output += " "
		}
		output += w
	}
	return output
}
