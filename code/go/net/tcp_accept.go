// 测试tcp的accept的返回值
// 设置了Deadline后返回net.error timeout=true temporary=true

package main

import (
    "fmt"
    "time"
    "net"
)

func main() {
    // 监听
    l := tcp_listen(":8888")
    defer l.Close()

    for {
        l.SetDeadline(time.Now().Add(time.Second * 3))
        conn, err := l.AcceptTCP()
        if err != nil {
            if err, ok := err.(net.Error); ok {
                fmt.Printf("%s timeout:%t temporary:%t\n", time.Now(), err.Timeout(), err.Temporary())
                continue
            } else {
                panic("accept err:" + err.Error())
            }
        }
        fmt.Printf("%s tcp conn:%v\n", time.Now(), conn)
    }
}
