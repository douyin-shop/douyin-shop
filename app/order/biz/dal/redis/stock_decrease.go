package redis

import (
	"context"
	"github.com/pkg/errors"
	Redis "github.com/redis/go-redis/v9"
)

const (
	// 扣减库存lua
	decreaseStockScript = "if redis.call('get', KEYS[1]) >= tonumber(ARGV[1]) " +
		"then return redis.call('decrby', KEYS[1], ARGV[1]) " +
		"else " +
		"return -1 " +
		"end"
)

type StockDecrease struct {
	redis      *Redis.Client
	ctx        context.Context
	cancelFunc context.CancelFunc
	key        string
	stock      uint32
	status     bool
}

func (s *StockDecrease) TryDecrease() (bool, error) {
	// 使用Redis客户端执行Lua脚本
	result, err := s.redis.Eval(s.ctx, decreaseStockScript, []string{s.key}, "1").Result()
	if err != nil {
		// 如果执行脚本出错，返回错误
		return false, errors.Wrap(err, "执行扣减库存脚本失败")
	}

	// 根据脚本返回的结果判断扣减是否成功
	if result == int64(-1) {
		// 库存不足，扣减失败
		return false, nil
	}

	// 扣减成功
	s.status = true
	return true, nil
}

func NewStockDecrease(key string, stock uint32) *StockDecrease {
	ctx, cancelFunc := context.WithCancel(context.Background())
	return &StockDecrease{
		redis:      DB(),
		ctx:        ctx,
		cancelFunc: cancelFunc,
		key:        key,
		stock:      stock,
	}
}
