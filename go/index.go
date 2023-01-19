// Tamabie tenemos que llamar un paqueta para empesar
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type datos struct {
	names     string
	lastnames string
	bodys     string
}

func validador(name, lastname, body string) {
	valor := reflect.ValueOf(name)

	if valor == reflect.ValueOf(name) {
		panic("no podemos retornar un numero")
	} else {
		almacen := []datos{
			{names: name, lastnames: lastname, bodys: body},
		}

		contenido, erro := json.Marshal(almacen)

		if erro != nil {
			panic(erro)
		} else {
			fmt.Println(string(contenido))
		}
	}
}

func main() {
	validador("kelvin", "abache", "hello")
	validador("kelvin", "abache", "hello")
}
