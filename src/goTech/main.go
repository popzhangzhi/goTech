package main

import (
	"fmt"
	"goTech/compile"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {

	//asmInt()
	asmString()
}

//go汇编使用
func asmInt() {
	fmt.Println(compile.Id)

}

//go 汇编字符串声明
func asmString() {
	//fmt.Println(compile.NameData)
	//compile.NameData[0] = 'a'
	fmt.Println(compile.Name)

}

/*
	匹配所有的符合要求的正则
	return [][]string
	第一维 标识个数，第二维匹配的项，以及（）内的string
	ex.
	[[r812 | lihao | 2018-10-19 18:46:59 812 lihao 2018-10-19 18:46:59] [r810 | lihao | 2018-10-19 11:18:13 810 lihao 2018-10-19 11:18:13]]
*/
func MactchGrepString(pattam string, data string) [][]string {
	r, _ := regexp.Compile(pattam)
	rel := r.FindAllStringSubmatch(data, -1)

	return rel
}

func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}
