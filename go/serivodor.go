package main

import (
	//Paquete que nos ayuadad a realizar o crear un servidor

	"log"
	"net/http"
)

/* Usa la misma sintaxis de express pero pero cambiar algunas cosas ejemplo
// 1)tenemos que inicia con un metodo (HandleFunc) despuedes pone ("/") puedes dejarlo vacio o ponerle nombre

// 2)Parametros de respuesta y requimiento (utilizamos una letra para diferenciar)ejemplo (w metodo, r metodo)

// 3)Metodos de pantalla y repuesta (http.ResposeWrite) (*http.Request) usando estos dos metodos podemos hacer un servidor basico

// 4)Metodo multi pantalla o pagina (hacemos una variable que se llame mux) dentro usamos (http.NewServeMux) pasando los mismo parametros que igual a las funciones

// 5)Subarboles en las rutas

// 6) conexion no encontrada o 404

*/

/*Forma simple de hacer un servidor
// func main() {
// 	http.HandleFunc("/locos", func(w http.ResponseWriter, r *http.Request) {
// 			w.Write([]byte("<h1>hola soy kelvin</h1>")) //w.Write([]byte (repuesta con etiquetas de html))
// 	})                                                  //pintame esto en la pantalla

// // muestrame esto en consola cuando inicie http.ListenAndServe(puerto,error en conexion)
// 	log.Fatalln(http.ListenAndServe(":3000", nil))
// }
*/

/*Forma de convertir la respuesta del servidor una funcion independiente

 //Aqui dividimos los campos del servidor poniendo, la repuesta y requerimiento en otra funcion
func servidor(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>cm estas soy otra parte de la repuesta</h1>"))
}

func main() {
// Despues llamando la funcion para utlizar, y teniendo codigo mas limpio
	http.HandleFunc("/", servidor)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
*/

/* Multi pagina o multiples rutas

func servidor(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hola soy el pensamiento del ser humano</h1>"))
}

func main() {
	mux := http.NewServeMux()                    // Crear o agrega nuevas rutas al servidor
	mux.HandleFunc("/", servidor)                // poniedo el mismo metodo que utilizamos al momento de usar http
	mux.HandleFunc("/locos2", servidor)          // Podemos poner nombres a las rutas entre otras cosas
	log.Fatal(http.ListenAndServe(":8080", mux)) // lofFatal nos alerta de cualquier tipo de error al momento de y iniciar el servidor
    //mux perosnaliza y tiene que usarse para no tener errores
}

*/

/* subrutas o subarboles en las rutas

// func servidor(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("<h1>hola soy subruta</h1>"))
// }

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/tienda", servidor) //Aqui tenemos una ruta fija
// 	mux.HandleFunc("/tienda/productos", servidor) //Aqui tenemos una subRuta o subArbol

// 	log.Fatal(http.ListenAndServe(":8080", mux))
 }
*/

/* 404 o ruta no encontrada

// func servidor(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {  //condicion que que quiere decir verifica que las ruta tenga una barra expaciadora
// 		http.NotFound(w, r) // respuesta si la condicion es verdadera pasa error
// 	}
// 	w.Write([]byte("<h1>error en busqueda</h1>"))
// }

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/servido", servidor)
// 	log.Fatal(http.ListenAndServe(":8080", mux))
}
*/

// Redirecionamiento

func servido(url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, url, 301)
	}
}

func main() {
	http.HandleFunc("/ejemplo", servido("http://ejemplo"))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
