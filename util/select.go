package utils

import "strconv"

func TriSelectInt(tab [][]string, col int) [][]string {
	//converti en nombre
	keys := make([]int64, len(tab))
	for i := range tab {
		keys[i], _ = strconv.ParseInt(tab[i][col], 10, 64)
	}

	index := 0
	var val int64 = 0

	for i := 0; i < len(tab); i++ {
		index = i
		val = 0

		for j := i; j < len(tab); j++ {
			if (val > keys[j]) || val == 0 {
				index = j
				val = keys[j]
			}
		}

		temp := tab[index]
		tempKey := keys[index]

		for k := index; k > i; k-- {
			tab[k] = tab[k-1]
			keys[k] = keys[k-1]
		}
		tab[i] = temp
		keys[i] = tempKey
	}

	return tab
}
