package handler

import (
	"context"
	"srvs/userop_srv/global"
	"srvs/userop_srv/model"
	"srvs/userop_srv/proto/gen"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (*UserOpServer) GetFavList(_ context.Context, req *proto.UserFavRequest) (*proto.UserFavListResponse, error) {
	var rsp proto.UserFavListResponse
	var userFavs []model.UserFav
	var userFavList []*proto.UserFavResponse

	result := global.DB.Where(&model.UserFav{User: req.UserId, Goods: req.GoodsId}).Find(&userFavs)
	rsp.Total = int32(result.RowsAffected)

	for _, userFav := range userFavs {
		userFavList = append(userFavList, &proto.UserFavResponse{
			UserId:  userFav.User,
			GoodsId: userFav.Goods,
		})
	}

	rsp.Data = userFavList

	return &rsp, nil
}

func (*UserOpServer) AddUserFav(_ context.Context, req *proto.UserFavRequest) (*emptypb.Empty, error) {
	var userFav model.UserFav

	userFav.User = req.UserId
	userFav.Goods = req.GoodsId

	global.DB.Save(&userFav)

	return &emptypb.Empty{}, nil
}

func (*UserOpServer) DeleteUserFav(_ context.Context, req *proto.UserFavRequest) (*emptypb.Empty, error) {
	if result := global.DB.Unscoped().Where("goods=? and user=?", req.GoodsId, req.UserId).Delete(&model.UserFav{}); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "Fav doesn't exist")
	}
	return &emptypb.Empty{}, nil
}

func (*UserOpServer) GetUserFavDetail(_ context.Context, req *proto.UserFavRequest) (*emptypb.Empty, error) {
	var userfav model.UserFav
	if result := global.DB.Where("goods=? and user=?", req.GoodsId, req.UserId).Find(&userfav); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "Fav doesn't exist")
	}
	return &emptypb.Empty{}, nil
}
