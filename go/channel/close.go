// 本例子的主要作用是验证往已经关闭的close读写的情况
// 结论：
// 1. 往已经关闭的channl写会panic
// 2. 读已关闭的channel会返回ok == false, 可以借此判断channel关闭, 可以多次读

package main

import (
    "fmt"
)

func main() {

    // 无缓冲写
    // panic: send on closed channel
    /*ch := make(chan int)
    close(ch)
    ch <- 1
    */

    // 有缓存写
    // panic: send on closed channel
    /*ch := make(chan int, 1)
    close(ch)
    ch <- 1
    */

    // 无缓冲读
    // 不会panic可以多次读
    /*ch := make(chan string)
    close(ch)
    s,ok := <-ch
    fmt.Printf("%s,%b\n", s, ok)
    s2,ok := <-ch
    fmt.Printf("%s,%b\n", s2, ok)
    */

    ch := make(chan []byte, 1)
    close(ch)
    s,ok := <-ch
    fmt.Printf("%v,%b,%b\n", s, ok, s == nil)
    s2,ok := <-ch
    fmt.Printf("%v,%b\n", s2, ok)
}
