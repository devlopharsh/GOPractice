package main 

import (
	"fmt"
)

type Sample struct{
	x int
	y string
}

func main(){
	// sample :=Sample{1, "Hello"}	
	// fmt.Println(sample.x, sample.y)
	//arrays in go 
	var ar [5] int
	for i:=0 ; i<5 ; i++{
		ar[i]=i
	} 
	fmt.Println(ar)

}