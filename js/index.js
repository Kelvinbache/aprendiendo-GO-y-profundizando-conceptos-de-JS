function calculador(numero1,numero2,datos){
   class operacion {
     constructor(numero1,numero2){ 
     this.suma = numero1 + numero2
     this.resta = numero1 - numero2
   }
 };

if (datos === 1){
   const suma = new operacion(numero1,numero2);
   return suma.suma

} else if(datos === 2){
   const resta = new operacion(numero1,numero2);
   return resta.resta
};

}

console.log(calculador(45,6,2))


// function persona(nombre,apellido){
//    this.nombre = nombre,
//    this.apellido = apellido
// }

// const kelvin = new persona("kelvin","abache")
// console.log(kelvin);






/**
 * construir una calculador con objectos que puedas poner un numero en especifico para hacer una operacion
 * 
 * 
 */