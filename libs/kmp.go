package libs

import (
	"fmt"
	"time"
)

//originStr 原始串 ，patternStr 匹配串
func Kmp(originStr string, patternStr string) int {
	var nextArr = getNext(patternStr)
	fmt.Println(nextArr)
	var i = 0
	var j = 0
	var count = 0
	for {
		count++
		if i >= len(originStr) || j >= len(patternStr) {
			break
		}
		if j == -1 || originStr[i] == patternStr[j] {
			i++
			j++
		} else {
			j = nextArr[j] //j回退
		}
	}
	fmt.Println("kmp count:", count)
	if j >= len(patternStr) {
		return i - len(patternStr)
	} else {
		return -1
	}
}
func getNext(patternStr string) []int {
	var plen int = len(patternStr)
	var nextArr = make([]int, plen)
	nextArr[0] = -1
	var j int = 0  // 遍历下标
	var k int = -1 // next值
	var count = 0
	for {
		count++
		if j >= plen-1 {
			break
		}
		if k == -1 || patternStr[j] == patternStr[k] { //刚开始或者 j 对应位置的值 == k 对应位置的值
			k++
			j++
			// if patternStr[j] == patternStr[k] { //后面字符也相等，跳过
			// 	nextArr[j] = nextArr[k]
			// } else {
			// 	nextArr[j] = k
			// }
			nextArr[j] = k
		} else {
			k = nextArr[k]
		}
	}
	fmt.Println("next count:", count)
	return nextArr
}

func TestKmp() {
	fmt.Println(time.Now())
	var originStr = "ABCDEFGHIGASBMSSEELLKSNVASDWWPPAKSSKGGHIGHIAGHIAM"
	//var patternStr = "GHIGHIAGHIAM"
	var patternStr = "ASDIUANKSAKDL"
	var points = Kmp(originStr, patternStr)
	if points >= 0 {
		fmt.Println(originStr[points : len(patternStr)+points])
	} else {
		fmt.Println("NO MATCH")
	}
	fmt.Println(time.Now())

}
