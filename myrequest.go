package mywebserver

import (
	"net/http"
	"sync"
)

// MyRequest é a estrutura de dados que armazena os dados de requisição
type MyRequest struct {
	HTTPRequest  *http.Request
	HTTPResponse http.ResponseWriter
	wg           sync.WaitGroup
}

// New inicializa a estrutura de dados MyRequest
func (req *MyRequest) New(hrw http.ResponseWriter, hreq *http.Request) {
	req.HTTPResponse = hrw
	req.HTTPRequest = hreq
	req.wg.Add(1)
}

// Done encerra a execução da requisição
func (req *MyRequest) Done() {
	req.wg.Done()
}

// Wait espera que a execução do requisição termine
func (req *MyRequest) Wait() {
	req.wg.Wait()
}
