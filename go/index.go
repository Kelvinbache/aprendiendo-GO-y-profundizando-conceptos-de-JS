package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

// Examen del los conceptos de hoy
/*crear un sistema que pueda crear archivos,comprobar esos archivos, poner contenido en ellos

1)Crear una estructura que podamos pasar los datos
2) pasar la estructura donde estan los datos guardados, a la funcion
3) crear una funcion donde podamos pasar cualquier nombre al archivo
4) hacer una comprobacion del archivo,para ver su existencia
5) padar en consola una notia donde digamos que el archivo fue creado

*/

// estructura de datos

type usuarios struct {
	Nombre   string
	Apellido string
	Telefono int
	Edad     int
}

//funcion donde esta el codigo de creacion

func creadorDeArchivo(nombreDeArchivo string) {

	archivo := nombreDeArchivo + ".json"     //pasndo y usando una cocadenacion para crear el archivo
	contenido, errores := os.Create(archivo) //creando el archivo

	if errores != nil {
		log.Fatal("tenemos problemas para crear el arcivo")
	}

	defer contenido.Close() //diciendo que el archivo estara serca

	if errors.Is(errores, os.ErrNotExist) { //haciendo la comporbacion del archivo creado
		log.Fatal("archivo no existe")

	} else {
		usuario := []usuarios{ // crando la informacion con la base de datos
			{Nombre: "kelvin", Apellido: "abache", Telefono: 45465465465, Edad: 18},
			{Nombre: "kelvin", Apellido: "abache", Telefono: 45465465465, Edad: 18},
			{Nombre: "kelvin", Apellido: "abache", Telefono: 45465465465, Edad: 18}}
		jsons, _ := json.Marshal(usuario)                // creando un json para despues pasar la informacion
		_, error := contenido.WriteString(string(jsons)) //solo permite formato estring

		if error != nil {
			log.Fatal("tenemos problemas colocando el contenido")
		}
		fmt.Println("contenido colocado en json") //notificacion en la consola para decir que , el archivo fue creado con exicto
	}
}

func main() {
	creadorDeArchivo("user")
	creadorDeArchivo("user2")
}

/*colocando contenido el los archivos mediante el uso de os

func main() {
	archivo := "texto.txt"
	contenido, error := os.Create(archivo) //creando el archivo

	if error != nil {
		log.Fatal("no podemos crear el archivo")
	}

	defer contenido.Close() // contenido esta cerca

	nombres := []string{"kelvin", "mabuel", "rebeca", "santiago"}

	for _, nombre := range nombres { // recorriendo el arreglo donde estan los datos
		_, erro := contenido.WriteString(nombre + "__") // permite colocar informacion dentro de los archivos

		if erro != nil {
			log.Fatal("error en mostrar informacion") // condicion donde la informacion no este
		}
	}
	fmt.Println("archivo en pantalla") // noticia cuando termina de hacer todo
}
*/

/*elimir archivos metoddo de os
func main(){
  archivo := "./go/archivo.json" // Pasando la direcion del archivo que quieres eliminar

  error := os.Remove(archivo) // solo llenva el parametro del error

  if error != nil{ // comprobando los errores entre los archivos
	log.Fatal("error en eliminar archivo")
  }
  fmt.Println("archivo:",archivo)
}
*/

/*metodo de comprobaciones entre los diferentes archivos, para validar su existencia
func main() {
	//json := "./go/archivo.json"
	archivo := "texto.txt"

	_, error := os.Stat(archivo) //comienza con este archivo

	if errors.Is(error, os.ErrNotExist) { //condicion que buscar entre los archivos

		fmt.Println("archivo no existe")

   } else {
		fmt.Println("archio existe")
	}

}
*/

/* os metodos en los archivos
func main() {

	archivo := "texto.txt"

	contenido, error := os.Create(archivo)//Permite crear archivos,por defecto los crear afuera de los archivos
	//hacemos un codicion para comprobar el contenido

	defer contenido.Close() // por defecto el contenido estara serca

	if error == nil { //comprobando el error al momento de crear el archivo
		log.Fatal(error)
	}

	log.Fatal(contenido) // pasando cuando el archivo fue creardo
}
*/
/* os paquete que permite comunicar los archivos externos en go
func main() {
	archivo := "./go/archivo.json"

	abriendo, error := os.Open(archivo)//abre el archivo
	if error != nil {
		fmt.Println(error)
	}

	contenido, error := abriendo.Stat() //comienza el archivo
	if error != nil {
		fmt.Println(error)
	}
	log.Panicln(contenido)
}

*/
