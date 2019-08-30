package test

import (
	"fmt"
	"testing"
)

func mergeSum(intArr []int, channel chan int) {
	len := len(intArr)
	if len <= 1 {
		channel <- intArr[0]
	} else if len <= 2 {
		channel <- intArr[0] + intArr[1]
	} else {
		mid := len / 2

		currentChannel := make(chan int)
		go mergeSum(intArr[:mid], currentChannel)
		go mergeSum(intArr[mid:], currentChannel)

		sum1, sum2 := <-currentChannel, <-currentChannel
		channel <- sum1 + sum2
	}
}

func TestGetSumByMerge(t *testing.T) {
	// LEN 比较大的时候，可能开启了比较多线程，消耗还是挺大的
	LEN := 10000000
	arr := make([]int, LEN + 1)
	for index := 1; index <= LEN; index++ {
		arr[index] = index
	}
	allSumChan := make(chan int)
	go mergeSum(arr, allSumChan)
	fmt.Println(<-allSumChan)
}
