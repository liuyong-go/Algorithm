package libs

import "fmt"

//originStr 原始串 ，patternStr 匹配串
func Kmp(originStr string, patternStr string) (resultStr string) {
	return
}
func getNext(patternStr string) (nextValue []int) {
	var plen = len(patternStr)
	fmt.Print(plen)
	return
}

func TestKmp() {
	//var originStr = "ABCDEFGHIASBMSSEEASSIIWELLKSNVASDWWPPAKSSKG"
	var patternStr = "BMSSEEASS"
	var nextValue []int = getNext(patternStr)
	fmt.Print(nextValue)

}
