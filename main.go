package main 

import (
	"fmt"
)

type Sample struct{
	x int
	y string
}

func main(){
	sample :=Sample{1, "Hello"}	
	fmt.Println(sample.x, sample.y)
}