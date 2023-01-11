/*metodo get en js para perdir datos */
const api = new Request("http://127.0.0.1:5500/archivo.json")


const promesa =  new Promise((resolve, reject) => {
   if(api.status === 504 ) reject(new error("tenemos error de conexion"));

/*primera vez usando fetch de esta manera */
fetch(api).then((result) => result.json())

//recorriendo los datos   
// mejorar el orden de los elementos 
   .then((datos) => {   
      for (const element of datos) {
         const resultado = element;
         //cambia las condicion para que empieze a buscar por los id 
           if(resultado.id === "1") resolve(resultado);
           if(resultado.id === "2") resolve(resultado);
       } 
   });
});

promesa.then(consola => console.log(consola))
promesa.catch(error => console.log(error));
   

/***
 * cosas que hacen para mini proyecto 
 * 1)conectar el archivo json a los js y go
 * 2)mostrar los datos en consola 
 * 3)manejador de errores 
 * 4)poder achadir otros datos al json 
 */
