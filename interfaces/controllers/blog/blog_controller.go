package blog

import (
	"github.com/set2002satoshi/8-4/interfaces/database"
	DBblog "github.com/set2002satoshi/8-4/interfaces/database/blog"
	"github.com/set2002satoshi/8-4/models"
	"github.com/set2002satoshi/8-4/pkg/module/dto/response"
	usecase "github.com/set2002satoshi/8-4/usecase/blog"
)

type BlogsController struct {
	Interactor usecase.BlogInteractor
}

func NewBlogsController(db database.DB) *BlogsController {
	return &BlogsController{
		Interactor: usecase.BlogInteractor{
			DB:   &database.DBRepository{DB: db},
			Blog: &DBblog.BlogRepository{},
		},
	}
}

func (bc *BlogsController) convertActiveToDTO(bu *models.ActiveBlog) *response.ActiveBlogEntity {

	o := response.Options{
		Revision:  int(bu.GetRevision()),
		CreatedAt: bu.GetCreatedAt(),
		UpdatedAt: bu.GetUpdatedAt(),
	}

	b := response.ActiveBlogEntity{}
	b.ID = bu.GetID()
	b.Name = bu.GetName()
	b.ActiveUserID = bu.GetActiveUserID()
	b.Title = bu.GetTitle()
	b.Context = bu.GetContext()
	b.Option = o

	return &b
}

func (bc *BlogsController) convertActiveToDTOs(bu []*models.ActiveBlog) []*response.ActiveBlogEntity {

	abr := []*response.ActiveBlogEntity{}

	for _, v := range bu {

		e := &response.ActiveBlogEntity{}
		o := response.Options{}
		o.Revision = int(v.GetRevision())
		o.CreatedAt = v.GetCreatedAt()
		o.UpdatedAt = v.GetUpdatedAt()

		e.ID = v.GetID()
		e.ActiveUserID = v.GetActiveUserID()
		e.Name = v.GetName()
		e.Title = v.GetTitle()
		e.Context = v.GetContext()
		e.Option = o
		abr = append(abr, e)
	}

	return abr
}

func (bc *BlogsController) convertHistoryToDTOs(bu []*models.HistoryBlog) []*response.HistoryBlogEntity {

	abr := []*response.HistoryBlogEntity{}

	for _, v := range bu {

		e := &response.HistoryBlogEntity{}
		o := response.Options{}
		o.Revision = int(v.GetRevision())
		o.CreatedAt = v.GetCreatedAt()
		o.UpdatedAt = v.GetUpdatedAt()

		e.ID = int(v.GetID())
		e.ActiveBlogID = int(v.GetActiveBlogID())
		e.Title = v.GetTitle()
		e.Context = v.GetContext()
		e.Option = o
		abr = append(abr, e)
	}

	return abr
}

func (bc *BlogsController) convertHistoryToDTO(bu *models.HistoryBlog) *response.HistoryBlogEntity {

	e := &response.HistoryBlogEntity{}
	o := response.Options{}
	o.Revision = int(bu.GetRevision())
	o.CreatedAt = bu.GetCreatedAt()
	o.UpdatedAt = bu.GetUpdatedAt()

	e.ID = int(bu.GetID())
	e.ActiveBlogID = int(bu.GetActiveBlogID())
	e.Title = bu.GetTitle()
	e.Context = bu.GetContext()
	e.Option = o

	return e
}
