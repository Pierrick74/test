package utils

import "strconv"

func TriSelectIAInt(tab [][]string, col int) [][]string {
	//converti en nombre
	keys := make([]int64, len(tab))
	for i := range tab {
		keys[i], _ = strconv.ParseInt(tab[i][col], 10, 64)
	}

	for i := 0; i < len(tab)-1; i++ {
		min := i

		for j := i + 1; j < len(tab); j++ {
			if keys[j] < keys[min] {
				min = j
			}
		}

		if min != i {
			tab[i], tab[min] = tab[min], tab[i]
			keys[i], keys[min] = keys[min], keys[i]
		}
	}

	return tab
}
