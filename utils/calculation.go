package utils

import (
	"fmt"
	"time"
)

// 计算停车费用
func CalculationMoney(str int) float64 {
	money := str * 2
	return float64(money)
}

// 计算时间
func CalculationTime(time1, time2 string) int {
	timeLayout := "2006-01-02 15:04:05" // 时间的格式
	timeStr1 := time1                   // 第一个时间字符串
	timeStr2 := time2                   //第二个时间字符串
	t1, _ := time.Parse(timeLayout, timeStr1)
	t2, _ := time.Parse(timeLayout, timeStr2)
	duration := t2.Sub(t1)
	fmt.Println(duration)
	return int(duration)
}
