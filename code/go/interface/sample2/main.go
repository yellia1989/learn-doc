// 如果实现了接收者是值类型的方法，会隐含地也实现了接收者是指针类型的方法。

package main

import "io"

type myWriter struct {

}

func (w myWriter) Write(p []byte) (n int, err error) {
    return
}

func main() {
    // 检查 *myWriter 类型是否实现了 io.Writer 接口
    var _ io.Writer = (*myWriter)(nil)

    // 检查 myWriter 类型是否实现了 io.Writer 接口
    var _ io.Writer = myWriter{}
}
