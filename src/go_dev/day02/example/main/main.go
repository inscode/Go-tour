package main

import (
	"fmt"
	"time"
	"go_dev/day02/example/demo"
)

func main() {
	fmt.Println(time.Month(1).String())
	a := 10
	b := 20
	demo.Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
	for i := 1; i < 10; i++ {
		str := demo.IsSushu(i)
		fmt.Println(str)
	}
	demo.GetFlowers()
	//1 2, 6 24 120
	fmt.Println(demo.JS(5))

}