# spike
12306 秒杀

## 运行代码
```shell
go run main.go
```

## 压力测试
```shell
ab -s 120 -n 10000 -c 100 http://127.0.0.1:3005/buy/ticket # 120秒内发起10000个请求，100个并发
```