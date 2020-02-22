package main

import (
    "net"
    "fmt"
)

func tcp_listen(addr string) *net.TCPListener {
    _laddr, err := net.ResolveTCPAddr("tcp4", addr)
    if err != nil {
        panic("listen on address:" + addr + ", err:" + err.Error())
    }

    l, err := net.ListenTCP("tcp4", _laddr)
    if err != nil {
        panic("listen on address:" + addr + ", err:" + err.Error())
    }
    fmt.Printf("listen on %s\n", _laddr)
    return l
}
