// exer 1.1 1.2 1.3 prints its command-line arguments from 0.
package main

import (
    "fmt"
    "os"
    "strings"
    "path"
    "time"
)

func main() {
    // 1.1
    fmt.Println(strings.Join(os.Args[0:], " "))
    new_Args := make([]string, 0)
    new_Args  = append(new_Args, path.Base(os.Args[0]))
    new_Args  = append(new_Args, os.Args[1:]...)
    fmt.Println(strings.Join(new_Args[:], " "))
    // 1.2
    for index, val := range os.Args {
        fmt.Println(index, val)
    }
    // 1.3
    fmt.Printf("\n+= format: ")
    start := time.Now()
    var s, sep string
    for i := 0; i < 10000; i++ {
        s += sep + fmt.Sprintf("%d", i)
        sep = " "
    }
    end := time.Now()
    fmt.Println(end.Sub(start).Seconds())
    fmt.Printf("\nJoin format: ")
    start = time.Now()
    arr := make([]string, 0, 10000)
    for i := 0; i < 10000; i++ {
        arr = append(arr, fmt.Sprintf("%d", i))
    }
    s = strings.Join(arr, " ")
    end = time.Now()
    fmt.Println(end.Sub(start).Seconds())
}

//!-
