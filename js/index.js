//funcion creadora de elementos 

function creardorDeElemntos(name,lastname,body) {
   //primero ponemos una condicion para verificar los elementos
  if(typeof name === 'string' && typeof lastname === 'string'&& typeof body === 'string') return [name,lastname,body];
  else return 'estos no son elementos tipo string';
}

function pasar([name,lastname,body] = creardorDeElemntos) {
   const person = {
      name: name,
      lastname: lastname,
      body:body
   }

 return person;
}


/***
 * cosas que hacen para mini proyecto CREANDO DATOS JSON
 * 1) crear una funcion que ordene los datos mediante uso de objectos
 * 2) crear otra funcion que separe los datos, colocando un id a cada elemento
 * 3) crear una forma de mostrar los datos requeridos 
 * 4)poner una funcion que permita una conexion entre dos funciones
 * donde esta guardado el codigo (https://github.com/Kelvinbache/aprendiendo-GO-y-profundizando-conceptos-de-JS);
 */
