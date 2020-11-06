package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// Server is a mock conduit server.
type Server struct {
	server  *httptest.Server
	handler *handler
}

// New creates a new empty conduit server.
func New() *Server {
	h := newHandler()
	ts := httptest.NewServer(h)

	return &Server{
		server:  ts,
		handler: h,
	}
}

type handlerResponse struct {
	HTTPCode int
	Payload  map[string]interface{}
}

type handler struct {
	routes map[string]handlerResponse
}

func newHandler() *handler {
	return &handler{
		routes: make(map[string]handlerResponse),
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Serve only POST requests.
	if req.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response, ok := h.routes[req.RequestURI]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := json.Marshal(response.Payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.HTTPCode)
	w.Write(data)
}

func (h *handler) RegisterMethod(
	method string,
	httpCode int,
	response map[string]interface{},
) {
	h.routes[method] = handlerResponse{
		HTTPCode: httpCode,
		Payload:  response,
	}

}

// RegisterCapabilities adds a default handler for the
// `conduit.getcapabilities` API endpoint.
func (s *Server) RegisterCapabilities() {
	s.RegisterMethod("conduit.getcapabilities", 200, gin.H{
		"result": map[string][]string{
			"authentication": {"token", "session"},
			"signatures":     {"consign"},
			"input":          {"json", "urlencoded"},
			"output":         {"json"},
		},
	})
}

// RegisterMethod adds a handler for a specific conduit API method.
func (s *Server) RegisterMethod(
	method string,
	httpCode int,
	response map[string]interface{},
) {
	s.handler.RegisterMethod(fmt.Sprintf("/api/%s", method), httpCode, response)
}

// GetURL returns the URL of the root of the server.
func (s *Server) GetURL() string {
	return s.server.URL
}

// Close shuts down the server. This should be called at the end of every test
// or by using defer.
func (s *Server) Close() {
	s.server.Close()
}
