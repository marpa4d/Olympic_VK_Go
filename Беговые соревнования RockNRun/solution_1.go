package main

import "fmt"

func main() {
    var n int
    fmt.Scanf("%d", &n)
    line := make([]int, n)
    for i, _ := range line {
        fmt.Scanf("%d", &line[i])
    }
    max := line[0]
    for _, v := range line {
    if v > max {
    max = v
}
}
    fmt.Println(max)
}