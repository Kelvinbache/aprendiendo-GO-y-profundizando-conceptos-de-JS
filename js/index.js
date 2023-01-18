
function datos(nombre,apellido,objecto) {
   //condiciones para validar los datos(haciendo validacion por parte)
   if(typeof nombre === 'number')  return 'valor no valido';

   else if(typeof apellido === 'number') return 'valor no valido';

   else if(typeof objecto === 'number')  return 'valor no valido';

   else {
        const almacen = {nombres:nombre, apellidos:apellido, objectos:objecto};
        return almacen
   };
};

function datosJson(valores){
   //  const {nombres,apellidos,objectos} = valores;
   const cambiandoJson = JSON.stringify(valores);
   console.log(cambiandoJson);
};

datosJson(datos("kelvin",2,"hello"));

datosJson(datos("manuel","peres","como estas"));



//funcion creadora de elementos 

// function creardorDeElemntos(name,lastname,body) {
//    //primero ponemos una condicion para verificar los elementos
//   if(typeof name === 'number' && typeof lastname === 'number'&& typeof body === 'number') return 'datos numericos no acectados';
//   else return [name,lastname,body];
// }

// function pasar([name,lastname,body] = creardorDeElemntos) {
//    const person = {
//       name: name,
//       lastname: lastname,
//       body:body
// }

// function creadorDeJson(){

//   if(typeof name === 'number' && typeof lastname === 'number' && typeof body === 'number') return {}
//       else {
//           const objectos = JSON.stringify(person);
//          return objectos
//       }
//  }
//   console.log(creadorDeJson());
// }


// pasar(creardorDeElemntos("kelvin","abache","hello"))

// pasar(creardorDeElemntos("kelvin",2,5))



/***
 * cosas que hacen para mini proyecto CREANDO DATOS JSON
 * 2) crear otra funcion que separe los datos, colocando un id a cada elemento
 * donde esta guardado el codigo (https://github.com/Kelvinbache/aprendiendo-GO-y-profundizando-conceptos-de-JS);
 */
