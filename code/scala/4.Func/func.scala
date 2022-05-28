object Func extends App {

  // 只有返回类型为空时可以不写返回值
  def fun1(a: Int*) {
    println(a)
  }

  fun1(1,2,3)

  // 默认参数和命名参数
  def fun2(name: String, sex: Int = 1) = {
    name + "" + sex 
  }
  println(fun2("yellia"), fun2(sex=0, name="1"))
}
