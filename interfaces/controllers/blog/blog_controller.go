package blog

import (
	"github.com/set2002satoshi/8-4/interfaces/database"
	DBblog "github.com/set2002satoshi/8-4/interfaces/database/blog"
	usecase "github.com/set2002satoshi/8-4/usecase/blog"

)

type BlogsController struct {
	Interactor usecase.BlogInteractor
}

func NewBlogsController(db database.DB) *BlogsController {
	return &BlogsController{
		Interactor: usecase.BlogInteractor{
			DB: &database.DBRepository{DB: db},
			Blog: &DBblog.BlogRepository{},
		},
	}
}