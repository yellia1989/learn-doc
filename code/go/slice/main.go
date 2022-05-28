package main

import (
    "fmt"
)

func out(p *[]int) {
    *p = make([]int,2)
    p1 := *p
    fmt.Printf("p1:%p\n", p1)

    p1[0] = 10
}

func main() {
    s1 := []byte("hello")
    fmt.Printf("s:%s len=%d\n", string(s1), len(s1))

    //s2 := s1[:len(s1)+1]
    //fmt.Printf("s:%s len=%d\n", string(s2), len(s2))

    s3 := s1[len(s1):]
    fmt.Printf("s:%s len=%d\n", string(s3), len(s3))

    var p []int
    out(&p)
    fmt.Printf("p:%p, %v\n", p, p)
}
