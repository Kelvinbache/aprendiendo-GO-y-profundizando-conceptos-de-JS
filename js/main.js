/**
 * @symbol(es un typo valor unico que combina los valores primitivos en js)pero cunado usas synbol el valor   es unico no puedes comparar, el mismo valor porque es diferente
 * 
 * @function rest Parametro (las sintaxis es la suiente ...)estos quiere decir agarra el resto de valor y guardalos en un array, otra cosa que quieras hacer con eso 
 */

/*symbol
 sonsola(nos muestra un valor unico que es una etiqueta llamada symbol o simbolo)
 1) metodos en symbol
 *for(clave del simbolo) nos permite reconocer si la clave pasada o actual son iguales 
 *

/forma simple de crear un simbol y como crear un symbol con clave/ 
// const valor1 = Symbol.for("hola");
// const valor2 = Symbol.for("hola");
// // const valores = Object.getOwnPropertySymbols(valor1);
// console.log(valor1 == valor2)

//no podemos comprar los valores del symbol porque son unicos 
// //te devuelve falso 
// console.log(valor1 === valor2);

/ejemplo de symbol/

const clave = Symbol();
const myCajon = {};
myCajon[clave] = 'hola';
myCajon["nombre"] = "kelvin";

console.log(myCajon); //mostrando todo el contenido 

console.log(clave in myCajon); //Preguntado si clave esta en cajon 

console.log(Object.getOwnPropertySymbols(myCajon)); // revisando si cajos tiene un objecto symbol

console.log(Object.keys(myCajon)); // comprobando la llave de cajo, pero no la de symbol 

console.log(myCajon[clave]); //mostrando la llave de del cajon

console.log(Reflect.ownKeys(myCajon)); //muestra las clave del objecto y del symbol

*/

/*examen de symbol (Crear un objecto 2 objecto unicos y comparar los valores de ello) 
 acciones para tomar:

 1)Crear una funcion para pasar los valores por defecto y transformalos en symbol
 2)hacer una compracion entre esos valores y mostrar en consola true o false 
 3)El valor a retornar tiene que estar un objecto 

2) examen de symbol(crear una variable que su valor sea un objecto y que tenga una clave para acceder)
1) Crear una constante donde su valor sea un objecto 
2) hacer que symbol tenga la clave para acceder a ese objecto
3) comprobar si symbol esta en ese objecto en cuestion 

*/

/*primer ejercicio de symbol 
const transformacion = (valor1,valor2)=>{
  const cambiar1 = Symbol.for(valor1);
  const cambiar2 = Symbol.for(valor2);
   
  if (cambiar1 == cambiar2) return {'valor':'true'}; 
  else return {'valor':'false'};
};

console.log(transformacion(1,1))



const clave  = Symbol();
const cajon = {};

cajon[clave] = 'hola';
cajon['objectos'] = ['pelota'];

console.log(Object.getOwnPropertySymbols(cajon));
console.log(cajon[clave]);
*/

/**************************************************************************************************************************************** */
/*funtion Rest parament 

//forma simple 

//... resto de valor dentro del parametro una funcion, siempre pasa un arraglo del resto de valores
function Resto (a,b,...otros) {
  return  otros  
}

console.log(Resto(1,2,5,34,6))

// puedes hacer varias formas para entender mejor el lenguaje
//nota: no puedes poner dos veses restos, porque despues tienes un error
const valor =(a,b,...Resto)=>{
  for (const elemto of Resto) {
    const suma1 = elemto + b ;
    const suma2 = elemto + a;
      console.log(`primera suma:${suma1}, segunada suma:${suma2}`);
  }
}

valor(1,3,4,64,8,498,7,9,556,4,87,5,87,897,97)

//nota: tambien puedes utilizar metos de array como map,for,len,entre otros 

*/

/**@example de restos una funcion o examen de ello 
 * 1) crear una funcion donde los principales parametros puedan (sumar y restas,el resto de los valores)
 * pasos a seguir:
 * 1)crear una funcion donde recoja mas de 10 valores numerico 
 * 2)hacer una condicion donde sia un valor de texto muestre error, en hacer la operacion
 * 3)pasar en un arreglo los valor de la suma y resta 
 * 4)utilizando metos de js6 
 * 5)pero ahora el usuario quiere ver mejor los valor y diferencialos


const recojedor = (a,b,...resto)=>{
  for (const valores of resto) {
    if (typeof valores === 'number'){
        const suma = valores + a;
        const resta = valores + b;
        console.log([`suma: ${suma} | resta:${resta}`]);
    }
                                      
    else {
        console.log(`valores no numerico `, valores)
    }
  }
}

recojedor("1",3,48,34,8,"holas",748,4,6,"tenemos","hola",8,48,48,45,489,4)

*/

/**modificacion del codigo anterior utilizando variantes de js6 */
const recoletor = (a,b,...c)=>{
   for(let i = 0; i <= c.length; i++){
     if (typeof c[i] == 'string'){
        console.log("valor no numrico", c[i].length)
     } else {
        let suma,resta;
        suma = a + c[i];
        resta = b - c[i];
         
         console.log([`primer resultado: ${suma}, segundo resultado:${resta}`])
     }
   }
}

recoletor(12,45,645,87,89,4564,"hola","hello","mia");

setTimeout(()=>{console.clear()},10000);
