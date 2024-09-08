package main

import (
	"fmt"
)

func main() {
	var i int /* not ready to initialize yet */
	i = 50
	var j string = "naveen"
	x := "what?" // ready to initialize
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(x)
	fmt.Printf("%v, %T \n", i, i)
	fmt.Printf("%v, %T \n", j, j)
	fmt.Printf("%v, %T \n", x, x)

	var y float32 = 3.147
	fmt.Printf("%v, %T \n", y, y)

	var (
		actorName string = "Al pacino"
		movieName string = "Ragin Bull"
		companion String = "Pain"
		partNo    int    = 1
	)

}
