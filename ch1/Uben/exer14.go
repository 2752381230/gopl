// exer 1.4 prints same line file-name
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// add read info from files
func main() {
	counts := make(map[string]int)
	record := make(map[string]string)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3:%v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
			record[line] = filename
		}
	}
	for line, n := range counts {
		//fmt.Printf("%d\t%s in files: %s\n", n, line, record[line])
		if n > 1 {
			fmt.Printf("%d\t%s in files: %s\n", n, line, record[line])
		}
	}
}
/*
重复的行可能不在同一个文件中，可能在同一个文件中，适当修改即可
*/
