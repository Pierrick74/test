package utils

import "strconv"

func TriInsertIAInt(tab [][]string, col int) [][]string {

	if len(tab) <= 1 {
		return tab
	}

	//converti en nombre une seule fois
	keys := make([]int64, len(tab))
	for i := range tab {
		keys[i], _ = strconv.ParseInt(tab[i][col], 10, 64)
	}

	for i := 1; i < len(tab); i++ {
		temp := tab[i]
		tempNumber := keys[i]
		j := i

		for j > 0 && keys[j-1] > tempNumber {
			tab[j] = tab[j-1]
			keys[j] = keys[j-1]
			j -= 1
		}

		tab[j] = temp
		keys[j] = tempNumber
	}
	return tab
}
