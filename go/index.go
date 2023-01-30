// Tamabie tenemos que llamar un paqueta para empesar
package main

/* 2)Ejercio de typo de dato basico

   crear una codicion para validar el tipo de datos
   1) funcion para recojer los datos y otra para mostralos
   2) condicion para validar los valores
   3) usar que los datos basicos

*/

// func recojer(valor interface{}) {
// 	if reflect.TypeOf(valor) == reflect.TypeOf(valor) {
// 		fmt.Println(reflect.TypeOf(valor))
// 	} else {
// 		fmt.Println("ningun dato es permitido aqui")
// 	}
// }
// func main() {
// 	recojer("kelvin")
// 	recojer(12.1)
// }

/*examen de los conceptos que apredendiste hoy

Crear un sistema de organizador de valores
1) crear una funcion que resiva los valores
2) hacer una condicion para ver el valor y tipo de dato
3) mostrarlo en consola 4 diferentes valores y su tipo de dato
interface == valor que no sabes cual tipo de valor tenga
*/

/*funcion de valores*/
// func values(value interface{}) {
// 	//aqui estamos dicion si diferente o igual pasa esto por consola
// 	if reflect.TypeOf(value) != nil {
// 		t := reflect.TypeOf(value)
// 		v := reflect.ValueOf(value)
// 		fmt.Println("tipo de dato:", t, "||tipo de valor:", v)
// 	} else {
// 		fmt.Println("tenemos un error")
// 	}
// }
// func main() {
// 	values("kelvin")
// 	values(-154)
// 	values(4.5)
// 	values(true)
// }

/*valor basico valueOf y typeOf
  1)valueof nos permite saber el valor de una varible
  2) typeOf nos permite saber que valor primitivo es
*/

//Tipos de tados que hay golang .

/*
1(String) "hola soy estring"(pondeos hacerlo directo o usamos su nombre abrebiado que es string)
2(boolean) que son true o fasel (pondemos  colocar valor directo o usar bool, pero casi siempre sera false)
3(desimales) 4.5 esto es un decimal(tenemos que poner float64 o float32 dependiendo de la computadora)
4(number enteros) 4 sepreseta como numero sin decimal (int o solo el numero)
5(unit) son longitudes que puede tener un valor o unidad que no tiene signo(pero tiene limite de longitud)(investigar mas sobre esto)
6(rune) podemos hacer transfomaciones entre otras cosas que tendras que investigar mas a fondo
7(complex64 o cumplex128) almacena numeros complejos de pendiendo cual utlises
*/

/*2) boolean formar de escribir el codigo
func main() {
	var n bool
	m := true

	fmt.Println(reflect.TypeOf(n), reflect.TypeOf(m))
	fmt.Println(n, m)
}
*/

/* 3) floatante o float
func main() {
	var n float32
	n = 4.44

	m := 4.5

	fmt.Println(reflect.TypeOf(n), reflect.TypeOf(m))
	fmt.Println(n, m)
}
/*

/* 4) int ejemplo como se puede escribir el codigo
func main() {
	var n int = 45
	m := 4

	fmt.Println(reflect.TypeOf(n), reflect.TypeOf(m))
}
*/

/* 5)unit entero pero sin signo ejemplo de uso
func main() {
	var n uint
	n = -5
	fmt.Println(reflect.TypeOf(n))
}

*/

/* 7)ejemplo del uso complex64 y 128
func main() {
	var n complex64
	n = 2 + 4i

	m := 25 + 1i
	fmt.Println(reflect.TypeOf(n), reflect.TypeOf(m))

}
*/

//DATOS POR DEFECTO EN GOLAND
// type datas struct {
// 	name     string
// 	lastname string
// 	age      int
// }
// // Para pasar datos por defecto en una estructuras de datos
// // fill_defaults {poniendo esto una funcion y conjunto por el valor de variable}
// func (data *datas) fill_defaults() {
// 	if data.name == "" {
// 		data.name = "abac"
// 	}
// 	if data.lastname == "" {
// 		data.lastname = "hga"
// 	}
// 	if data.age == 0 {
// 		data.age = 10
// 	}
// }
// func main() {
// 	person1 := datas{name: "kelvin", age: 2}
// 	person1.fill_defaults()
// 	fmt.Println(person1)
// 	person2 := datas{lastname: "abache"}
// //Cambiamos los datos y si falta uno ponemos uno por defecto
// //ejemplo de pesona 1 y 2 faltan datos distintos y rellena por defecto con fill_defaults
// 	person2.fill_defaults()
// 	fmt.Println(person2)
// }
