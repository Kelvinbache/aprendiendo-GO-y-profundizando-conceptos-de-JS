function sistemaDeOrganizacion(elemento,elemento2,indice){
 const almacen = [];

 //condicion que estamos colocando a las varibles
     if(typeof elemento === 'number' && typeof elemento2 === 'number' ){
        return 'no puedo almacenar numeros';
    } else almacen.push(elemento,elemento2);
   
// retorna los indices de los elementos y los busca 
  return almacen[indice]
}

console.log(sistemaDeOrganizacion("persona","persona2",1));


console.log(sistemaDeOrganizacion("persona3","persona4",1));


console.log(sistemaDeOrganizacion("persona5","persona6",1));






/*crear un sistema de organizacion
1)donde cada elementos sea separado por indice 
2)poder ver elemento deseado y no ver otros elementos
3)poder crear nuevos elemetos y aumentar la cola
4)separar numeros y string
*/

// //forma de hace una busqueda de triangulo js  
// function suma(x,y){
//    const total = 2 * (x + y);
//    console.log(total)
//    const area = (x * y)
//    return area
// }

// console.log("area del triangulo", suma(12,45));