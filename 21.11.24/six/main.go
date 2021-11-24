package main

import "fmt"

func find20(a int) (int, bool) {
	for i := 1; i <= 9; i++ {
		if a*i == 20 {
			return i, true
		}
	}
	return 0, false
}

func main() {

	var find bool = false
	i := 1
	j := 1
	for ; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if i*j == 45 {
				find = true
				break
			}
		}
		if find {
			break
		}
	}
	fmt.Printf("%d* %d= %d\n", i,j, i*j)
}