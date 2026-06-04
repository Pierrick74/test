package utils

import "strconv"

func TriInsertInt(tab [][]string, col int) [][]string {

	if len(tab) == 1 {
		return tab
	}

	for i := 0; i < len(tab); i++ {
		temp := tab[i]
		tempNumber, _ := strconv.ParseFloat(tab[i][col], 64)
		j := i

		for j > 0 {
			prev, _ := strconv.ParseFloat(tab[j-1][col], 64)
			if prev <= tempNumber {
				break
			}
			tab[j] = tab[j-1]
			j -= 1
		}

		tab[j] = temp
	}
	return tab
}
