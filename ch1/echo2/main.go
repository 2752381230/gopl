// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

//!-
/*
var s, sep string
    // means var s string; var sep string
s := ""
    // 只能在函数内使用，不能直接在包内使用
var s = ""
    // 没有指明类型，此时能推导出来类型string
    // 等价于 var s string
    // var x = 1  此时只能取默认 int 类型
var s string = ""
    // "标准"写法，但“less is more”

推荐第一第二种写法，注意第二种在函数内
*/
