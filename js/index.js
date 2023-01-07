async function api(dato){
   const json = await fetch("http://127.0.0.1:5500/archivo.json");
   if(json.status === 500 ) return "tenemos problemas para conectar";
   
   const cotenido = json.json().then( valor => console.log(valor))
   return cotenido
}

console.log(api());


/***
 * cosas que hacen para mini proyecto 
 * 1)conectar el archivo json a los js y go
 * 2)mostrar los datos en consola 
 * 3)manejador de errores 
 * 4)poder achadir otros datos al json 
 */
