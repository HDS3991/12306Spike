package logic

import (
	"context"
	"github.com/redis/go-redis/v9"
)

// LuaScript
// tonumber redis 内部函数，将字符串转换为数字，如果转换失败，返回 nil
const LuaScript = `
	local ticket_key = KEYS[1]
	local ticket_total_key = ARGV[1]
	local ticket_sold_key = ARGV[2]
	local ticket_total_nums = tonumber(redis.call('HGET', ticket_key, ticket_total_key))
	local ticket_sold_nums = tonumber(redis.call('HGET', ticket_key, ticket_sold_key))
	-- 查看是否还有余票,增加订单数量,返回结果值
	if(ticket_total_nums >= ticket_sold_nums) then
		return redis.call('HINCRBY', ticket_key, ticket_sold_key, 1)
	end
	return 0
`

type RemoteSpikeKeys struct {
	SpikeOrderHashKey  string //redis中秒杀订单hash结构key
	TotalInventoryKey  string //hash结构中总订单库存key
	QuantityOfOrderKey string //hash结构中已有订单数量key
}

func (RemoteSpikeKeys *RemoteSpikeKeys) RemoteDeductionStock(ctx context.Context, conn *redis.ClusterClient) bool {
	lua := redis.NewScript(LuaScript)
	keys := []string{RemoteSpikeKeys.SpikeOrderHashKey, RemoteSpikeKeys.TotalInventoryKey, RemoteSpikeKeys.QuantityOfOrderKey}
	result, err := lua.Run(ctx, conn, keys).Int()
	if err != nil {
		return false
	}
	return result != 0
}
