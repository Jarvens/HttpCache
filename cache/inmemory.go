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

//从缓存中删除key = k的值
func (c *inMemoryCache) Del(k string) error {
	//加锁
	c.mutex.Lock()
	defer c.mutex.Unlock()

	//从map中取出对应key的value
	v, exist := c.c[k]

	//存在该值则直接调用删除方法，该删除方法是由Stat提供，由于cache模型匿名聚合了Stat所以可以直接调用
	//调用的原因在于更新stat状态信息
	//删除map中key=k的值
	if exist {
		delete(c.c, k)
		c.del(k, v)
	}
	return nil
}

//获取缓存当前状态
func (c *inMemoryCache) GetStat() Stat {
	return c.Stat
}

//初始化，类似于Java的构造函数
func newInMemoryCache() *inMemoryCache {
	return &inMemoryCache{make(map[string][]byte), sync.RWMutex{}, Stat{}}
}
