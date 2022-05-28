object Control extends App {
  // 每个语句都有一个值
  
  // v的类型是string和int的共同父类Any
  val v = if (true) "1" else 0
  println(v)


  // 没有break,continue
  var v1 = 10
  val w = while (v1 > 0) {
    v1 -= 1 
    println(v1)
  }
  println("while 的value = ", w)

  // for循环嵌套
  println("for for")
  for (i <- 0 to 1; j <- 0 to 1) {
    println(i,j) 
  }

  // for循环定义临时变量
  println("for tmp")
  for (i <- 0 to 2; tmp = i; j <- 0 to tmp) {
    println(i, j) 
  }

  // for if
  println("for if")
  for (i <- 0 to 10 if i % 2 == 0) {
    println(i)
  }
  
  // for yield
  val vv = for (i <- 0 to 10 if i%2 == 0)
  yield {
    i + 1  
  }
  println(vv)
}
