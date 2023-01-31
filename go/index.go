// Tamabie tenemos que llamar un paqueta para empesar
package main

import (
	"fmt"
	"time"
)

/*importado*/

/*Ejercios de practica
3) crear una matriz donde puedas guardar los dia y mes de las pesonas, tambien poder mostralos en consola

3) Ejercicio (Creador de matriz donde se guarda los meses y dias)
* Crear un array donde agarra valor numero
* hacer una condicion donde los parametros sean numerico y no tipo string
* retorne un array donde salga los valores del mes y dia
* utilizar paquete de tiempo
*/

func valores(valor1, valor2 int) {
	ano := time.Now().Year()
	mes := time.Now().Month()

	var almacen []int = []int{ano, int(mes)}

	fmt.Println(almacen)

}

func main() {
	valores(4, 5)
}

/* 2 Ejercicio Etructura de objectos (Pasos a seguir)
* Crear un funcion que resiva los parametro
* guarda los datos en un arreglo
* Pasar los datos a otra funcion
* Poner orden los elementos



/Creando los objectos donde se guarda las variables
type data struct {
	Nombre   string
	Apellido string
	Edad     int
	Cuerpo   string
}

func personas(nombre, apellido, cuerpo string, edad int) {
	//pasando los valores un contenedor
	persona := []data{
		{Nombre: nombre, Apellido: apellido, Cuerpo: cuerpo, Edad: edad},
	}

	// Utilizando un bucle para recorrer los arreglo,despues guardar los datos en una variable
	// Manera de recorrer los arreglos pero recuerdad esto hay otras maneras solo tienes que buscar
	for i := 0; i < len(persona); i++ {

		// Guardando los valores de rrecorrido de persona y despues usar los valores como objectos normales
		valor := persona[i]
		fmt.Println("nombre:", valor.Nombre, "||", "apellido:", valor.Apellido, "||", "edad:", valor.Edad, "||", "Cuerpo:", valor.Cuerpo)
	}
}

func main() {
	//Pasando los datos a la funcion personas
	personas("kelvin", "abache", "hola", 44)
	personas("santi", "abache", "como esta", 6)
	personas("marina", "fuentes", "bien", 13)
}

*/

/* 1) Ejercicio de clave y valor| | pasos a seguir:
 * Primero hacer una map
 * pasar una clave des una funcion para y comparar los valores
 * crear la codicional para comparar los dato y si no tiene mismo valor borrarlo
 * despues hacer un consola para mostrar

// func almacen(clave, valor string) {
// 	caja := make(map[string]string)

// 	// Para entender mejor nombre de la variable[clave] || valor = variable[clave y valor] el corchete almacena los dos valores
// 	caja[clave] = valor

// 	// primero comparamos los valores recordando que clave, va tener el contenido del valor, mas la clave que tiene
// 	if valor == caja[clave] {
// 		fmt.Println("valor de caja: ", caja[clave])
// 	} else {
// 		// Si la condicion no cumple los valores requeridos
// 		delete(caja, clave)
// 		fmt.Println("valores eliminados")
// 	}
// }

// // utlizando una llama de otra funcion para utilizarla
// func main() {
// 	almacen("kelvin", "hola")

// 	almacen("kelvin", "como estas")
// }
*/

/*------------------------------------------------Otro concepto---------------------------------------------------------------------------*/

/*Concepto de tiempo
  1)Como llamar al tiempo y usar los metos basico
  2) Poner datos o consfigurar datos
*/
/*2) cambiar datos o modificarlos
func main() {
	//podemos pasar datos y manipular como queramos(esto resive 7 valor{ano,mes,dia,hora,minutos,segundos,tiempo de localida})
	times := time.Date(2023, 1, 30, 8, 51, 30, 100, time.Local)
	fmt.Println(times)

	//formato o formatiar la como quieres que se vea
	fmt.Println(times.Format("2023-1-30 08:051:3 "))
}

*/

/* 1) como llamar al tiempo

// func main() {
// 	//now nos tira todo los datos de hora y tiempo
// 	times := time.Now()
// 	// podemos pedir diferentes datos mediante los siguientes metodos

// 	//podemos imprimir todos los valores en console
// 	fmt.Println(
// 		times.Year(),
// 		times.Month(),
// 		times.Day(),
// 		times.Hour(),
// 		times.Minute())
// }
*/

/*------------------------------------------------Otro concepto--------------------------------------------------------------------------*/
/* Concepto de estructura de datos

1)Array {nos permite cuandar datos en forma de matriz(poniendo indicadores de cuantos datos van entrar y salir)}
2)Slice {Permite poner datos sin pertir una parametro en especifico}
3)Map {Es igual a un array pero esto usar clave y valor para haceder al contenido}
4)Struct {Podemos separar los diferentes valores, mediante valores unicos de una estructura}

*/

/* 1) array tipo de dato que almacena como una matriz
func main() {

	// Cuando ponemos el numero dentro del array, significa que va tener diez entradas de datos, o puede ser menos de pende que nesecites
	var x [10]int

	//len para ver la longitud que pueda tener un array o una matriz
fmt.Println("led:", len(x))
	fmt.Println("emp:", x)
}
*/

/* 2)Slice o revanada
// func main() {
// 	// Crear dos array con los vimos valores, pero separa por un igual esto una forma mas flexible de guardar datos
// 	//  tambien permite no indicar cuantos cortes va tener
// 	var x []int = []int{1, 4, 8, 7, 4, 4484, 84, 8, 5}

// 	fmt.Println("logintud:", len(x))

// 	fmt.Println("valor:", x)
 }
*/

/* 3) Map funciona con clave y valor

// func main() {
// 	// aqui estamos iniciando el map dentro de los corchetes iniciamos un valor |generalmete iniciamos con string despues asignamos el valor
// 	m := make(map[string]int)

// // Aqui podemos ver que adentro de los corchetes ponemos la clave y despues el valor
// 	m["kelvin"] = 40

// // mostramos por consola el valor, pero ojo nombre de la variable y clave
// 	fmt.Println(m["kelvin"])

// // delete permite borrar la clave del map
// 	delete(m, "kelvin")

// // ahora el valor vaser 0
// 	fmt.Println(m["kelvin"])
 } */

/*  4)Struct o estructura

// type personas struct {
// 	//cuando creas una estructura recuerdad que los valor van hacer unicos, y no pueden ser cambiados
// 	nombre   string
// 	apellido string
// 	edad     int
// }

// func main() {
// 	//utilizamos una varia para implementar otros datos, a la estructura

// 	//manera simple para guardar un solo dato
// 	kelvin := personas{nombre: "kelvin", apellido: "abache", edad: 16}

// 	// utilizando una array para guardar multiples objectos
// 	persona := []personas{
// 		{nombre: "kelvin", apellido: "abache", edad: 18},
// 		{nombre: "alejandro", apellido: "abache", edad: 45},
// 	}

// 	fmt.Println(persona)
// 	fmt.Println(kelvin)
 } */
