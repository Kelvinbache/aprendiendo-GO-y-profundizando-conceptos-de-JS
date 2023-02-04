/** 
 * @ new Number methods o nuevo metodos de numero
 * @ new methods math o nuevo metodo matematico  
 */

/* examen de aprendido de hoy
   1)Crear un sistema donde pueda pasar un texto  y convierte en numero (calcular la raiz cuadrada del numero)
   2)crear otro sistema donde puedas calcular un triangulo  
*/


/*2) sistema mediante un objecto 
* Crear una clase donde podamos pasar un texto y convertir en numero 
* calcular los tres lados de un tringulo
* retornar los valores de la clase 
*/

class tringulo{
     constructor(arriba,ezquierda,derecha){
          this.arriba = arriba;
          this.ezquierda = ezquierda;
          this.derecha = derecha;
     }
}

const numero = new tringulo("4","8","4");

const arriba =  numero.arriba;
const izuierda = numero.izuierda;
const derecha = numero.derecha;

function resultado() {
 const numero1 = Number(arriba);
 const numero2 = Number(izuierda);
 const numero3 = Number(derecha);

 return numero2 * numero3 - Math.PI(numero1) /180;
}

console.log(resultado());

/*1) sistema de creasion de texto en numero 
* usar una funcion donde pases un texto numerico
* hacer una constante donde se convierta en numero 
* retornar el valor atrior y el actual
 
const sistema = (texto) => {
  const numero = Number(texto);  // Estamos pasomos un texto que tiene un los numerico 
  return `valor anterior: ${texto} | valor nuevo: ${Math.sqrt(numero)}`; // Despues hacemos la operacion entre valor de texto, retornamos un los resultados
}

console.log(sistema("8"));
console.log(sistema("6"));
console.log(sistema("78"));
console.log(sistema("24848"));
*/

/*Primero iremos metodos matematicos y convertidor de numeros


// const Numero = "4"; // Number es un meto que convierte los valores tipo texto en numero
 
// if( typeof Numero == 'string') {  //podemos hacer gualquier tipo de comparacion entre los valor 
//      const creador = Number(Numero); 
//      console.log(creador + 5);
// }

*/
/* metodos matematico nos permite tener todos los metodos matematicos 
 
const numero = 8;
console.log(Math.SQRT1_2 + numero)
*/