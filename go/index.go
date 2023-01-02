// Tamabie tenemos que llamar un paqueta para empesar
package main

//Aqui tenemos que importar unas funciones para usar

import (
	"fmt"
	"reflect"
)

func array(elemeneto, elemeneto2 string, indixe int) {
	var array []string

	var valor, valor2 = reflect.ValueOf(elemeneto).Kind(),
		reflect.ValueOf(elemeneto2).Kind()

	if valor == reflect.Int && valor2 == reflect.Int {
		fmt.Println("no haceto valores numericos son lentras")
	} else {
		array = append(array, elemeneto, elemeneto2)
	}

	fmt.Println(`los objectos del array:`, array[indixe], `\ donde estan ubicado:`, indixe)
}

func main() {
	array("pelota", "bate", 0)

	array("pez lobo", "manzana", 1)

	array("chocolate", "tomate", 1)

	array("galleta", "maltiada de proteina", 0)
}

//Output: Company is {ABC [{Deven 10000 Full-Stack Developer} {Alex 7000 Back-end Developer}]}
