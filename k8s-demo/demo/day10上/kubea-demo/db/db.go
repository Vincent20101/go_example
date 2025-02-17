package db
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
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/wonderivan/logger"
	"kubea-demo/config"
	"kubea-demo/model"
)

var (
	isInit bool //是否已经初始化
	GORM   *gorm.DB
	err    error
)

//db的初始化函数
func Init() {
	//判断是否已经初始化
	if isInit {
		return
	}
	//组装数据库连接的数据
	//parseTime是查询结果是否自动解析为时间
	//loc是Mysql的时区设置
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DbUser,
		config.DbPwd,
		config.DbHost,
		config.DbPort,
		config.DbName)
	GORM, err = gorm.Open(config.DbType, dsn)
	//打印sql语句
	GORM.LogMode(config.LogMode)
	//开启连接池
	GORM.DB().SetMaxIdleConns(config.MaxIdleConns)
	GORM.DB().SetMaxOpenConns(config.MaxOpenConns)
	GORM.DB().SetConnMaxLifetime(config.MaxLifeTime)

	isInit = true
	GORM.AutoMigrate(model.Event{}, model.Chart{})
	logger.Info("连接数据库成功")
}

func Close() error {
	logger.Info("关闭数据库连接")
	return GORM.Close()
}