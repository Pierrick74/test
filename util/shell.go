package utils

import "strconv"

func TriShellInt(tab [][]string, col int) [][]string {
	longueur := len(tab);
  	var valTemp int64 = 0;
	var valTempTab []string;
  
  	n := 0;
  	//trouve la plus grande largeur 
	for n<longueur {
		n = (3*n+1);
	}

	//converti en nombre une seule fois
	keys := make([]int64, len(tab));
	for i := range tab {
		keys[i], _ = strconv.ParseInt(tab[i][col], 10, 64);
	}

  	for (n!= 0){
  		//reduit la largeur de trie
    	n = int(n/3);
		//compare la premiere largeur puis on décale
    	for i := n; i<longueur; i= i + 1 {
        	valTemp = keys[i];
			valTempTab = tab[i];
        	//on recupere l'index actuel
        	j := i;
        	//tant qu'o n'a pas verifier l'index plus petit 
        	// et que l'index inferieur est plus petit
        	for((j>n-1) &&(keys[j-n]>valTemp)){
	        	//switch de val
				tab[j] = tab[j-n];
           		keys[j] = keys[j-n];
           		// regarde l'index plus petit
           		j = j-n;
        	}
        	// insert la valeur
        	tab[j] = valTempTab;
			keys[j] = valTemp
    	}
  	}
  return tab;
}