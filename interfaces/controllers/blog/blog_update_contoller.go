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
	UpdateBlogResponse struct {
		response.UpdateBlogResponse
	}
)


func (r *UpdateBlogResponse) SetErr(err error, errMsg string) {
	r.CodeErr = err
	r.MsgErr = errMsg
}


func (bc *BlogsController) Update(ctx c.Context) {
	req := &request.BlogUpdateRequest{}
	res := UpdateBlogResponse{}

	if err := ctx.BindJSON(&req); err != nil {
		res.SetErr(err, "BindErr")
		ctx.JSON(404, res)
		return
	}
	reqModel, err := bc.toModel(req)
	if err != nil {
		res.SetErr(err, "Models")
	}
	updatedBlog, err := bc.Interactor.Update(reqModel)
	if err != nil {
		res.SetErr(err, "")
		ctx.JSON(404, res)
		return
	}

	res.Result = &response.ActiveBlogResult{Blog: bc.convertActiveToDTO(updatedBlog)}
	ctx.JSON(201, c.NewH("ok", res))
}


func (bc *BlogsController) toModel(req *request.BlogUpdateRequest) (*models.ActiveBlog, error) {
	return models.NewActiveBlog(
		req.ID,
		temporary.INITIAL_ID, 
		req.Title,
		req.Context,
		time.Time{},
		time.Time{},
		time.Time{},
		temporary.REVISION(req.Revision),
	)
}