package db

import (
	"fmt"
	"github.com/JackTJC/backend/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const dialect = "mysql"
var(
	args string
	dbConn *Connect
)

type Connect struct {
	*gorm.DB
}

func Init(){
	args=fmt.Sprintf("%s:%s@(%s:%s)/%s",
		conf.Conf.Mysql.UserName,
		conf.Conf.Mysql.Passwd,
		conf.Conf.Mysql.Host,
		conf.Conf.Mysql.Port,
		conf.Conf.Mysql.Database,
		)
	db,err:=gorm.Open(dialect,args)
	if err!=nil{
		panic(err)
	}
	dbConn=&Connect{db}
}
