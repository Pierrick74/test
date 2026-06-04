package main

import (
	utils "example/hello/util"
	"fmt"
	"slices"
	"time"
)

func main() {

	defer func(start time.Time) {
        fmt.Printf("temps insertIA fr %v\n", time.Since(start))
    }(time.Now())

	datas, err := utils.ReadCSV("fr.csv");

	if err != nil {
    	fmt.Println("Error:", err)
    	return
	}

	header := datas[0];
	//remove de j à j+1 not include
	datas = slices.Delete(datas, 0 , 1);

	orderDatas := utils.TriInsertIAInt(datas, 9)

	orderDatas = append([][]string{header}, orderDatas...)

	utils.WriteCSV("frO.csv", orderDatas)
}