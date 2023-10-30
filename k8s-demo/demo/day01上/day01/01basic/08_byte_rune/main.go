package main
/**

此课程提供者：微信imax882

+微信imax882
办理会员 课程全部免费看

课程清单：https://leaaiv.cn

全网最全 最专业的 一手课程

成立十周年 会员特惠 速来抢购

**/
import "fmt"

func main() {
	strs := "我是谁abc"
	a := []rune(strs)
	b := []byte(strs)
	fmt.Printf("%T %d\n", a, a)
	fmt.Printf("%T %d\n", b, b)

	//单引号
	var a1 = 'a'
	fmt.Printf("%T %d", a1, a1)
}
