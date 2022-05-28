package main

import (
    "net"
    "fmt"
    "os/exec"
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

func run_shell(file string) string {
    cmd := exec.Command("/bin/bash", "-c", file)
    output, err := cmd.Output()
    if err != nil {
        if err,ok := err.(*exec.ExitError); ok {
            if err.ExitCode() != 1 {
                panic("run " + file + " err:" + err.Error())
            }
        } else {
            panic("run " + file + " err:" + err.Error())
        }
    }
    return string(output)
}
