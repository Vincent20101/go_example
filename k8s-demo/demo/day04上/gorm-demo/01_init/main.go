package main
/**

此课程提供者：微信imax882

+微信imax882
办理会员 课程全部免费看

课程清单：https://leaaiv.cn

全网最全 最专业的 一手课程

成立十周年 会员特惠 速来抢购

**/
//引入orm包时，需要同时引入驱动包
//"gorm.io/driver/mysql"就是gorm，mysql的驱动包
import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//数据库初始化
func main() {
	//定义连接地址
	//parseTime=True&loc=Local
	//parseTime是查询结果是否自动解析为时间
	//id name type  create_time(time.Time)  数据库字段格式datetime
	//loc是Mysql的时区设置
	dsn := "root:123456@tcp(127.0.0.1:3306)/test_db?charset=utf8&parseTime=True&loc=Local"
	//建立连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println(db)
}
