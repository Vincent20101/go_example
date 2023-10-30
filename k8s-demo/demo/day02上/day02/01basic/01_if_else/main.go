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
	//第一种,变量在if条件判断之外,变量作用域在整个main函数中
	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if (score > 75) {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	//第二种，变量在if条件判断中，变量作用域在if函数中
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if (score > 75) {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

}
