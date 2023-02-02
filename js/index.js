/**
 * examen de los metos y conceptos del otro dia:

 * 1)Crea una forma vitual para pasar una frase y comparar los primeras letras de la frase 
 * 2)conversacion de entre persona (compara las frases de las dos persona  y devuelva true o false)

 1(Comparador de primera letra) pasos a seguir:
 * crear un objecto donde los parametros sean frase1 y frases2
 * hacer una funcion donde se compare las primeras letra
 * depues pasar en consola (son igules o no son iguales)
*/

const frases = (frase1,frase2) =>{
     if(frase1.startsWith("A") == true) return 'es igual';
}


console.log(frases("Aaaah"))


/**
 * @stringsWith (puede devolver true o falso) 
 * @includes (permite conparal las frases o oraciones de variables)
 * @endsWith (puede devolver true o false)
 */

/*StrinsWith devuelve true o false dependiendo la condion que tiene, que es la primera letra es mayuscula en una oracion o nombre
// const strin = 'Estoy aqui';
// //estos Busca las letras mayusculas de una oracion o nombre 
// console.log(strin.startsWith('Es'));
// console.log(strin.startsWith('es',1))
*/

/*include podemos revisar si es igual a la oraciones o frases

const frase = 'estoy fuera de liga';

 const liga = 'estoy en mi liga';

//podemos comparar dos valores entre si con estos metodos 
 console.log(`"${frase}" ${frase.includes(liga) ? 'is': 'is not'}`);

*/

/*endsWith compara las ultimas palabras una frase o oracion 
const frase = 'kelvin tiene un telefono';
// los para metros que resive es (el valor final de la oracion o frase)
console.log(frase.endsWith('telefono'))
*/