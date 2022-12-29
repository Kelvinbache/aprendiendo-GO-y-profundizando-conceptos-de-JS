// Tamabie tenemos que llamar un paqueta para empesar
package main

//Aqui tenemos que importar unas funciones para usar

import (
	"fmt"
)

// creando la funcion para sumar numeros
func suma(numero1 int, numero2 int) int {
	var resultado = numero1 + numero2
	return resultado
}

// creando un bloque de valores
var (
	valor1 = suma(12, 35)
	valor2 = suma(45, 6)
	valor3 = suma(59, 98)
	valor4 = suma(638, 87)
)

func main() {
	//array para almacenar los elementos en una posicion
	var array [4]int = [4]int{valor1, valor2, valor3, valor4}

	//cuando interes una array en for recuerda que es indice y elemento
	for index, elemento := range array {
		total := elemento / 8

		//pasando los valores por consola
		fmt.Println("la localizacion:", index, "valor de los elementos:", total)
	}
}

// segunda forma de declaracion de variables en arrays
// func main() {
// 	x := [5]int{1, 2, 5, 8, 9} //esta inicia con valores
// 	fmt.Println(x)
// 	var y [5]int = [5]int{45, 578, 5488, 8754, 578} // es esta estamos asinando parcialmente los valores
// 	fmt.Println(y)
// 	person := [4]string{"kelvin", "maria", "miguel", "gabriel"} //solo tienes que ponele el tipo de valor que tenga la variable
// 	fmt.Println(person)
// }

//declaracion de array primera forma
// func main() {
// 	var array [3]string
// 	array[0] = "kelvin"
// 	array[1] = "maria"
// 	array[2] = "nada"
// 	fmt.Println(array[1])
// 	fmt.Println(array[0])
// 	fmt.Println(array[2])
// }

// //forma con dos parametro
// func triangulo(x int, y int) (area int) {
// 	// aguerdate que tenemos que poner var cuando estamos dentro del bloque
// 	var parametro int
// 	parametro = 2 * (x + y)
// 	fmt.Println(parametro)
// 	area = (y * x)
// 	return // estamos retornando los valores para esperar valores
// }
// func main() {
// 	//por esto estamos haciendo un retorno de triangulo
// 	fmt.Println("area del tringulo", triangulo(12, 35))
// }

// otra forma de declarar una function
// esta forma no retorna dana todo se hace dentro de,su bloque
// func res(x int, y int) {
// 	total := x - y
// 	fmt.Println(total)
// }

// esto una forma de declarar una function con retorno
// x int, y init = igual espera un para metro para inicia
// init comenzara para dar un retorno los valores a la variable pedida
// func add(x int, y int) int {
// 	total := 0
// 	total = x + y
// 	return total
// }

// func main() {
// 	sum := add(20, 12)
// 	res(12, 35)
// 	fmt.Println(sum)
// }

// //nos retorna 5 veses cualquier recorrido
// 	for range "hello" {
// 		fmt.Println("hello")
// 	}

//forma 1 de declarar un for
// func main() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(i, "<", ">")
// 	}
// }

//forma larga de declarar una constante
// const hello string = "hello"

//forma corta de declarar una constante
// const name = "kelvin"

//bloque de una constante
// const (
// 	name     = "kelvin"
// 	lastName = "abache"
// 	telefono = 6456465464
// 	hello    = "hello word"
// )

// func main() {
// 	fmt.Println(name, lastName, hello)

// 	fmt.Println(telefono)
// }

/*cosas para aprender hoy
5)hacer una mini aplicacion para entender el tema
*/
