package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

type SumFunc func(int64, int64) int64

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func timedSumFunc(f SumFunc) SumFunc {
	return func(start int64, end int64) int64 {
		defer func(t time.Time) {
			fmt.Printf("--- Time Elapsed (%s): %v ---\n", getFunctionName(f), time.Since(t))
		}(time.Now())
		return f(start, end)
	}
}

func Sum1(start, end int64) int64 {
	var sum int64
	sum = 0
	if start > end {
		start, end = end, start
	}
	for i := start; i <= end; i++ {
		sum += i
	}
	return sum
}

func Sum2(start, end int64) int64 {
	if start > end {
		start, end = end, start
	}
	return (end - start + 1) * (end + start) / 2
}

//reference: https://coolshell.cn/articles/21228.html
func main() {
	sum1 := timedSumFunc(Sum1)
	sum2 := timedSumFunc(Sum2)
	fmt.Printf("%d, %d\n", sum1(-1000, 10000000), sum2(-1000, 10000000))
}