package test

import (
	"context"
	"fmt"
	"github.com/JackTJC/backend/conf"
	"github.com/JackTJC/backend/dal"
	"testing"
)

func TestCache(t *testing.T) {
	conf.InitConf()
	tel:="17683969426"
	limit:=int32(1)
	dal.InitDal()
	data, err := dal.QueryClientWithCache(context.Background(), &tel, nil, nil, &limit)
	if err!=nil{
		panic(err)
	}
	fmt.Println(data)
}
