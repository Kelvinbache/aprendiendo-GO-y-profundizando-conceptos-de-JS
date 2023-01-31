/**
 * 1)Examen de progreso utilizando los dos conceptos crear un sistema para separar numero y letras
 * criterio: los numero se guardaran en arreglos y las letras en objectos
 * 
 * Pasos a seguir:
 * 1) Crear por un lado, la funcion que resive y otras que separe los datos
 * 2) utilizar una condicional para ver tipo de valor que resive, la funcion
 * 3) dividir donde estara las funciones 
 * 4) utliza los conceptos vienjo que aprediste ayer 
 */


/*creando la funcion que resive FALTA RESOLVER ESTE EJERCICIO*/

import { tipoDeValor } from "./main.js"

const almacen = (valor1, valor2,valor3, valor4)=>{
     const contenido = [valor1,valor2,valor3,valor4];
     tipoDeValor(contenido);
}
almacen("kelvin",5,4,5)



/** Modulos nota: Para usar los modulos con node.js y js tiene que cambiar el tipo de valor ,a module(archivo package.json)
 * 1(import y export) usando las parala export y import (podemodemos exportar una funcion o otro tipo de valor) (los mismo podemos hacer con import podemos importar quier valor que pidamos)
 
* 2(exporta multiple valores) muy util al momento de exportar diferentes valores 

*/

/* 2)Exportar con multiples valore 
function ejemplo(params) {
     return [1,2]
}

function ejemlo2 (){
     return [3,4]
}

asi podemos exportar multiples variables 
export {
     ejemplo,ejemlo2
}

*/

/* 1)import y export basica
estamos exportando la funcion a otro archivo que se encarga de ejecutar esa funcion
export function ejemplo(){
     return [1,2,5,8]
}     
*/

/********************************************OTROS CONCEPTOS***************************************************************************** */

/*Destructuring en js*
 * @description metodo de destructuring en js(nos permite quitar algunos valor y asinalos )

* 1(desestruracion basica) como podemos desetructurar y cual es su sintaxis @valor  

* 2(intercambiar valor) podemos intercambiar mediante un array entre algorimos de bajo nivel 

* 3(funcion que retorna un array) Podemos returnar cualquier valores en un arrglo (pero tambien podemos separar las variables)@ejemlo

* 4(ignorar valores) utilizando [,,] dos comas dentro de un array podemos ignorar todos los valores, y [a,,b] podemos ignorar un      valor@ejemplo 
 
* 5(asignar valores) Podemos asignar el resto de valores a una sola variable utilizado [...c] tres punto que significa asigna , el resto de la cade a este algorimo

 * 6(Desestructuracion de objectos o asignar un valor) Podemos desestructurar mediante un const y asignar objetos utiliznado ({}) los parentesis y los corchetes 
 
 
* 7(asignar nuevos nombre a las variables) Podemos cambiar los nombres de objectos y poner otros nuevos, pero tiendo encuenta que el nombre tiene tener referencia al antiguo
*/
/*asignacion de nuevos valoresen objetos 
const objecto = {
     user: "kelvin",
     id:1
}

//Podemos cambiar algunos nos nombres de varias o cambiar los valores, pero ten encuenta que para hacer un cambio tienes que utilizar 
{varible actual: variable nueva} = donde esta guardad 

forma con variable de alto rango 
const {user:name,id:adreen} = objecto;
console.log(name,adreen);

forma con una varible de menor rango 
let {user:nombre, id:direcion} = objecto;
console.log(nombre,direcion);


*/
/* desestruracion de objeto 

//forma 1 de hacer una desestructuracion de objetos
const objecto = {
     user : "kelvin",
     id: 1
};

//misma forma de hacer una desestructuracion pero aqui estamos utilizando una const para guardar los valores 
const {user,id} = objecto
console.log(user,id)

//forma 2 de hacer una desestructuracion de objetos o asignar valores a las variables

let a,b;

//utilizando la sintaxis de la siguiente forma podemos asignar valores ({a,b} = {a:2,b:2}), utilizando los parentesis 
({a,b} = {a:1,b:2});
console.log(a,b)
*/
/*asignar el resto de valores de una cadena 

gracias a los tres ... podemos asignar el resto de valores a una sola variable (pero ojo mejor usar esto de ultimo para no tener errores)

function ejemlo(){
     return [1,3,48,4,7,8,78,7];
}

let a,b,c;
[a,b,...c] = ejemlo();
console.log(a,c,b)

*/

/* 4)ignorar valores 

let a,b; 

podemos ignorar valores usando (a, ,b) y salta al siguente valor de la cadena, o tambien no mostrar nada (,,) asi ignora todo los valores
[a,,b]=[1,2,3];
console.log(a,b)

function ejemlo(){
     return [1,2,3,4];
}

let c,d,e;
[a,,d,e]= ejemlo();
console.log(c,d,e)

*/

/* 3) funcion que retorna un arreglo 

function ejemplo() {
 retorna un arreglo que puede terne calquier valor
     return [1,2]
}

asignamos los valor a y b 
let a,b;

mediante la llamada ejemplo a y b recojeran los valores que retorne, y los guardamos en esa variables
[a,b] = ejemplo();
console.log(a,b)
*/

/* 2) intercambio de valores mediante un algorimo de bajo nivel 
let a = 1;
let b = 2;

Podemos cambian valores en alugonos casos mediante algorimos de bajo nivel como let (aqui estamos intercambiando los valore a y b)
[a,b]= [b,a];
console.log(a,b)
*/

/* 0)valores predeterminados al momento de desetructurar  
let a,b ;

puede pasar que los valores se cambien 
[a = 1, b= 0] = [2,6];
console.log(a,b);
*/

/* 1)a y b escoge los primeros valores de array, pero rest recoje los valos que quedaron (utilizando ...)los 3 punto  asignamos a rest que los recoja

[a,b] = [1,2];
console.log(a,b);
Podemos hacerlo de este modo otro otras forma 
let c,d
const valor = [1,2];
[c,d] = valor;
*/


/** (const y let los template strings)
 * @constant (solo podemos leer el valor, mas no podemos cambiar su valor)
 * @let (podemos cambiar el valor y podemos utilizalo en sierto rango del codigo)
 * @template strings (podemos hacer frases mas extensas y poner valor adentro de ellos, usando ${})
 */

/*podemos cubrir siertas zonas del codigo, osea esto es gobla 
let nombre,apellido;


function pasamos(nombres,apellidos){
     nombre = nombres;
     apellido = apellidos
     //podemos hacer condiciones desde siertas partes del codigo 
     if(typeof nombre == "number") return `no es valido los numero en una operacion de texto`;
 
     if(typeof apellido == "number") return `no es valido los numero en una operacion de texto`;
    
     return `hola soy ${nombre} ${apellido}`
}

console.log(pasamos("kelvin",2))
*/

// (Podemos definir una funcion o cualquir valor con let y cont 
// let valores = (nombres,apellidos = ` `)=>{
//     const nombre = nombres;
//     const apellido = apellidos;
//     if(typeof nombre == "number") return `valores numerico no validos`;
//     if(typeof apellido == "number") return `valores numerico no validos`;
//     return `hola soy ${nombre} ${apellido}`
// }
// console.log(valores("kelvin",5))

/** (funcion flecha y valores por defecto)
 * @ejemplo (parametro) podemos tener cualquier valor 
 * @funcionFlecha (parametros = podemos poner un valor por defecto si no hay valor agulno de entrada) los mismos pero mas corto y mas facil de usar
 * typeOf (podemos saber que valor)
 * @returns (Podemos hacer hacer algunas conparacion y otras cosas, a a parte de sumar o restar)
 * 
 * mini programa para ver si entendiste---------------------------------------------------------------
 * Crear dos functiones para pasar cualquier valor y otra para ver que valor tiene
 * 1)crear la function donde se ubicaran los valores por defecto
 * 2)crear la function donde se ubicaran los typos de valor 
 * 3)separa los valor con constantes
 * ADICCIONAL: usar los dos tipos de functiones  
 */

//function flecha 
// const defecto = (a,b = ["otros", true])=>{
//     return b;
// }
// const tipoDeValor = ()=>{
//     const valor = defecto("kelvin");
//     if (valor == false)return "valor no son igual"
//     return typeof valor 
// }
// console.log(tipoDeValor());





//function normal
//function donde esta los para metros por defecto
// function defecto (a,b = [{c :"hello"}]){
//     return {a,b}
// }


// function tipoDeValor (){
//  const valor1 = defecto("kelvin");
//   return  valor1
// }

// console.log(tipoDeValor());





















//funcion flecha o arrow function
//  const funcionFlecha = (a,b)=> {
//    return a + b 
//  }
//  const resultado = funcionFlecha("4","5");
//  console.log(typeof resultado);

 


//funcion pasndo parametros por defecto
// function ejemplo(a,b){
//     return a + b 
// }
// const guardarValor = ejemplo();
// console.log(guardarValor);