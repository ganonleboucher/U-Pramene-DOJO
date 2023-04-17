package main

import (
	//"container/list"
	"fmt"
)

type BeerProduct struct {
	name     string
	price    float64
	quantity uint
}

func vendre(nomBiere string, quantiteDemandee uint, stock []BeerProduct) { //nom de la bière, la quantité souhaitée, stock
	fmt.Println("Je voudrais", quantiteDemandee, "de", nomBiere, "svp.")
	var indexBiere int = -1
	for i := 0; i < len(stock); i++ {
		if nomBiere == stock[i].name {
			indexBiere = i
			break
		}
	}

	if indexBiere == -1 {
		fmt.Println("Nous n'avons pas de", nomBiere, "en stock !")
	} else if quantiteDemandee > stock[indexBiere].quantity {
		fmt.Println("J'ai pas assez de", stock[indexBiere].name, "sale poch' !")
	} else {
		stock[indexBiere].quantity -= quantiteDemandee
		fmt.Println("Il reste desormais", stock[indexBiere].quantity, "de la biere", stock[indexBiere].name, "en stock !")
	}
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
	var svijany = BeerProduct{"Svijany", 2.50, 20}
	var stock = []BeerProduct{
		svijany,
		{"Kozel", 3.50, 12},
		{"Bernard", 1.80, 42},
		{"Radegast", 2.70, 27},
	}

	for i := 0; i < len(stock); i++ {
		fmt.Println("Il reste", stock[i].quantity, "de la biere", stock[i].name, "en stock !")
	}

	vendre("Kozel", 2, stock)
	vendre("Radegast", 220, stock)
	vendre("Sibeeria", 22, stock)

	for i := 0; i < len(stock); i++ {
		fmt.Println("Il reste", stock[i].quantity, "de la biere", stock[i].name, "en stock !")
	}

	/*
		var caisse float64 = 0
		var biereDemandee = stock[1]
		var quantiteDemandee uint = 2
		fmt.Println("Je voudrais", quantiteDemandee, "de", biereDemandee.name, "svp.")

		stock[1].quantity -= quantiteDemandee
		// same as: stock[1].quantity = stock[1].quantity - 2
		// *= += ...
		caisse += stock[1].price * float64(quantiteDemandee)
		for i := 0; i < len(stock); i++ {
			fmt.Println("Il reste desormais", stock[i].quantity, "de la biere", stock[i].name, "en stock !")
		}
		fmt.Println("Il y a desormais", caisse, "dollars dans la caisse !")
	*/
}
