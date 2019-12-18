package base

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"xorm.io/xorm"
)

type Model struct {
	Id       int64     `xorm:"pk autoincr comment('唯一ID')" json:"id"`
	CreateAt time.Time `xorm:"created comment('创建时间')"`
	UpdateAt time.Time `xorm:"updated comment('更新时间')"`
	DeleteAt time.Time `xorm:"deleted comment('删除时间')"`
}

var Engine *xorm.Engine

func init() {
	var err error
	Engine, err = xorm.NewEngine("mysql", "saler:saler@(192.168.0.36)/base?charset=utf8mb4")
	if err != nil {
		log.Fatal("数据库连接失败：", err)
	}
	Engine.ShowSQL(true)

}
