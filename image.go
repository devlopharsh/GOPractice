package main

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		row := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			row[j]=uint8(( i ^ j))
		}
		result[i]=row
	}
	return result
}
