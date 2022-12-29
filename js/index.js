const total = [];

function suma(numero1,numero2){
    const array = [];

    const resultado = numero1 + numero2;
    array.push(resultado);

    return array
}

function restal(numero1,numero2,numero3){

   const resta = suma(numero1,numero2);
   const resultado = numero3 - resta;
    
   return total.push(resultado)
}


restal(7,5,98)

restal(4,59,98)

restal(9,9,98)

restal(4,58,54)


for (let  i = 0;  i < total.length; i++){ 
      console.log(total[i]);
}

// //forma de hace una busqueda de triangulo js  
// function suma(x,y){
//    const total = 2 * (x + y);
//    console.log(total)
  
//    const area = (x * y)
//    return area
// }

// console.log("area del triangulo", suma(12,45));