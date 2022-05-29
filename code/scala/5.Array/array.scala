import scala.collection.mutable._

object MArray extends App {

    // 定长数组
    val v = Array(1, 2)
    v(0) = 3
    println(v.mkString(","))

    // 变成数组
    var v2 = ArrayBuffer(1,2)
    v2 += (3,4)
    println(v2)
}
