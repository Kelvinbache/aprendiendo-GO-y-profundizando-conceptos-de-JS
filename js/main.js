/**
 * @arraykeys metos que devuelve los indices de los elementos 
 * @find metodo para recorrer el arreglo y devolver el valores ,mediante un callbak
*/


/**examen para evaluar los conceptos 
 * Crear un sistema donde pasemos un arreglo y devuelva los numero que son mas altos, y borre los demas 
 * Crear otro donde pasemos un arraglo que solo muestre los indices
 * 
 * 1) sistema de valores grades 
 * crear una function donde usemos resto de valores y los guarde dentro, del arreglo
 * hacer una condicion doden verifique los valores mas grande del arreglo
 * retor los indices y valor dentro de un objecto
 * pide que sia una string muestre donde esta el string y valor que tiene
*/

// sistema 2 mostrando los indices de un array
const indece = (...reto)=>{
   const clave = reto.keys(reto);
   for(const index of clave){
     console.log(index)
   }
}

indece(1,231,32,44,45,456,4,8789,7)





//sistema 1 
const sistema1 = (...array)=>{
 array.find((valor,index)=>{

     if(typeof valor == 'string'){
       console.log([{indice:index,valor:valor}])
    } else if(valor > 10){
        console.log({indice:index,valor:valor});
     } else {
        delete array[valor]
     }
 })
}

// sistema1(1,24,"manuel",8,7,"kelvin",4,"estoy",6,48,"hola",99,46,487,654,8,78,978,"hello")


/*************************************************************************************************************/
/*array.find metodo de array  funcion para metro (elementos,indece,array)
  function de forma assendente osea del primer elemento al ultimo; tiene que pasar un valor verdadero para ejecutar correctamente la funcion
  
  se utliza para comparar valores un array, bueno cualquier tipo de valor que tenga el arreglo
  nota: solo puede del volver el valor dentro de la funcion que tiene, de resto no de vuelve nada  
*/

/*ejemplo basico 

const base = (...array)=>{
array.find((elemento)=>{
      if(elemento > 10) console.log(elemento);
})
}

base(1,4,78,31,8,2,1,4);
*/

/*********************************************************************************************************** */
/*uso basico de array keys 
 next metodo de pasar elementos dentro de un objecto,pero solo  funciona con objecto 


const array = [1,2,3];
const keys = array.keys(array);
const rest = [...array.keys()];

console.log(keys.next());
// for(const clave of keys){ // estamos pasando el keys del array 
//     console.log(clave); // estamos mostrando los indices de los arreglos 
// }
*/

