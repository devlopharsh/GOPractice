// package main

// import "fmt"

// // import (
// // 	"fmt"
// // )

// // import "golang.org/x/tour/pic"

// // type Sample struct{
// // 	x int
// // 	y string
// // }
// type Vertex struct{
// 	Lat , Long float64
// }

// var m map[string]Vertex

// func main(){
// 	// // sample :=Sample{1, "Hello"}
// 	// // fmt.Println(sample.x, sample.y)
// 	// //arrays in go
// 	// var ar [5] int
// 	// for i:=0 ; i<5 ; i++{
// 	// 	ar[i]=i
// 	// }
// 	// fmt.Println(ar)
// 	// pic.Show(Pic)
// 	m = make(map[string]Vertex)
// 	m["Bells"]=Vertex{
// 		40 ,50 ,
// 	}
// 	fmt.Println(m["Bells"])
// }

package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	count := make(map[string]int)

	for _,word := range words{
		count[word]++
	}
	return count
}

func main() {
	wc.Test(WordCount)
}
