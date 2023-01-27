// Tamabie tenemos que llamar un paqueta para empesar
package main

import (
	"encoding/json"
	"fmt"
)

// Tienes que ponder los nombres en MAYUSCULAS cuando utilizes el (Json)
type datos struct {
	Names     string
	Lastnames string
	Bodys     string
}

func validador(name, lastname, body string) {
	//Almacenando los datos
	almacen := []datos{
		{Names: name, Lastnames: lastname, Bodys: body},
	}
	// Transformando los datos en json
	contenido, erro := json.Marshal(almacen)
	// Condicion que valida la salida de los datos
	if erro != nil {
		panic(erro)
	}
	// Pasondo los datos
	fmt.Println(string(contenido))
}

func main() {
	//Llamamos la funcion pasando los datos
	validador("kelvin", "abache", "hello")
	validador("manuel", "abache", "tengo hanbre")
}
