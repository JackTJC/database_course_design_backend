package main

import (
	"github.com/JackTJC/backend/conf"
	"github.com/JackTJC/backend/dal"
	"github.com/JackTJC/backend/server"
)

func main() {
	conf.InitConf()
	dal.InitDal()
	server.SetReqFormat(server.PB)
	server.StartServer()
}

//grpc 实现
//port:=fmt.Sprintf(":%s",conf.Conf.Server.Port)
//lis,err:=net.Listen("tcp",port)
//if err!=nil{
//	quick_log.Fatalf("listen error:%v\n", err)
//}
//s:=grpc.NewServer()
//pb_gen.RegisterGoodsInfoAdminServer(s,&GoodsInfoAdminServer{})
//if err:=s.Serve(lis);err!=nil{
//	quick_log.Fatalf("server serve error:%v",err)
//}
