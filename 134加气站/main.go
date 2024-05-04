package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	var sum int
	for i := range gas {
		gas[i] = gas[i] - cost[i]
		sum += gas[i]
	}
	if sum < 0 {
		return -1
	}
	var index, curSum int
	index = -1
	for i := 0; i < len(gas); i++ {
		if i <= index {
			continue
		}
		if gas[i] < 0 {
			// 没有资格作为开始
			index = i
			curSum += gas[i]
			continue
		} else {
			// 有资格开始则尝试一下
			left := gas[i]
			for j := i + 1; j < len(gas); j++ {
				left += gas[j]
				if left < 0 {
					index = j
					curSum += left
					break
				}
			}
			if left >= 0 && left+curSum >= 0 {
				return i
			}
		}
	}
	return -1
}

func canCompleteCircuit1(gas []int, cost []int) int {
	var sum int
	for i := range gas {
		gas[i] = gas[i] - cost[i]
		sum += gas[i]
	}
	if sum < 0 {
		return -1
	}
	var index, curSum int
	index = -1
	for i := 0; i < len(gas); i++ {
		if i <= index {
			continue
		}
		if gas[i] < 0 {
			// 没有资格作为开始
			index = i
			curSum += gas[i]
			continue
		} else {
			// 有资格开始则尝试一下
			left := gas[i]
			for j := i + 1; j < len(gas); j++ {
				left += gas[j]
				if left < 0 {
					index = j
					curSum += left
					break
				}
			}
			if left >= 0 && left+curSum >= 0 {
				return i
			}
		}
	}
	return -1
}

func main() {
	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}
	fmt.Println(canCompleteCircuit(gas, cost))
}

//func canCompleteCircuit1(gas []int, cost []int) int {
//	var sum int
//	for i := range gas {
//		gas[i] = gas[i] - cost[i]
//		sum += gas[i]
//	}
//	if sum < 0 {
//		return -1
//	}
//	var index, lastSum, left int
//	left := math.MaxInt
//	for {
//		if left == math.MaxInt {
//			// 没人开始
//			if gas[] < 0 {
//				// 没有资格作为开始
//				i++
//				lastSum += gas[i]
//				continue
//			}
//		}
//
//
//		i := index
//		if gas[i] < 0 {
//			// 没有资格作为开始
//			i++
//			lastSum += gas[i]
//			continue
//		} else {
//			// 有资格开始则尝试一下
//			left := gas[i]
//			stop := 0
//			for j := i + 1; j < len(gas); j++ {
//				left += gas[j]
//				if left < 0 {
//					stop = j
//					i = j
//					lastSum += left
//					break
//				}
//			}
//			if left >= 0 && left+lastSum >= 0 {
//				return i
//			}
//			if left < 0 {
//				// 没走通
//			}
//		}
//	}
//}
