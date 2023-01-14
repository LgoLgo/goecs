package handler

import (
	"context"
	"srvs/goods_srv/global"
	"srvs/goods_srv/model"
	"srvs/goods_srv/proto/gen"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *GoodsServer) CategoryBrandList(ctx context.Context, req *proto.CategoryBrandFilterRequest) (*proto.CategoryBrandListResponse, error) {
	var categoryBrands []model.GoodsCategoryBrand
	categoryBrandListResponse := proto.CategoryBrandListResponse{}

	var total int64
	global.DB.Model(&model.GoodsCategoryBrand{}).Count(&total)
	categoryBrandListResponse.Total = int32(total)

	global.DB.Preload("Category").Preload("Brands").Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Find(&categoryBrands)

	var categoryResponses []*proto.CategoryBrandResponse
	for _, categoryBrand := range categoryBrands {
		categoryResponses = append(categoryResponses, &proto.CategoryBrandResponse{
			Category: &proto.CategoryInfoResponse{
				Id:             categoryBrand.Category.ID,
				Name:           categoryBrand.Category.Name,
				Level:          categoryBrand.Category.Level,
				IsTab:          categoryBrand.Category.IsTab,
				ParentCategory: categoryBrand.Category.ParentCategoryID,
			},
			Brand: &proto.BrandInfoResponse{
				Id:   categoryBrand.Brands.ID,
				Name: categoryBrand.Brands.Name,
				Logo: categoryBrand.Brands.Logo,
			},
		})
	}

	categoryBrandListResponse.Data = categoryResponses
	return &categoryBrandListResponse, nil
}

func (s *GoodsServer) GetCategoryBrandList(ctx context.Context, req *proto.CategoryInfoRequest) (*proto.BrandListResponse, error) {
	brandListResponse := proto.BrandListResponse{}

	var category model.Category
	if result := global.DB.Find(&category, req.Id).First(&category); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Product category does not exist")
	}

	var categoryBrands []model.GoodsCategoryBrand
	if result := global.DB.Preload("Brands").Where(&model.GoodsCategoryBrand{CategoryID: req.Id}).Find(&categoryBrands); result.RowsAffected > 0 {
		brandListResponse.Total = int32(result.RowsAffected)
	}

	var brandInfoResponses []*proto.BrandInfoResponse
	for _, categoryBrand := range categoryBrands {
		brandInfoResponses = append(brandInfoResponses, &proto.BrandInfoResponse{
			Id:   categoryBrand.Brands.ID,
			Name: categoryBrand.Brands.Name,
			Logo: categoryBrand.Brands.Logo,
		})
	}

	brandListResponse.Data = brandInfoResponses

	return &brandListResponse, nil
}

func (s *GoodsServer) CreateCategoryBrand(ctx context.Context, req *proto.CategoryBrandRequest) (*proto.CategoryBrandResponse, error) {
	var category model.Category
	if result := global.DB.First(&category, req.CategoryId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Product category does not exist")
	}

	var brand model.Brands
	if result := global.DB.First(&brand, req.BrandId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "brand does not exist")
	}

	categoryBrand := model.GoodsCategoryBrand{
		CategoryID: req.CategoryId,
		BrandsID:   req.BrandId,
	}

	global.DB.Save(&categoryBrand)
	return &proto.CategoryBrandResponse{Id: categoryBrand.ID}, nil
}

func (s *GoodsServer) DeleteCategoryBrand(ctx context.Context, req *proto.CategoryBrandRequest) (*emptypb.Empty, error) {
	if result := global.DB.Delete(&model.GoodsCategoryBrand{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "Brand category does not exist")
	}
	return &emptypb.Empty{}, nil
}

func (s *GoodsServer) UpdateCategoryBrand(ctx context.Context, req *proto.CategoryBrandRequest) (*emptypb.Empty, error) {
	var categoryBrand model.GoodsCategoryBrand

	if result := global.DB.First(&categoryBrand, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Brand category does not exist")
	}

	var category model.Category
	if result := global.DB.First(&category, req.CategoryId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Product category does not exist")
	}

	var brand model.Brands
	if result := global.DB.First(&brand, req.BrandId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Brand does not exist")
	}

	categoryBrand.CategoryID = req.CategoryId
	categoryBrand.BrandsID = req.BrandId

	global.DB.Save(&categoryBrand)

	return &emptypb.Empty{}, nil
}
