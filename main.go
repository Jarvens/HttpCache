// date: 2019-03-13
package main

import (
	"github.com/Jarvens/HttpCache/cache"
	"github.com/Jarvens/HttpCache/http"
)

func main() {
	c := cache.New("inMemory")
	http.New(c).Listen()
}
