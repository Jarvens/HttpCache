# HttpCache

编译
```bash
$ cd HttpCache
$ go build main.go
#启用缓存服务
$ ./main

#查看缓存状态
$ curl 127.0.0.1:1234/status

#写入缓存
$ curl -v 127.0.0.1:1234/cache/testKey -XPUT -dtestValue

#查看缓存值
$ curl 127.0.0.1:1234/cache/testKey

```


