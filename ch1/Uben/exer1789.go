// exec 1.7 1.8 1.9 for fetch
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	// 1.7
    //"io/ioutil"
	"io"
	"net/http"
	"os"
    // 1.8
    "strings"
)

func main() {
	for _, url := range os.Args[1:] {
        // 1.8
        if !strings.HasPrefix(url, "http://") {
            url = "http://" + url
        }
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
        // 1.9
		fmt.Printf("%s", resp.Status)
        // 1.7
		//b, err := ioutil.ReadAll(resp.Body)
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		//fmt.Printf("%s", b)
	}
}

//!-
