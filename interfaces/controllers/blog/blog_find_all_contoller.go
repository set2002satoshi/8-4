package blog

import (

	c "github.com/set2002satoshi/8-4/interfaces/controllers"

	"github.com/set2002satoshi/8-4/pkg/module/dto/response"
)

type (
	FindAllBlogResponse struct {
		response.FindAllBlogResponse
	}
)

func (r *FindAllBlogResponse) SetErr(err error, errMsg string) {
	r.CodeErr = err
	r.MsgErr = errMsg
}


func (bc *BlogsController) Find(ctx c.Context) {


	res := &FindAllBlogResponse{}
	blogAll, err := bc.Interactor.FindAll()
	if err != nil {
		res.SetErr(err, "err")
		ctx.JSON(404, c.NewH(err.Error(), res))
		return 
	}
	res.Results = &response.ActiveBlogResults{Blogs: bc.convertActiveToDTOs(blogAll)}
	
	ctx.JSON(200, c.NewH("success", res))
}