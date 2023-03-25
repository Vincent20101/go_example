~~~go
从1.18版本开始，Go添加了对泛型的支持，即类型参数

1、声明一个泛型函数
package main
 
import "fmt"
 
func printslice[T any](s []T) { //[T any]表示支持任何类型的参数  （s []T表示形参s是一个T类型的切片）
	for _, v := range s {
		fmt.Printf("%v\n", v)
	}
}
func main() {
	//常规调用
	printslice[int]([]int{66, 77, 88, 99, 100})
	printslice[string]([]string{"dudu", "yiyi", "8号"})
	//省略显示类型
	printslice([]int{88, 99, 100})
}
2、声明一个泛型切片
package main
 
import "fmt"
 
type vestor[T any] []T
 
func printslice[T any](s []T) { //[T any]表示支持任何类型的参数  （s []T表示形参s是一个T类型的切片）
	for _, v := range s {
		fmt.Printf("%v\n", v)
	}
}
func main() {
	v1 := vestor[int]{58, 1881}
	printslice(v1)
	v2 := vestor[string]{"dudu", "yiyi", "8号"}
	printslice(v2)
}
3、声明一个泛型map
package main
 
import "fmt"
 
type M[K string, V any] map[K]V
 
func main() {
	m1 := M[string, int]{"key": 1}
	m1["key"] = 2
 
	m2 := M[string, string]{"key": "value"}
	m2["key"] = "dudu"
	fmt.Println(m1, m2)
}
4、声明一个泛型通道
package main
 
import "fmt"
 
type C[T any] chan T
 
func main() {
	a := make(C[int], 10)
	a <- 1
	fmt.Println(<-a)
	b := make(C[string], 10)
	b <- "dudu"
	fmt.Println(<-b)
}
5、泛型约束
1）使用interface中规定的类型约束泛型函数的参数
package main
 
import "fmt"
 
type NumStr interface {
	Num | Str
}
type Num interface {
	~int | ~int32 | ~uint64
}
type Str interface {
	string
}
 
func add[T NumStr](a, b T) T {
	return a + b
}
 
func main() {
fmt.Println(add(3,4))
fmt.Println(add("dudu","yiyi"))
}
上面的 NumStr，新增了类型列表表达式，它是对类型参数进行约束。

使用 | 表示取并集

如果传入参数不在集合限制范围内，就会报错。

2）使用interface中规定的方法来约束泛型的参数
package main
 
import (
	"fmt"
	"strconv"
)
 
type Price int
 
func (i Price) String() string {
	return strconv.Itoa(int(i))
}
 
type Price2 string
 
func (i Price2) String() string {
	return string(i)
}
 
type ShowPrice interface {
	String() string
	~int | ~string
}
 
func ShowPriceList[T ShowPrice](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return
}
func main() {
	fmt.Println(ShowPriceList([]Price{1, 2}))
	fmt.Println(ShowPriceList([]Price2{"a", "b"}))
}
3)使用interface中规定的方法和类型来双重约束泛型的参数
package main
 
import (
	"fmt"
	"strconv"
)
 
type Price int
 
func (i Price) String() string {
	return strconv.Itoa(int(i))
}
 
type Price2 string
 
func (i Price2) String() string {
	return string(i)
}
 
type ShowPrice interface {
	String() string
	~int | ~string
}
 
func ShowPriceList[T ShowPrice](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return
}
func main() {
	fmt.Println(ShowPriceList([]Price{1, 2}))
	fmt.Println(ShowPriceList([]Price2{"a", "b"}))
}
4）使用泛型自带comparacle约束，判断比较

package main
 
import "fmt"
 
func findFunc[T comparable](a []T, v T) int {
	for i, e := range a {
		if e == v {
			return i
		}
	}
	return -1
}
func main() {
    fmt.Println(findFunc([]int{1, 2, 3, 4, 5, 6}, 5))
    fmt.Println(findFunc([]string{"dudu", "yiyi", "8号"}, "dudu"))
}

~~~