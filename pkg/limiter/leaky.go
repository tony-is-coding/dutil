package limiter

import (
	"sync"
	"time"
)

/**
漏桶算法, 漏桶算法相当于将请求放在了一个限定速度出的 标准队列中, 每固定周期时间放出一个请求, 将网络请求进行平滑分配
超出队列的部分进行
*/


type LeakyBucket struct {
	l         sync.Locker
	size      int64
	limit     int64
	lastLeaky int64
	rate      time.Duration
	reqs      []func(...interface{}) interface{} // 为每个请求过来分配一个唯一的 id, 然后串联这个唯一id 与 请求处理事务

}

//Try 尝试往池中注入一个请求, 等待漏桶调用
func (l *LeakyBucket) Try() error {
	now := int64(time.Now().Nanosecond())
	pieces := (now - l.lastLeaky) / int64(l.rate)
	if (l.limit - (l.size - pieces)) < 1 {
		// 池不足以容纳
		return PoolFullErr
	}
	l.l.Lock()
	l.size = l.size - pieces + 1
	// 添加请求
	l.reqs = append(l.reqs, func(i ...interface{}) interface{} { return nil })
	l.lastLeaky = now
	l.l.Unlock()
	return nil
}

//Serve 以固定速率整型流量处理
func (l *LeakyBucket) Serve() {
	// 按照固定速率处理请求, 平滑流量
	tk := time.NewTicker(l.rate)
	for {
		select {
		case <-tk.C:
			l.l.Lock()
			handler := l.reqs[0]
			l.reqs = l.reqs[1:]
			l.l.Unlock()
			go handler()
		}
	}
}
