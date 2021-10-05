package main

import "fmt"

func main() {
	m := [][]int{{1,2,3}, {4,5,6}, {7,8,9}}

	fmt.Println(m)

	Rotate(m)

	fmt.Println(m)
}

func Rotate(m [][]int) {
	Reverse(m)

	for i := 0; i < len(m); i++ {
		for j := 0; j < i; j++ {
			m[i][j], m[j][i] = m[j][i], m[i][j]
		}
	}
}

func Reverse(m [][]int) {
	i, j := 0, len(m) - 1

	for i < j {
		m[i], m[j] = m[j], m[i]

		i++
		j--
	}
}