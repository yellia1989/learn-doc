// 本例子的作用是测试t.Run方法和t.Parallel方法
// 结论：
// 1. 在一个test case中调用t.Run方法, 会产生一个新的协程来跑这个test case
// 并且父携程会阻塞在t.Run方法中直到新的携程退出
// 2. 在一个test case中调用t.Run方法, 会产生一个新的协程来跑这个test case, 
// 如果不希望父携程阻塞新的携程执行结束,可以在新的携程开头执行t.Parallel方法
// 这样子携程会在父携程执行完后测试用例后开始执行，但是父携程会等待子携程退出

package test

import (
    "testing"
)

func TestRun1(t *testing.T) {

    t.Log("开启子测试用例,等待子测试用例结束")

    t.Run("child1", func (t *testing.T) {
        t.Log("我是子测试用例")
    })

    t.Log("子测试用例结束，测试用例继续执行")
}

func TestRun2(t *testing.T) {

    t.Log("开启子测试用例,不等待子测试用例结束, 测试继续执行")

    t.Run("child2", func (t *testing.T) {
        t.Parallel()
        t.Log("我是子测试用例,现在开始执行")
    })

    t.Log("继续执行，直到完成")
}
