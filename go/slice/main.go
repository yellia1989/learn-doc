package main

import (
    "fmt"
)

func main() {
    s1 := []byte("hello")
    fmt.Printf("s:%s len=%d\n", string(s1), len(s1))

    //s2 := s1[:len(s1)+1]
    //fmt.Printf("s:%s len=%d\n", string(s2), len(s2))

    s3 := s1[len(s1):]
    fmt.Printf("s:%s len=%d\n", string(s3), len(s3))
}
