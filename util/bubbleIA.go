package utils

import "strconv"

func TriBubbleIAInt(tab [][]string, col int) [][]string {

	if len(tab) <= 1 {
		return tab
	}

	//converti en nombre une seule fois
	keys := make([]int64, len(tab))
	for i := range tab {
		keys[i], _ = strconv.ParseInt(tab[i][col], 10, 64)
	}

	//la partie triée remonte à la fin, on réduit la portée à chaque passe
	for fin := len(tab) - 1; fin > 0; fin-- {
		isEnd := true
		for i := 0; i < fin; i++ {
			if keys[i] > keys[i+1] {
				tab[i], tab[i+1] = tab[i+1], tab[i]
				keys[i], keys[i+1] = keys[i+1], keys[i]
				isEnd = false
			}
		}
		//si aucun échange : déjà trié, on arrête
		if isEnd {
			break
		}
	}

	return tab
}
