package category

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	empty "github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"

	"apis/goods-web/api"
	"apis/goods-web/forms"
	"apis/goods-web/global"
	"apis/goods-web/proto/gen"
)

func List(c context.Context, ctx *app.RequestContext) {
	r, err := global.GoodsSrvClient.GetAllCategorysList(c, &empty.Empty{})
	if err != nil {
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}

	data := make([]interface{}, 0)
	err = json.Unmarshal([]byte(r.JsonData), &data)
	if err != nil {
		zap.S().Errorw("[List] search error", err.Error())
	}

	ctx.JSON(http.StatusOK, data)
}

func Detail(c context.Context, ctx *app.RequestContext) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	reMap := make(map[string]interface{})
	subCategorys := make([]interface{}, 0)
	if r, err := global.GoodsSrvClient.GetSubCategory(c, &proto.CategoryListRequest{
		Id: int32(i),
	}); err != nil {
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	} else {
		for _, value := range r.SubCategorys {
			subCategorys = append(subCategorys, map[string]interface{}{
				"id":              value.Id,
				"name":            value.Name,
				"level":           value.Level,
				"parent_category": value.ParentCategory,
				"is_tab":          value.IsTab,
			})
		}
		reMap["id"] = r.Info.Id
		reMap["name"] = r.Info.Name
		reMap["level"] = r.Info.Level
		reMap["parent_category"] = r.Info.ParentCategory
		reMap["is_tab"] = r.Info.IsTab
		reMap["sub_categorys"] = subCategorys

		ctx.JSON(http.StatusOK, reMap)
	}
	return
}

func New(c context.Context, ctx *app.RequestContext) {
	categoryForm := forms.CategoryForm{}
	if err := ctx.BindAndValidate(&categoryForm); err != nil {
		api.HandleValidatorError(ctx, err)
		return
	}

	rsp, err := global.GoodsSrvClient.CreateCategory(c, &proto.CategoryInfoRequest{
		Name:           categoryForm.Name,
		IsTab:          *categoryForm.IsTab,
		Level:          categoryForm.Level,
		ParentCategory: categoryForm.ParentCategory,
	})
	if err != nil {
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}

	request := make(map[string]interface{})
	request["id"] = rsp.Id
	request["name"] = rsp.Name
	request["parent"] = rsp.ParentCategory
	request["level"] = rsp.Level
	request["is_tab"] = rsp.IsTab

	ctx.JSON(http.StatusOK, request)
}

func Delete(c context.Context, ctx *app.RequestContext) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	_, err = global.GoodsSrvClient.DeleteCategory(c, &proto.DeleteCategoryRequest{Id: int32(i)})
	if err != nil {
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}

	ctx.Status(http.StatusOK)
}

func Update(c context.Context, ctx *app.RequestContext) {
	categoryForm := forms.UpdateCategoryForm{}
	if err := ctx.BindAndValidate(&categoryForm); err != nil {
		api.HandleValidatorError(ctx, err)
		return
	}

	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	request := &proto.CategoryInfoRequest{
		Id:   int32(i),
		Name: categoryForm.Name,
	}
	if categoryForm.IsTab != nil {
		request.IsTab = *categoryForm.IsTab
	}
	_, err = global.GoodsSrvClient.UpdateCategory(c, request)
	if err != nil {
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}

	ctx.Status(http.StatusOK)
}
