/*@como recorrer un ojecto*/


/**metodos del object veremos tres formas de recorrer un objecto 
 * 1) object.entries parametros que seguir [clave,valor] podemos retonar una valores of separarlos 
 * 2) object.value solo tiene un unico parametro [value] como su nombre indica solo devuelve el valor 
 * 3) object.keys solo devuelve la calve del objecto [clave] 
*/

const ejemploKeys = {a:1,b:2,c:3}
// console.log(Object.keys(ejemploKeys)) solo devuelve la clave del objecto

const clave = Object.keys(ejemploKeys);
const valor = Object.values(ejemploKeys);


[clave,valor].forEach((elemeto) => {
  if (elemeto.length == 0){
     console.log(elemeto);

  } else if (elemeto.length == 1){
   console.log(elemeto)
  }

});

 /* const ejemploValue = {numero1:1, numero2:2,numero3:3}
// console.log(Object.values(ejemploValue)); solo devuelve el valor 
for(const value of Object.values(ejemploValue)){
     console.log(value); // solo devuelve el valor del objecto 
}
*/
 /* const ejemploEntries = [{nombre:"kelvin",apellido:"abache",telefono:12345345}]

// ejemploEntries.find( valores => { //ejemplo de objecto dentro de array
//       for(const [clave,valor] of Object.entries(valores)){
//            console.log(clave,":",valor)
//       }
// })

//console.log(Object.entries(ejemploEntries)) // esto retorna un arreglo ten encuenta eso

// for(const [clave,valor] of Object.entries(ejemploEntries)){ // una forma tambien para recorrer objectos mediante este metodo
//      console.log(`clave: ${clave} || valor:${valor}`)
// }

*/