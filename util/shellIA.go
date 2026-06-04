package utils

import "strconv"

func TriShellIAInt(tab [][]string, col int) [][]string {
	longueur := len(tab)

	if longueur <= 1 {
		return tab
	}

	//converti en nombre une seule fois
	keys := make([]int64, longueur)
	for i := range tab {
		keys[i], _ = strconv.ParseInt(tab[i][col], 10, 64)
	}

	//trouve la plus grande largeur (suite de Knuth : 1, 4, 13, 40...)
	n := 1
	for n < longueur {
		n = 3*n + 1
	}

	for n > 1 {
		//reduit la largeur de trie
		n = n / 3
		//compare la premiere largeur puis on décale
		for i := n; i < longueur; i++ {
			valTemp := keys[i]
			valTempTab := tab[i]
			//on recupere l'index actuel
			j := i
			//tant qu'on n'a pas verifier l'index plus petit
			// et que l'index inferieur est plus grand
			for j >= n && keys[j-n] > valTemp {
				//décale la val vers la droite
				tab[j] = tab[j-n]
				keys[j] = keys[j-n]
				// regarde l'index plus petit
				j = j - n
			}
			// insert la valeur
			tab[j] = valTempTab
			keys[j] = valTemp
		}
	}
	return tab
}
