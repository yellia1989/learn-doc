package main

import (
    "os"
    _ "syscall"
    "fmt"
    "os/signal"
)

func main() {
    c := make(chan os.Signal)
    signal.Notify(c)

    s := <-c
    fmt.Printf("recv signal:%s\n", s)
}
