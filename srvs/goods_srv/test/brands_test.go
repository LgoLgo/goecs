package main

import (
	"context"
	"fmt"
	"testing"

	"google.golang.org/grpc"

	"srvs/goods_srv/proto/gen"
)

var brandClient proto.GoodsClient
var conn *grpc.ClientConn

func TestGetBrandList2(t *testing.T) {
	Init()
	rsp, err := brandClient.BrandList(context.Background(), &proto.BrandFilterRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	for _, brand := range rsp.Data {
		fmt.Println(brand.Name)
	}
}
