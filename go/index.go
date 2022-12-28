// Tamabie tenemos que llamar un paqueta para empesar
package main

//Aqui tenemos que importar unas funciones para usar
import "fmt"

// asi es como colocamos los tipos de variable que podemos colocar
func simple(name string) {
	fmt.Print("helllo word " + name)
}

// esta es la funcion que inicia todo el programa
func main() {
	simple("kelvin")
}
