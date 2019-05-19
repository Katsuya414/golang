package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scanf("%d %d %d", &x, &y, &z)

 	if x*y*z == 5*7*5 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
