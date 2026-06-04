package main

import (
	utils "example/hello/util"
	"fmt"
	"slices"
)

func main() {
	datas, err := utils.ReadCSV("small.csv");

	if err != nil {
    	fmt.Println("Error:", err)
    	return
	}

	header := datas[0];
	//remove de j à j+1 not include
	datas = slices.Delete(datas, 0 , 1);

	orderDatas := utils.TriBubbleInt(datas, 9)

	orderDatas = append([][]string{header}, orderDatas...)

	utils.WriteCSV("output.csv", orderDatas)
}