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

type Caisse struct {
	startDay float64
	balance  float64
}

func getIndexBiere(nomBiere string, stock []BeerProduct) int { //find algorithm => rechercher l'index d'un element dans une collection a partir d'un critère discriminant ( ici le nom de la bière )
	for i := 0; i < len(stock); i++ {
		if nomBiere == stock[i].name {
			return i
		}
	}
	return -1
}

func withdrawFromStock(nomBiere string, quantiteDemandee uint, stock []BeerProduct) uint {
	fmt.Println("Je voudrais", quantiteDemandee, "de", nomBiere, "svp.")
	var indexBiere int = getIndexBiere(nomBiere, stock)
	if indexBiere == -1 {
		fmt.Println("Nous n'avons pas ou plus de", nomBiere, "en stock !")
		return 0
	} else if quantiteDemandee > stock[indexBiere].quantity {
		fmt.Println("J'ai pas assez de", stock[indexBiere].name, "sale poch' !")
		return 0
	} else {
		stock[indexBiere].quantity -= quantiteDemandee
		fmt.Println("Il reste desormais", stock[indexBiere].quantity, "de la biere", stock[indexBiere].name, "en stock !")
		return quantiteDemandee
	}
}

func showAllStock(stock []BeerProduct) {
	for i := 0; i < len(stock); i++ {
		fmt.Println("Il y a en tout", stock[i].quantity, "de", stock[i].name, "en stock !")
	}
}

func showItemStock(name string, stock []BeerProduct) {
	var indexBiere int = getIndexBiere(name, stock)
	if indexBiere == -1 {
		fmt.Println("Il n'y a pas de", name, "en stock !")
	} else {
		fmt.Println("Il y a", stock[indexBiere].quantity, "de la biere", name, "en stock !")
	}
}

/*func ajoutCash(beerWithdrawn uint, caisse Caisse, stock []BeerProduct) uint {
	var indexBiere int = -1
	for i := 0; i < len(stock); i++ {
		if prixBiere == stock[i].price {
			indexBiere = i
			break
		}
	}
}*/

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
	//var caisse = Caisse{startDay: 80.37, balance: 0}
	var svijany = BeerProduct{"Svijany", 2.50, 20}
	var stock = []BeerProduct{
		svijany,
		{"Kozel", 3.50, 12},
		{"Bernard", 1.80, 42},
		{"Radegast", 2.70, 27},
	}

	fmt.Println("===============================")
	showAllStock(stock)

	fmt.Println("=============VENTE DE KOZEL==================")

	var beerWithdrawn = withdrawFromStock("Kozel", 2, stock)
	fmt.Println("beer withdrawn :", beerWithdrawn)
	//ajoutCash("Kozel", stock, beerWithdrawn, caisse)

	fmt.Println("===============================")
	withdrawFromStock("Radegast", 220, stock)
	fmt.Println("===============================")
	withdrawFromStock("Sibeeria", 22, stock)

	fmt.Println("===============================")
	showAllStock(stock)

	fmt.Println("===============================")
	showItemStock("Radegast", stock)
	fmt.Println("===============================")
	showItemStock("Svijany", stock)
	fmt.Println("===============================")
	showItemStock("Sibeeria", stock)
	fmt.Println("===============================")

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
