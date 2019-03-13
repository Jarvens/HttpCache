// date: 2019-03-13
package cache

import "sync"

//定义缓存模型
type inMemoryCache struct {

	//用于存放 key  value 的map
	c map[string][]byte

	//读写锁防止脏数据
	mutex sync.RWMutex

	//缓存状态对象  匿名，相当于该模型直接拥有Stat的属性以及方法
	Stat
}

//设置缓存
func (c *inMemoryCache) Set(k string, v []byte) error {

	//加锁
	c.mutex.Lock()
	defer c.mutex.Unlock()

	//根据key查询map
	tmp, exist := c.c[k]

	//存在，删除该值，因为map如果存在该值不会覆盖
	if exist {
		c.del(k, tmp)
	}
	//重新赋值
	c.c[k] = v
	//更新Stat状态  count keySize valueSize
	c.add(k, v)
	return nil
}

//根据key从缓存读取数据
func (c *inMemoryCache) Get(k string) ([]byte, error) {

	//加锁防止脏数据
	c.mutex.Lock()
	defer c.mutex.Unlock()

	//从map中取出key=k的value
	return c.c[k], nil
}

func (c *inMemoryCache) Del(k string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	v, exist := c.c[k]
	if exist {
		delete(c.c, k)
		c.del(k, v)
	}
	return nil
}
