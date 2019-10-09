package main

import "fmt"

func main() {
	// 每21万个块，减半
	// 最初奖励50个比特币
	// 用一个循环来判断，累加
	total:=0.0
	blockInterval:=21.0 // 万
	currentReward:=50.0 //当前的奖励

	// 如果当前奖励大于0
	for currentReward > 0 {
		fmt.Println(currentReward)
		// 每一个区间内的奖励总量 = 当前的奖励 * 当前区间块的数量
		amount:=blockInterval*currentReward
		currentReward *= 0.5
		total += amount
	}

	// 比特币总量
	fmt.Println(total,"万")
}