package main

import (
	//"container/list"
	"fmt"
)

type BigBeerStock struct {
	name     string
	price    float64
	quantity uint
}

func main() {
	//KATA 1
	//var beers = []string{"Svijany", "Kozel", "Bernard", "Radegast"}
	//var beerStock = []int{5, 20, 12, 1}
	//fmt.Println("Hello", ",World!", "Hello", "again")
	//fmt.Println(beers[0])
	//for i := 0; i < 4; i++ { //TODO La bite ??
	//	fmt.Println("Il reste", beerStock[i], "de la biere", beers[i], "!")
	//}

	//KATA 2
	fmt.Println(BigBeerStock{"Svijany", 2.50, 20})
	var svijany = BigBeerStock{"Svijany", 2.50, 20}
	var beers = []BigBeerStock{
		svijany,
		{"Kozel", 3.50, 12},
		{"Bernard", 1.80, 42},
		{"Radegast", 2.70, 27},
	}

	for i := 0; i < len(beers); i++ {
		fmt.Println("Il reste", beers[i].quantity, "de la biere", beers[i].name, "en stock !")
	}
}
