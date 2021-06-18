package test

import (
	"fmt"
	"github.com/JackTJC/backend/conf"
	"github.com/JackTJC/backend/dal"
	"testing"
)

func TestDB(t *testing.T) {
	conf.InitConf()
	dal.InitDal()
	tel:="17683969426"
	limit:=int32(1)
	client, err := dal.QueryClient(&tel, nil, nil, &limit)
	if err!=nil{
		panic(err)
	}
	fmt.Println(client)
}
