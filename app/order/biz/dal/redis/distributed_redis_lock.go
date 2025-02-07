package redis

import (
	"context"
	"errors"
	"google.golang.org/appengine/log"
	"time"

	Redis "github.com/redis/go-redis/v9"
)

const (
	// 解锁lua
	unLockScript = "if redis.call('get', KEYS[1]) == ARGV[1] " +
		"then redis.call('del', KEYS[1]) return 1 " +
		"else " +
		"return 0 " +
		"end"

	// 看门狗lua
	watchLogScript = "if redis.call('get', KEYS[1]) == ARGV[1] " +
		"then return redis.call('expire', KEYS[1], ARGV[2]) " +
		"else " +
		"return 0 " +
		"end"
)

type DistributeRedisLock struct {
	redis      *Redis.Client
	ctx        context.Context
	cancelFunc context.CancelFunc
	key        string
	value      string
	expireTime time.Duration
	status     bool
	waitTime   time.Duration
}

func (d *DistributeRedisLock) TryLock() (bool, error) {
	if err := d.Lock(); err != nil {
		return false, err
	}

	d.status = true
	go d.Watchdog()

	return true, nil
}

func (d *DistributeRedisLock) Lock() error {
	now := time.Now()

	for time.Since(now) < d.waitTime {
		isLock, err := d.redis.SetNX(d.ctx, d.key, d.value, d.expireTime).Result()

		if err != nil {
			return err
		}

		if !isLock {
			time.Sleep(1000 * time.Millisecond)
		} else {
			return nil
		}
	}
	return errors.New("lock timeout")
}

func (d *DistributeRedisLock) Watchdog() {
	// 创建一个定时器NewTicker, 每过期时间的2分之1触发一次
	loopTime := time.Duration(d.expireTime*1000*1/2) * time.Millisecond
	expTicker := time.NewTicker(loopTime)
	for {
		select {
		case <-d.ctx.Done():
			return
		case <-expTicker.C:
			if d.status {
				args := []interface{}{d.key, d.value}
				res, err := d.redis.EvalSha(d.ctx, watchLogScript, []string{d.key}, args...).Result()

				if err != nil {
					log.Debugf(d.ctx, "redis eval error: %v", err)
					return
				}

				r, ok := res.(int64)
				if !ok || r == 0 {
					log.Debugf(d.ctx, "redis eval error: %v", res)
					return
				}
			}
		}
	}
}

func (d *DistributeRedisLock) Unlock() (bool, error) {
	d.cancelFunc()

	if d.status {
		err := d.redis.Eval(context.Background(), unLockScript, []string{d.key}, []string{d.value}).Err()
		if err != nil {
			return false, err
		}

		d.status = false
		return true, nil
	}

	return false, errors.New("unlock error")
}

func NewDistributeRedisLock(key string, expireTime time.Duration, value string, waitTime ...time.Duration) *DistributeRedisLock {
	wait := time.Second * 3
	if len(waitTime) > 0 {
		wait = waitTime[0]
	}

	ctx, cancelFunc := context.WithCancel(context.Background())

	return &DistributeRedisLock{
		redis:      DB(),
		key:        key,
		value:      value,
		expireTime: expireTime,
		waitTime:   wait,
		ctx:        ctx,
		cancelFunc: cancelFunc,
	}
}
