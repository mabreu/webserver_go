package mywebserver

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// MyWebServer é a estrutura de dados básica do nosso web server
type MyWebServer struct {
	handlers map[string]func(req *MyRequest)
	Port     int
}

// New inicializa a estrutura MyWebServer
func (ws *MyWebServer) New(port int) {
	if port == 0 {
		port = 8080
	}

	ws.Port = port
	ws.handlers = make(map[string]func(req *MyRequest))
	ws.AddHandler("/", ws.rootHandle)
}

// AddHandler adiciona um novo serviço ao nosso Servidor
func (ws *MyWebServer) AddHandler(name string, fun func(req *MyRequest)) {
	fmt.Println("AddHandler: ", name)

	for len(name) > 0 && name[0:1] == "/" {
		name = name[1:]
	}

	name = "/" + name
	ws.handlers[name] = fun
	http.HandleFunc(name, ws.listener)
}

// Start inicia o servidor web MyWebServer
func (ws *MyWebServer) Start() {
	fmt.Println("Starting MyWebServer with port", ws.Port, "...")
	http.ListenAndServe(":"+strconv.Itoa(ws.Port), nil)
}

func (ws *MyWebServer) rootHandle(req *MyRequest) {
	fmt.Fprintln(req.HTTPResponse, "MyWebServer is running at", ws.Port)
	req.Done()
}

func (ws *MyWebServer) listener(hrw http.ResponseWriter, hreq *http.Request) {
	fmt.Println("URI:", hreq.RequestURI)
	uri := hreq.RequestURI
	pos := strings.Index(uri, "?")

	if pos == 0 {
		http.Error(hrw, "MyWebServer: No service Specified.", http.StatusNotImplemented)
		return
	}

	if pos < 0 {
		pos = len(uri)
	}

	name := hreq.RequestURI[0:pos]
	fun, found := ws.handlers[name]

	if !found || fun == nil {
		http.Error(hrw, "MyWebServer: Service not found.", http.StatusNotImplemented)
		return
	}

	var myreq = MyRequest{}
	myreq.New(hrw, hreq)
	go fun(&myreq)
	myreq.Wait()
}
