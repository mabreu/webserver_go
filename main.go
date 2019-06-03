package main

import (
	"fmt"

	"./mywebserver"
)

// Somar soma 2 numeros
func somar(req *mywebserver.Request) {
	fmt.Fprintln(req.GetResponseWriter(), "A soma é")
	req.Done()
}

// Subtrair subtrai 2 numeros
func subtrair(req *mywebserver.Request) {
	fmt.Fprintln(req.GetResponseWriter(), "A substração é")
	req.Done()
}

// potencia_quadrado subtrai 2 numeros
func potenciaQuadrado(req *mywebserver.Request) {
	fmt.Fprintln(req.GetResponseWriter(), "A potencia quadrado é")
	req.Done()
}

// Subtrair subtrai 2 numeros
func potenciaCubo(req *mywebserver.Request) {
	fmt.Fprintln(req.GetResponseWriter(), "A potencia cubo é")
	req.Done()
}

func main() {
	var meuws = mywebserver.MyWebServer{}
	meuws.New(8080)
	meuws.AddHandler("somar", somar)
	meuws.AddHandler("subtrair", subtrair)
	meuws.AddHandler("power/square", potenciaQuadrado)
	meuws.AddHandler("power/cube", potenciaCubo)
	meuws.Start()
}
