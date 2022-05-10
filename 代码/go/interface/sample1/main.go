package main

import (
    "fmt"
)

type Go struct {}
func (g Go) sayHello() {
    fmt.Println("Hi, I am GO!")
}

type PHP struct {}
func (p PHP) sayHello() {
    fmt.Println("Hi, I am PHP!")
}

type C struct {}

type IGreeting interface {
    sayHello()
}

func sayHello(i IGreeting) {
    i.sayHello()
}

func main() {
    golang := Go{}
    php := PHP{}
    c := C{}

    sayHello(golang)
    sayHello(php)

    sayHello(c)
}
