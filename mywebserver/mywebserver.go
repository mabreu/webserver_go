package mywebserver

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// MyWebServer é uma estrutura (classe) para representar o meu servidor web
type MyWebServer struct {
	handlers map[string]func(req *Request)
	port     int
	version  string
}

// New inicaliza o servidor Web
func (ws *MyWebServer) New(port int) {
	ws.version = "1.0.0" // Versão de 02/06/2019

	if port == 0 {
		port = 8080
	}

	ws.SetPort(port)
	ws.handlers = make(map[string]func(req *Request))
	ws.AddHandler("/", ws.rootHandle)
}

// GetVersion retorna a versão do servidor web
func (ws *MyWebServer) GetVersion() string {
	return ws.version
}

// SetPort define a porta que será utilizada no servidor web
func (ws *MyWebServer) SetPort(port int) {
	ws.port = port
}

// GetPort retorna a porta definida
func (ws *MyWebServer) GetPort() int {
	return ws.port
}

// AddHandler adiciona a chamada na lista de chamadas possíveis
func (ws *MyWebServer) AddHandler(name string, fun func(req *Request)) {
	fmt.Println("AddHandler", name)

	for len(name) > 0 && name[0:1] == "/" {
		name = name[1:]
	}

	fmt.Println("AddHandler", name)
	name = "/" + name
	ws.handlers[name] = fun
	http.HandleFunc(name, ws.listener)
}

// Start inicia o servidor
func (ws *MyWebServer) Start() {
	var port = ":" + strconv.Itoa(ws.GetPort())
	fmt.Println("Starting MyWebServer with port", port, "...")
	http.ListenAndServe(port, nil)
}

func (ws *MyWebServer) rootHandle(req *Request) {
	fmt.Fprintln(req.GetResponseWriter(), "MyWebServer is running at", ws.GetPort())
	req.Done()
}

// Listener recebe as chamadas e distribui de acordo com o nome
func (ws *MyWebServer) listener(hrw http.ResponseWriter, hreq *http.Request) {
	fmt.Println("URI: ", hreq.RequestURI)
	uri := hreq.RequestURI
	pos := strings.Index(uri, "?")

	if pos == 0 {
		http.Error(hrw, "MyWebServer: No Service specified.", http.StatusNotImplemented)
		return
	}

	if pos < 0 {
		pos = len(uri)
	}

	name := hreq.RequestURI[0:pos]
	fmt.Println("Name: ", name)
	fun, found := ws.handlers[name]

	if !found || fun == nil {
		http.Error(hrw, "MyWebServer: Service not found.", http.StatusNotImplemented)
		return
	}

	var req = Request{}
	req.New(hrw, hreq)
	go fun(&req)
	req.Wait()
}
