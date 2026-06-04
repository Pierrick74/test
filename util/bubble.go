package utils

import "strconv"

func TriBubbleInt(tab [][]string, col int) [][]string {

	//converti en nombre
	keys := make([]int64, len(tab))
	for i := range tab {
		keys[i], _ = strconv.ParseInt(tab[i][col], 10, 64)
	}

	var isEnd bool = false

	for {
		if isEnd != false {
			break
		}

		isEnd = true
		for i := 0; i < len(tab)-1; i++ {
			if keys[i] > keys[i+1] {
				tab[i], tab[i+1] = tab[i+1], tab[i]
				keys[i], keys[i+1] = keys[i+1], keys[i]
				isEnd = false
			}
		}
	}

	return tab

	/*
	   //Comparaison de string
	   var isEnd bool = false;

	   	  	for {
	   			if (isEnd != false) {
	   				break;
	   			}

	   		isEnd = true;
	   	 	for  i:=0; i<len(tab)-1; i++ {
	   	        if (tab[i][col] > tab[i+1][col]) {
	   	            temp := tab[i];
	   	            tab[i] = tab[i + 1];
	   	            tab[i + 1] = temp;
	   				isEnd = false;
	   	       }
	   	    }
	   	}

	   	return tab;
	*/
}
