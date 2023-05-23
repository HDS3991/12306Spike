package main

import (
	"context"
	"net/http"
	"spike/global"
	"spike/init/redis"
	"spike/log"
	"spike/logic"
	"spike/util"
	"strconv"
)

var (
	localSpike  logic.LocalSpike
	remoteSpike logic.RemoteSpikeKeys
	done        chan int
)

func init() {
	localSpike = logic.LocalSpike{
		LocalInStock:     150,
		LocalSalesVolume: 0,
	}
	remoteSpike = logic.RemoteSpikeKeys{
		SpikeOrderHashKey:  "ticket_hash_key",
		TotalInventoryKey:  "ticket_total_nums",
		QuantityOfOrderKey: "ticket_sold_nums",
	}
	global.Redis = redis.NewPool()
	done = make(chan int, 1)
	done <- 1
}

func main() {
	http.HandleFunc("/buy/ticket", handleReq)
	http.ListenAndServe(":3005", nil)
}

func handleReq(w http.ResponseWriter, r *http.Request) {
	redisConn := global.Redis
	msg := ""
	<-done
	//全局读写锁
	if localSpike.LocalDeductionStock() && remoteSpike.RemoteDeductionStock(context.TODO(), redisConn) {
		util.RespJson(w, 1, "抢票成功", nil)
		msg = msg + "result:1,localSales:" + strconv.FormatInt(localSpike.LocalSalesVolume, 10)
	} else {
		util.RespJson(w, -1, "已售罄", nil)
		msg = msg + "result:0,localSales:" + strconv.FormatInt(localSpike.LocalSalesVolume, 10)
	}

	// 将抢票状态写入到log中
	done <- 1
	log.Write(msg, "./stat.log")
}
