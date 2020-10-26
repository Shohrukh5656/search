package main

import (
	"context"
	"fmt"





	
	"github.com/Shohrukh5656/search/pkg/search"



)


func main() {

	res := search.Any(context.Background(), "HTTP", []string{"./test.txt", "./test copy.txt"})

	r, ok := <-res
	if !ok {
		fmt.Println("error ok =>", ok)
	}

	//fmt for for 
	fmt.Println("---------------")
	fmt.Println("res.Phrase) => ", r.Phrase)
	fmt.Println("res.Line) => ", r.Line)
	fmt.Println("res.LineNum) => ", r.LineNum)
	fmt.Println("res.ColNum) => ", r.ColNum)
	fmt.Println("---------------")
}
