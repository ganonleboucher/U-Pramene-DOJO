package main

import (
	//"container/list"
	"fmt"
)

func main() {
	var beers = []string{"Svijany", "Kozel", "Bernard", "Radegast"}
	var beerStock = []int{5, 20, 12, 1}

	fmt.Println("Hello", ",World!", "Hello", "again")
	fmt.Println(beers[0])
	for i := 0; i < 4; i++ { //TODO La bite ??
		fmt.Println("Il reste", beerStock[i], "de la biere", beers[i], "!")
	}

	//SystemPropertiesAdvanced.exe
	//"editor.parameterHints.enabled": false

	/*
		beerStock := list.New()
		beerStock.PushFront("Svijany")
		beerStock.PushFront("Kozel")
		beerStock.PushFront("Bernard")
		beerStock.PushFront("Radegast")
	*/
}
