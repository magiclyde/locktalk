/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2023/7/13 下午3:09
 * @note: 基于 ZooKeeper 实现分布式阻塞锁

基于 ZooKeeper 的锁与基于 Redis 的锁的不同之处在于 Lock 成功之前会一直阻塞，这与我们单机场景中的 mutex.Lock 很相似。

其原理也是基于临时 Sequence 节点和 watch API，例如我们这里使用的是 /lock 节点。
Lock 会在该节点下的节点列表中插入自己的值，只要节点下的子节点发生变化，就会通知所有 watch 该节点的程序。
这时候程序会检查当前节点下最小的子节点的 id 是否与自己的一致。如果一致，说明加锁成功了。

这种分布式的阻塞锁比较适合分布式任务调度场景，但不适合高频次持锁时间短的抢锁场景。
按照 Google 的 Chubby 论文里的阐述，基于强一致协议的锁适用于 粗粒度 的加锁操作。
这里的粗粒度指锁占用时间较长。我们在使用时也应思考在自己的业务场景中使用是否合适。
*/

package main

import (
	"time"

	"github.com/go-zookeeper/zk"
)

func main() {
	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second) //*10)
	if err != nil {
		panic(err)
	}
	l := zk.NewLock(c, "/lock", zk.WorldACL(zk.PermAll))
	err = l.Lock()
	if err != nil {
		panic(err)
	}
	println("lock succ, do your business logic")

	// do sth.
	time.Sleep(time.Second * 10)

	l.Unlock()
	println("unlock succ, finish business logic")
}
