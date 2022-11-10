package blog

import (
	"time"

	c "github.com/set2002satoshi/8-4/interfaces/controllers"
	"github.com/set2002satoshi/8-4/models"
	"github.com/set2002satoshi/8-4/pkg/module/dto/request"
	"github.com/set2002satoshi/8-4/pkg/module/dto/response"
	"github.com/set2002satoshi/8-4/pkg/module/temporary"
)

type (
	CreateBlogResponse struct {
		response.CreateBlogResponse
	}
)


func (r *CreateBlogResponse) SetErr(err error, errMsg string) {
	r.CodeErr = err
	r.MsgErr = errMsg
}


func (bc *BlogsController) Create(ctx c.Context) {
	req := &request.BlogCreateRequest{}
	res := CreateBlogResponse{}

	if err := ctx.BindJSON(&req); err != nil {
		res.SetErr(err, "BindErr")
		ctx.JSON(404, res)
		return
	}
	reqModel, err := toModel(req)
	if err != nil {
		res.SetErr(err, "Models")
	}
	createdBlog, err := bc.Interactor.Post(reqModel)
	if err != nil {
		res.SetErr(err, "")
		ctx.JSON(404, res)
		return
	}

	res.Result = &response.ActiveBlogResult{Blog: bc.convertActiveToDTO(createdBlog)}
	ctx.JSON(201, c.NewH("ok", res))


}


func toModel(req *request.BlogCreateRequest) (*models.ActiveBlog, error) {
	return models.NewActiveBlog(
		temporary.INITIAL_ID,
		temporary.INITIAL_ID, // ここは後から修正
		req.Title,
		req.Context,
		time.Time{},
		time.Time{},
		time.Time{},
		temporary.INITIAL_REVISION,
	)
}