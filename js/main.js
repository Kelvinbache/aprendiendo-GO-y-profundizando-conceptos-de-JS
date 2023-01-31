const tipoDeValor = ([valor,valor2,valor3,valor4])=>{
   let a,b,c,d;

    if(typeof valor == "string") return ({a} = {a:valor});
    else [a] = [valor];
   
    if(typeof valor2 == "string") return ({b} = {b:valor2});
    else  [b] = [valor2];
   
    if(typeof valor3 == "string") return ({c} = {c:valor3});
    else  [c] = [valor3];
   
    if(typeof valor4 == "string") return ({d} = {d:valor4});
    else [d] = [valor4];
   
};

export {
    tipoDeValor
}