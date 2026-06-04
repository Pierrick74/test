package utils

import "strconv"

func TriBubbleInt(tab [][]string, col int) [][]string {

  var isEnd bool = false;

  for {
	if (isEnd != false) {
		break;
	}

    isEnd = true;
    for  i:=0; i<len(tab)-1; i++ {
        a, _ := strconv.ParseFloat(tab[i][col], 64);
        b, _ := strconv.ParseFloat(tab[i+1][col], 64);
        if (a > b) {
            temp := tab[i];
            tab[i] = tab[i + 1];
            tab[i + 1] = temp;
           isEnd = false;
       }
     }
  }

  return tab;
}