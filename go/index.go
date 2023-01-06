// Tamabie tenemos que llamar un paqueta para empesar
package main

import (
	"fmt"
)

//calculadora con go

type numeros struct {
	numero1 int
	numero2 int
}

func operaciones(dato1, dato2, operacion int) {
	var numero numeros
	numero.numero1 = dato1
	numero.numero2 = dato2

	if operacion == 1 {
		suma := numero.numero1 + numero.numero2
		fmt.Println(suma)

	} else if operacion == 2 {
		reta := numero.numero1 - numero.numero2
		fmt.Println(reta)

	} else if operacion == 3 {
		multiplicacion := numero.numero1 * numero.numero2
		fmt.Println(multiplicacion)

	} else {
		fmt.Println("operacion no encontradad")
	}

}

func main() {
	operaciones(8, 9, 3)

	operaciones(81, 5, 5)

	operaciones(7, 8, 2)
}

//Aqui tenemos que importar unas funciones para usar

// func main() {
// 	array := [10]int{4, 5, 7, 5, 8, 9, 54, 8, 64, 9}
// 	for index, element := range array {
// 		if element != int {
// 			fmt.Println("hay un 5 en este arrlos")
// 		}
// 		fmt.Println("localizando el elemento:", index, "elemento:", element)
// 	}
// }

//forma de estructural los distintos campos de informacion
// donde estamos definiendo la structura del arreglo

//puedes guardar diferentes datos
// type persona struct {
// 	name     string
// 	apellido string
// 	pais     string
// 	telefono int
// }
// func main() {
// 	// donde estamos pasando los datos
// 	kelvin := persona{"kelvin", "abache", "venezuela", 1213155652}
// 	fmt.Println(kelvin)
// 	jose := persona{"jode", "peres", "peru", 464654564654}
// 	fmt.Println(jose)
//para poder cambiar los datos de la estructura
// 	kelvin.telefono = 456465465465
// 	fmt.Println(kelvin)
// }

// // i es el para metro para poner el tiempo de ejecucucion
// func index(elemento, elemento2 string, i int) {
// 	objecto, objecto2 := (elemento), (elemento2)
// 	fmt.Println(objecto, objecto2)
// }
// func main() {
// 	i := 10
// 	//retrasa a una funcion
// 	defer index("pelota", "palmera", i)
// 	//Aqui estamos poniendo el tiempo de ejecucion
// 	i = 20
// 	fmt.Printf("usando un manejador de tiempo    ")
// }

/*conceptos para aprender hoy
4)bloque de datos con var
*/
