package main

import (
    "os"
    _ "syscall"
    "fmt"
    "os/signal"
)

func main() {
    c := make(chan os.Signal, 1)
    signal.Notify(c)

    for {
        s := <-c
        fmt.Printf("recv signal:%s\n", s)
    }
}
