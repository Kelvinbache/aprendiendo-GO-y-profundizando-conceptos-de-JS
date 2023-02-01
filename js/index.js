/**
 * @ejercicio (separar los elementos mediante uso de oebjectos)
 * Crear una api donde el usuario guarde los datos
 * 1)Crear un objecto donde pueda para los datos 
 * 2)hacer un funcion donde puedas guardar un array
 * 3)otra donde puedas se para los datos 
 * 4) una condicion donde puedas separa datos numeros y strings
 * 5)usar herencia donde pueda meter mas informacion sobre el
 */

class api{
     constructor(nombre,apellido,edad){
          this.nombre = nombre;
          this.apellido = apellido;
          this.edad = edad;
     }
 
 get almacen(){
      return [this.nombre,this.apellido,this.edad];
  } 

}

const persona = new api("kelvin","abache",18);
let a,b,c ;

[a,b,c] = persona.almacen;
console.log(a,b,c)

/** clases y objecto this
 * 
 * @class (classe en js )
 * @this  obejcto this es una forma de ejecucion es como una funcion, tambien puede hacer diferentes llamadas a distintas funciones
 

/**
 * 1)Primero veremos 3 formas de declarar un clase (esta declaracion simple, declaracion anonima, declaracion con dos parametros)
 * 2)Consttrutor y forma del cuerpo en las clases (constructor es otro meto para, crear parametros o objectos dentro de una clase)
 * 3)Como llamar una clase(usamos new dentro una constante para poner nuevos valores)
 * 4)metodos de la clase(metos que podemos usar para clases son muchos pero los simples seria: funcion,get,set,stict)
 * 5)herencia (utilizando entends y super {permite usar los valores del primera clase})
 * 
 

/*1) forma simple de declarar una clases
class ejemplo {
     constructor(nombre,apellido){ // Primero declaramos los objectos o parametros del constructor
          this.nombre = nombre;
          this.apellido = apellido;
}
 
   almacen(){
     return {nombre:this.nombre,apellido:this.apellido};
   }
}

//5)herencia(extens y super (permite hacer herencia de otras clases)
  class ejemplo2 extends ejemplo {
        // tenemos que hacer un nuevo parametro para poner nuevos valores  
      constructor(nombre,apellido,edad){
          // super permite hacer una herencia de otras clases(nota: solo podemos pedir los datos )
          super(nombre,apellido);
          this.edad = edad; // this nos da una referencia un objecto o funcion
     }    
     almacen(){
          return{nombre:this.nombre,apellido:this.apellido,edad:this.edad};
     }
  }

/*3) como llamar una clase  (utlizamos new para crear nuevos de elementos)
  const persona = new ejemplo("kelvin","abache");
  console.log(persona);

  const persona1 = new ejemplo2("kelvin","abache",18);
  console.log(persona1);

*/

/* 1)Formas de declarar una clases 
//Forma declarar simple
/*comenzamos usando una palabra llama class (Que significa clase),despues ponemos un nombre a la clase 
class obejcto {
       // despues usamos otra palabra llamada constructor(Que significa construir) despudes ponemos nombre a las variable (parametor o nombre), usamos this para crear una variable y despues ponemos, el mismo valor 
     
       constructor(juego,telefono){ //Declarar un objecto que no puede duplicar, el mismo valor 

          // falta investigar un copo para saber mas
         this.juego = juego;
         this.telefono = telefono;
     }

//metodo de declarar una funcion 
     play(){
          return this.juego + " " +this.telefono; 
     }
     
// aqui estamos utlizando un metodo (llamado get, que esto significa que va ser global)  
    get alamacen(){
          return this.play()
     }


}


// new (hacermos nuevos valores para pasarlos)
const reultado = new obejcto("fifa","motorola");
console.log(reultado.alamacen);
*/

/* forma de expresion sin nombre de clase

// Aqui no ponemos un nombre a la clase, pero usamos una constante para declarar la clase
const obejcto2 = class{
      //mismo condicion para declarar los parametros del constructor
     constructor(juegos,telefonos){
          this.juego = juegos;
          this.telefonos = telefonos;
     }
}
*/

/* forma de expresion con dos nombres 

//Declar dos veses una variable pero utilizando otro nombre para diferenciar el padre del hijo
// const persona = class persona1 {
//      constructor(nombre,apellido){ 
//          this.nombre = nombre;
//          this.apellido = apellido; 
//      }
 }
 */