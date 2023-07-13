/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2023/7/13 下午1:31
 * @note: 不加锁并发地修改全局变量
 */

package main

import (
	"sync"
)

// 全局变量
var counter int

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}

	wg.Wait()
	println(counter)
}
