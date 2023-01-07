// Tamabie tenemos que llamar un paqueta para empesar
package main

import (
	"encoding/json"
	"fmt"
)

type datos struct {
	name   string
	number int
	body   string
}

func main() {

	person := datos{"kelvin", 465548798, "hello"}

	v := json.Marshal(person)
	b := []byte(`{"name":"kevin","lastname":"abache","body":"hello"}`)

	fmt.Println()

}
