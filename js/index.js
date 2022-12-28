const dinero = 5;

if(dinero >= 10 || dinero <= 10){
    const pan = 10;
    const cafe = 3;
    const pastel = 50;

    if(dinero >= pan){console.log("pan comprado")} 
    else console.log("pan:","no tienes dinero cara comprar"); 
 
    if(dinero >= cafe){console.log("cafe comprado")} 
    else console.log("cafe:","El dinero no alcanza");

    if(dinero >= pastel){console.log("pastel comprado")} 
    else console.log("pastel:","no tienes suficiente dinero");
}