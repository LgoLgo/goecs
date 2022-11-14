package main

import (
	"context"
	"fmt"
	"testing"

	"srvs/goods_srv/proto/gen"
)

func TestGetCategoryBrandList(t *testing.T) {
	Init()
	rsp, err := brandClient.CategoryBrandList(context.Background(), &proto.CategoryBrandFilterRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	fmt.Println(rsp.Data)
}
