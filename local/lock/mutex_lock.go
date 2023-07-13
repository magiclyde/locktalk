/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2023/7/13 下午1:45
 * @note: 加锁以创造临界区,并发地修改全局变量
 */

package main

import (
	"sync"
)

// 全局变量
var counter int
var l sync.Mutex

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			l.Lock()
			counter++
			l.Unlock()
		}()
	}

	wg.Wait()
	println(counter)
}
