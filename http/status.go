// date: 2019-03-13
package http

import (
	"encoding/json"
	"log"
	"net/http"
)

type statusHandler struct {
	*Server
}

//实现HTTP 包 ServeHTTP接口
func (h *statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	b, e := json.Marshal(h.GetStat())
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func (s *Server) statusHandler() http.Handler {
	return &statusHandler{s}
}
