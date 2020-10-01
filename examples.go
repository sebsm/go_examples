package main

import (
	"fmt"
)

var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah Smith Sade"
	doctorNumber int    = 15
	season       int    = 20
)

func main() {
	var i int
	i = 10
	var j int = 30
	k := 40.
	fmt.Println(i)
	i = 50
	var div int = 7
	var vid int8 = 8
	fmt.Println(i)
	var stat bool = true
	fmt.Printf("%v, %T\n", j, j)

	fmt.Printf("%v, %T\n", k, k)

	fmt.Printf("%v, %T\n", stat, stat)

	fmt.Println(i + div)
	fmt.Println(i - div)
	fmt.Println(i * div)
	fmt.Println(i / div)
	fmt.Println(i % div)

	fmt.Println(i + int(vid))

	fmt.Println(i & div)
	fmt.Println(i | div)
	fmt.Println(i ^ div)
	fmt.Println(i &^ div)

}
