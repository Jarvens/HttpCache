// date: 2019-03-13
package http

import (
	"github.com/Jarvens/HttpCache/cache"
	"net/http"
)

//定义服务结构体绑定事件
type Server struct {
	cache.Cache
}

func (s *Server) Listen() {
	http.Handle("/cache/", s.cacheHandler())
	http.Handle("/status", s.statusHandler())
	http.ListenAndServe(":1234", nil)
}

func New(c cache.Cache) *Server {
	return &Server{c}
}
