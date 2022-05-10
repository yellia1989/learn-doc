package main

import "fmt"

type Car struct{
    Name string
    Age int
    XList []int
}
func main(){
    a := Car{Name:"buick", Age: 10, XList:make([]int, 10)}
    b := a
    b.Age = 9
    b.XList[1] = 1
    fmt.Printf("address a:%p,%+v\n", &a, a)
    fmt.Printf("address b:%p,%+v\n", &b, b)
    fmt.Println("----------------------")
    fmt.Printf("Car a:%v\n", a)
    fmt.Printf("Car a:%v,%p\n", len(a.XList), &a.XList)
    fmt.Printf("Car b:%v\n", b)
    fmt.Printf("Car b:%v,%p\n", len(b.XList), &b.XList)
}
