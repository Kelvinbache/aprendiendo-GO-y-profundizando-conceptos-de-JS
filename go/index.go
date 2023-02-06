package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

/*examne del conocimiento aprendido
crear un ruta donde puedas petir datos y separar los datos pedidos mediante un bloque de tipos de datos

paso a seguir \:
1)hacer una funcion donde tenga el codigo de peticion
2) comprovar el estado de la peticion
3) hacer condion donde verifiquemos la url
4) separa los datos
5) mostrar los datos por consola que separdon por llaves
*/

type datas struct {
	Id    int
	Name  string
	email string
}

func main() {
	jsons, error := http.Get("https://jsonplaceholder.typicode.com/users")
	if error != nil {
		panic(error)
	}

	defer jsons.Body.Close() //por defecto los datos van a estar cerca
	cuerpo, error := ioutil.ReadAll(jsons.Body)

	if error != nil {
		panic(error)
	}

	repuesta := string(cuerpo)

	log.Panicf(repuesta)
}

/*
 1) metodos de peticion
 2)codigo de estado
 3)codigo de respuesta
 4)codigo de errores
*/

/* metodo Get
metodo get podemos pedir datos mediante la url de la pagin
func main() {
	respuesta, error := http.Get("https://jsonplaceholder.typicode.com/users")
	if error != nil {
		log.Fatal(error)
	}

	cuerpo, error := ioutil.ReadAll(respuesta.Body)
	if error != nil {
		log.Fatal(error)
	}

	repuestaDelServidor := string(cuerpo)
	log.Printf(repuestaDelServidor)
	log.Printf(respuesta.Status) // aqui estamo usando codigo de estado
}
*/
