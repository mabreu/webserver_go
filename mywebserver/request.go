package mywebserver

import (
	"net/http"
	"sync"
)

// Request é o "objeto" que será passado para os métodos
type Request struct {
	wg   sync.WaitGroup
	hrw  http.ResponseWriter
	hreq *http.Request
}

// New cria um novo Request
func (req *Request) New(hrw http.ResponseWriter, hreq *http.Request) {
	req.hrw = hrw
	req.hreq = hreq
	req.wg.Add(1)
}

// Done encerra a execução de um processo de requisição
func (req *Request) Done() {
	req.wg.Done()
}

// Wait aguarda até a execução do processo terminar
func (req *Request) Wait() {
	req.wg.Wait()
}

// GetResponseWriter retorna o ResponseWriter da requisição
func (req *Request) GetResponseWriter() http.ResponseWriter {
	return req.hrw
}

// SetResponseWriter define o ResponseWriter da requisição
func (req *Request) SetResponseWriter(hrw http.ResponseWriter) {
	req.hrw = hrw
}

// GetRequest retorna o HTTPRequest da requisição
func (req *Request) GetRequest() *http.Request {
	return req.hreq
}

// SetRequest define o HTTPRequest da requisição
func (req *Request) SetRequest(hr *http.Request) {
	req.hreq = hr
}
