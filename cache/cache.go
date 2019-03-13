// date: 2019-03-13
package cache

//缓存接口
type Cache interface {

	// set key,value to cache
	Set(string, []byte) error

	// get value via key from cache
	Get(string) ([]byte, error)

	// delete from cache via key
	Del(string) error

	// get cache stat : count , keySize , valueSize
	GetStat() Stat
}
