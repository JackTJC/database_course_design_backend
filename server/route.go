package server

import (
	"github.com/JackTJC/backend/method"
	"github.com/JackTJC/backend/pb_gen"
	"github.com/gin-gonic/gin"
	"net/http"
)

func setRoute(r *gin.Engine){
	r.Use(corsHandler())
	r.POST("/goods/create/",wrapCreateGoods)
	r.POST("/goods/query/",wrapQueryGoods)
	r.POST("/client/create/",wrapCreateClient)
	r.POST("/client/query/",wrapQueryClient)
	r.POST("/order/create/",wrapCreateOrder)
	r.POST("/order/query/",wrapQueryOrder)
	r.POST("/goods/m_create/",wrapMCreateGoods)
	r.POST("/order/update/status",wrapUpdateOrderStat)
}

func corsHandler()gin.HandlerFunc{
	return func(context *gin.Context) {
		requestMethod := context.Request.Method
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma,token,openid,opentoken")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
		context.Header("Access-Control-Max-Age", "172800")
		context.Header("Access-Control-Allow-Credentials", "false")
		if requestMethod == "OPTIONS" {
			context.JSON(http.StatusOK, `{"message":"success"}`)
		}
		//处理请求
		context.Next()
	}
}

func wrapUpdateOrderStat(c *gin.Context) {
	req:=&pb_gen.UpdateOrderStatusRequest{}
	if err:=readBody(c,req);err!=nil{
		c.Status(500)
		return
	}
	h:=method.NewUpdateOrderStatHandler(c,req)
	h.Run()
	res:=h.Res
	body, err := writeBody(res)
	if err!=nil{
		c.Status(500)
		return
	}
	setBody(c,body)
	afterHandler(c)
}

func wrapCreateClient(c *gin.Context){
	req:=&pb_gen.CreateClientRequest{}
	if err:=readBody(c,req);err!=nil{
		c.Status(500)
		return
	}
	h:=method.NewCreateClientHandler(c,req)
	h.Run()
	res:=h.Res
	body, err := writeBody(res)
	if err!=nil{
		c.Status(500)
		return
	}
	setBody(c,body)
	afterHandler(c)
}

func wrapQueryClient(c *gin.Context){
	req:=&pb_gen.QueryClientRequest{}
	if err:=readBody(c,req);err!=nil{
		c.Status(500)
		return
	}
	h:=method.NewQueryClientHandler(c,req)
	h.Run()
	res:=h.Res
	body, err := writeBody(res)
	if err!=nil{
		c.Status(500)
		return
	}
	setBody(c,body)
	afterHandler(c)
}

func wrapCreateGoods(c *gin.Context){
	req:=&pb_gen.CreateGoodsRequest{}
	if err:=readBody(c,req);err!=nil{
		c.Status(500)
		return
	}
	h:=method.NewCreateGoodsHandler(c,req)
	h.Run()
	res:=h.Res
	body, err := writeBody(res)
	if err!=nil{
		c.Status(500)
		return
	}
	setBody(c,body)
	afterHandler(c)
}

func wrapQueryGoods(c *gin.Context){
	req:=&pb_gen.QueryGoodsRequest{}
	if err:=readBody(c,req);err!=nil{
		c.Status(500)
		return
	}
	h:=method.NewQueryGoodsHandler(c,req)
	h.Run()
	res:=h.Res
	body, err := writeBody(res)
	if err!=nil{
		c.Status(500)
		return
	}
	setBody(c,body)
	afterHandler(c)
}

func wrapCreateOrder(c *gin.Context){
	req:=&pb_gen.CreateOrderRequest{}
	if err:=readBody(c,req);err!=nil{
		c.Status(500)
		return
	}
	h:=method.NewCreateOrderHandler(c,req)
	h.Run()
	res:=h.Res
	body, err := writeBody(res)
	if err!=nil{
		c.Status(500)
		return
	}
	setBody(c,body)
	afterHandler(c)
}


func wrapQueryOrder(c *gin.Context){
	req:=&pb_gen.QueryOrderRequest{}
	if err:=readBody(c,req);err!=nil{
		c.Status(500)
		return
	}
	h:=method.NewQueryOrderHandler(c,req)
	h.Run()
	res:=h.Res
	body, err := writeBody(res)
	if err!=nil {
		c.Status(500)
		return
	}
	setBody(c,body)
	afterHandler(c)
}
func wrapMCreateGoods(c *gin.Context){
	req:=&pb_gen.MCreateGoodsRequest{}
	if err:=readBody(c,req);err!=nil{
		c.Status(500)
		return
	}
	h:=method.NewMCreateGodsHandler(c,req)
	h.Run()
	res:=h.Res
	body,err:=writeBody(res)
	if err!=nil{
		c.Status(500)
		return
	}
	setBody(c,body)
	afterHandler(c)

}
func afterHandler(c *gin.Context) {

}

func setBody(c *gin.Context, body []byte) {
	switch reqFormat {
	case PB:
		c.Data(200,"application/octet-stream",body)
	case JSON:
		c.Data(200,"application/json",body)
	default:
		c.Data(200,"application/json",body)
	}
}