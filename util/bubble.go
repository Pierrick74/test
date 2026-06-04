package utils

func TriBubble(tab []int) {

  var isEnd bool = false;
  
  for {
	if (isEnd != false) {
		break;
	}

    isEnd = true;
    for  i:=0; i<len(tab)-1; i++ {
        if(tab[i] > tab[i+1]) {
            temp := tab[i];
            tab[i] = tab[i + 1];
            tab[i + 1] = temp;
           isEnd = false;
       }
     }
  }
}