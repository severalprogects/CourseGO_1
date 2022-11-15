package main

import "fmt"

func main() {
	const PI float64 = 3.1415
	var arr1 [3]int = [3]int{1, 2, 3}
	arr2 := [...]float64{3.14, 5.1, 11.5}
	fmt.Println((arr1[0]+arr1[2] == 4) || (arr2[0] == PI))
}
