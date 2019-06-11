package main

import (
	"fmt"
	"math"
	"strconv"

	"./mywebserver"
)

var myws mywebserver.MyWebServer
var versao = "1.2.3"

func main() {
	myws = mywebserver.MyWebServer{}
	myws.New(8080)
	myws.AddHandler("somar", somar)
	myws.AddHandler("potencia/quadrado", potenciaQuadrado)
	myws.AddHandler("versao", versaoHandle)
	myws.Start()
}

// Somar executa a soma de dois números
func somar(req *mywebserver.MyRequest) {
	x := req.HTTPRequest.FormValue("x")
	y := req.HTTPRequest.FormValue("y")

	if x == "" || y == "" {
		fmt.Fprintln(req.HTTPResponse, "Falta(m) parâmetro(s).")
	} else {
		fx, err := strconv.ParseFloat(x, 32)

		if err != nil {
			fmt.Fprintln(req.HTTPResponse, "Parâmetro X inválido.")
		}

		fy, err := strconv.ParseFloat(y, 32)

		if err != nil {
			fmt.Fprintln(req.HTTPResponse, "Parâmetro Y inválido.")
		}

		fmt.Fprintf(req.HTTPResponse, "A soma de %10.5f e %10.5f é %10.5f", fx, fy, fx+fy)
	}

	req.Done()
}

// potenciaQuadrado eleva ao quadrado o valor informado
func potenciaQuadrado(req *mywebserver.MyRequest) {
	x := req.HTTPRequest.FormValue("x")

	if x == "" {
		fmt.Fprintln(req.HTTPResponse, "Falta parâmetro.")
	} else {
		fx, err := strconv.ParseFloat(x, 32)

		if err != nil {
			fmt.Fprintln(req.HTTPResponse, "Parâmetro X inválido.")
		}

		fmt.Fprintf(req.HTTPResponse, "O quadrado de %10.5f é %10.5f", fx, math.Pow(fx, 2))
	}

	req.Done()
}

// versaoHandle informa a versão do web service
func versaoHandle(req *mywebserver.MyRequest) {
	fmt.Fprintln(req.HTTPResponse, "A versão da minha API é", versao)
	req.Done()
}
