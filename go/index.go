// Tamabie tenemos que llamar un paqueta para empesar
package main

//Aqui tenemos que importar unas funciones para usar
import (
	"fmt"
)

func main() {
	fmt.Println()
}

//pequena comparacion de declaracion de condiciones y variables
// func main() {
// 	dinero := 50

// 	if dinero >= 10 || dinero <= 10 {
// 		pan := 10
// 		cafe := 3
// 		pastel := 50

// 		if dinero >= pan {
// 			fmt.Println("pan comprado")
// 		} else {
// 			fmt.Println("pan:", "Dinero no alcanza")
// 		}

// 		if dinero >= cafe {
// 			fmt.Println("cafe comprado")
// 		} else {
// 			fmt.Println("cafe:", "no tienes dirnero suficiente")
// 		}

// 		if dinero >= pastel {
// 			fmt.Println("pastel comprado")
// 		} else {
// 			fmt.Println("pastel:", "esta muy costoso")
// 		}
// 	}
// }

// esto un bloque de declaracion de variables y esto es global
// var (
// 	name     = "kelvin"
// 	lastname = "abache"
// 	numero   = 2004
// 	telefono = 46545646564
// )

// func main() {
// 	fmt.Println("esto es nombre:", name)
// 	fmt.Println("esto es apellido:", lastname)
// 	fmt.Println("esto es numero:", numero)
// 	fmt.Println("esto es telefono:", telefono)
// }

//forma de declara una condicion en GO
// func main() {
// 	x := true
// 	name := "kelvin"

// 	//forma de declara una condicion
// 	if x {
// 		lastname := "abache"

// 		if x != false {
// 			fmt.Println(name)
// 			fmt.Println(lastname)
// 			fmt.Println(name + " " + lastname)
// 		}
// 	}

// 	fmt.Println(x)
// }

// func main() {
//     //forma corta de escribir una variable
// 	name := "kelvin"
// 	number := 2000
// 	verdad := true

// 	fmt.Println(reflect.TypeOf(name))
// 	fmt.Println(reflect.TypeOf(number))
// 	fmt.Println(reflect.TypeOf(verdad))
// }

// forma de claracion en go

// //Aqui estamos represendando un numero
// var i int

// //Aqui estamos representando un string
// var s string

// func main() {

// 	//iniciamos una mariable y despues seguimos declarando pero ojo cuando declaras
// 	var name, lastname string = "kelvin", "abache"

// 	//Aqui estamos dividiendo los numeros por cada lentra
// 	k, l, m := 1, 2, 3

// 	//igual aqui cada item = dinero y numero = 2000
// 	item, numero := "Dinero", 2000

// 	//Aqui estamos pengando el nombre y apellido
// 	fmt.Println(name + " " + lastname)

// 	//aqui estamos sumando los numeros
// 	fmt.Println(k + l + m)

// 	//igual aqui estamos pegando dinero y numero
// 	fmt.Println(item, "-", numero)
// }
