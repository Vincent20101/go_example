package main
/**

此课程提供者：微信imax882

+微信imax882
办理会员 课程全部免费看

课程清单：https://leaaiv.cn

全网最全 最专业的 一手课程

成立十周年 会员特惠 速来抢购

**/
import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 8
	var b int32 = 4
	var c int64 = 4
	d := 4
	fmt.Printf("%T %v\n", a, a)
	fmt.Printf("%T %v\n", b, b)
	fmt.Printf("%T %v\n", c, c)
	fmt.Printf("%T %v\n", d, d)
	//查看数据类型
	fmt.Println(reflect.TypeOf(d))
}
