package blog

import (

	c "github.com/set2002satoshi/8-4/interfaces/controllers"
	"github.com/set2002satoshi/8-4/pkg/module/dto/request"
	"github.com/set2002satoshi/8-4/pkg/module/dto/response"
)


type (
	DeleteBlogResponse struct {
		response.DeleteBlogResponse
	}
)

func (r *DeleteBlogResponse) SetErr(err error, ErrMsg string) {
	r.CodeErr = err
	r.MsgErr = ErrMsg
}

func (bc *BlogsController) Delete(ctx c.Context) {

	req := &request.BlogDeleteRequest{}
	res := &DeleteBlogResponse{}
	if err := ctx.BindJSON(&req); err != nil {
		res.SetErr(err, "bindErr")
		ctx.JSON(404, res)
		return 
	}

	blog, err := bc.Interactor.DeleteByID(req.ID)
	if err != nil {
		ctx.JSON(404, c.NewH(err.Error(), nil))
		return
	}
	res.Result = &response.HistoryBlogResult{Blog: bc.convertHistoryToDTO(blog)}
	ctx.JSON(200, c.NewH("success", res))
}
