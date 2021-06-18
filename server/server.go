package server

import (
	"fmt"
	"github.com/JackTJC/backend/conf"
	"github.com/gin-gonic/gin"
	"log"
)

type ReqFormat int32
const (
	JSON ReqFormat=1
	PB ReqFormat=2
)
var reqFormat ReqFormat

func StartServer(){
	port:=fmt.Sprintf(":%s",conf.Conf.Server.Port)
	c:=gin.Default()
	setRoute(c)
	err := c.Run(port)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func SetReqFormat(format ReqFormat) {
	reqFormat=format
}
