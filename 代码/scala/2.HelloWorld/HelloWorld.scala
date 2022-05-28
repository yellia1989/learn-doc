object HelloWorld extends App {
  println("Hello World")

  var p = "111"
  val p2 = "111"
  println(p, p2)

  p = "222"
  //val类型的变量不能赋值
  //p2 = "222"
  println(p, p2)

  def func1() {
    println("func1") 
  }

  func1()

  var f2 = (s:String) => println(s)
  f2("function")
}
