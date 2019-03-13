// date: 2019-03-13
package cache

import "log"

func New(typ string) Cache {
	var c Cache
	if typ == "inMemory" {
		c = newInMemoryCache()
	}

	if c == nil {
		panic("未知缓存类型: " + typ)
	}
	log.Println(typ, "缓存对象实例化成功")
	return c
}
