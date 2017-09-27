// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.
//!+

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"    // bufio 包实现了带缓存的I/O操作. io.Reader io.Writer
	"fmt"
	"os"
)

func main() {
    fmt.Println("ctrl+d to stop the input")
    fmt.Println("\ninput 'a' != 'a '\n")
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

//!-
/*
map[key]val
    key: 任意类型，其值可用 == 运算符的就可以
    val: 任意类型

make 函数创建的为空map
range map 时 key 无序
*/
