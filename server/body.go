package server

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
)

var (
	unmarshaler = jsonpb.Unmarshaler{AllowUnknownFields: true}
	marshaler   = jsonpb.Marshaler{EnumsAsInts: true, EmitDefaults: true, OrigName: true}
)

func readBody(ginCtx *gin.Context,reqModel proto.Message)error {
	r := ginCtx.Request
	var bodyData []byte
	var err error
	if bodyData, err = ioutil.ReadAll(r.Body); err != nil {
		return err
	}
	switch reqFormat {
	case JSON:
		err=unmarshaler.Unmarshal(bytes.NewReader(bodyData),reqModel)
	case PB:
		err=proto.Unmarshal(bodyData,reqModel)
	default:
		err=unmarshaler.Unmarshal(bytes.NewReader(bodyData),reqModel)
	}
	if err!=nil{
		return err
	}
	return nil
}
func writeBody(resModel proto.Message) ([]byte,error) {
	switch reqFormat {
	case JSON:
		str,err:=marshaler.MarshalToString(resModel)
		return []byte(str),err
	case PB:
		data, err := proto.Marshal(resModel)
		return data,err
	default:
		str,err:=marshaler.MarshalToString(resModel)
		return []byte(str),err
	}
}
