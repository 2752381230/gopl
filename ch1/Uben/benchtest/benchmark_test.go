package exer13

import (
    "fmt"
    "strings"
    "testing"
)

func BenchmarkPlus(b *testing.B) {
    for i := 0; i < b.N; i++ {
        PlusFormat()
    }
}
func BenchmarkJoin(b *testing.B) {
    for i := 0; i < b.N; i++ {
        JoinFormat()
    }
}

func PlusFormat() {
    var s, sep string
    for i := 0; i < 10; i++ {
        s += sep + fmt.Sprintf("%d", i)
        sep = " "
    }
}
func JoinFormat(){
    arr := make([]string, 0, 10)
    for i := 0; i < 10; i++ {
        arr = append(arr, fmt.Sprintf("%d", i))
    }
    strings.Join(arr, " ")
}
