package blog

import (

	c "github.com/set2002satoshi/8-4/interfaces/controllers"
	"github.com/set2002satoshi/8-4/pkg/module/dto/request"
	"github.com/set2002satoshi/8-4/pkg/module/dto/response"
)

type (
	FindByIdBlogResponse struct {
		response.FindByIDBlogResponse
	}
)

func (r *FindByIdBlogResponse) SetErr(err error, errMsg string) {
	r.CodeErr = err
	r.MsgErr = errMsg
}


func (bc *BlogsController) FindByID(ctx c.Context) {
	
	req := &request.BlogFindByIDRequest{}
	res := &FindByIdBlogResponse{}

	if err := ctx.BindJSON(req); err != nil {
		res.SetErr(err, "BindErr")
	}

	blog, err := bc.Interactor.FindByID(req.ID)
	if err != nil {
		ctx.JSON(404, c.NewH(err.Error(), nil))
		return
	}
	res.Result = &response.ActiveBlogResult{Blog: bc.convertActiveToDTO(blog)}

	ctx.JSON(200, c.NewH("success", res))
}

