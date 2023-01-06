class aplicacion{
   constructor(){
      this.json = fetch("http://127.0.0.1:5500/archivo.json")
    }
}
console.log(aplicacion.json);

/***
 * cosas que hacen para mini proyecto 
 * 1)conectar el archivo json a los js y go
 * 2)mostrar los datos en consola 
 * 3)manejador de errores 
 * 4)poder achadir otros datos al json 
 */
