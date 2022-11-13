package user

import (
	// "time"

	"github.com/set2002satoshi/8-4/interfaces/database"
	DBuser "github.com/set2002satoshi/8-4/interfaces/database/user"
	"github.com/set2002satoshi/8-4/models"
	"github.com/set2002satoshi/8-4/pkg/module/dto/response"
	usecase "github.com/set2002satoshi/8-4/usecase/user"
)

type UsersController struct {
	Interactor usecase.UserInteractor
}

func NewUsersController(db database.DB) *UsersController {
	return &UsersController{
		Interactor: usecase.UserInteractor{
			DB:   &database.DBRepository{DB: db},
			User: &DBuser.UserRepository{},
		},
	}
}

func (uc *UsersController) convertActiveToDTO(au *models.ActiveUser) *response.ActiveUserEntity {
	bl := []response.ActiveBlogEntity{}
	for _, v := range au.Blogs {
		b := response.ActiveBlogEntity{
			ID:           v.GetID(),
			ActiveUserID: v.GetActiveUserID(),
			Name:         v.GetName(),
			Title:        v.GetTitle(),
			Context:      v.GetContext(),
			Option: response.Options{
				Revision:  int(v.GetRevision()),
				CreatedAt: v.GetCreatedAt(),
				UpdatedAt: v.GetUpdatedAt(),
			},
		}
		bl = append(bl, b)
	}

	u := response.ActiveUserEntity{
		ID:       int(au.GetID()),
		Email:    au.GetEmail(),
		Name:     au.GetName(),
		Password: string(au.GetPassword()),
		Blogs:    bl,
		Option: response.Options{
			Revision:  int(au.GetRevision()),
			CreatedAt: au.GetCreatedAt(),
			UpdatedAt: au.GetUpdatedAt(),
		},
	}
	return &u
}

func (uc *UsersController) convertHistoryToDTO(au *models.HistoryUser) *response.HistoryUserEntity {
	u := response.HistoryUserEntity{}
	u.ID = int(au.GetID())
	u.ActiveUserID = int(au.GetActiveID())
	u.Email = au.GetEmail()
	u.Name = au.GetName()
	u.Password = string(au.GetPassword())
	o := response.Options{}
	o.Revision = int(au.GetRevision())
	o.CreatedAt = au.GetCreatedAt()
	o.UpdatedAt = au.GetUpdatedAt()
	u.Option = o
	return &u
}

func (uc *UsersController) convertActivesToDTOs(au []*models.ActiveUser) []*response.ActiveUserEntity {
	u := []*response.ActiveUserEntity{}
	for _, v := range au {
		bl := []response.ActiveBlogEntity{}
		for _, b := range v.Blogs {
			be := response.ActiveBlogEntity{
				ID:           b.GetID(),
				ActiveUserID: b.GetActiveUserID(),
				Name:         b.GetName(),
				Title:        b.GetTitle(),
				Context:      b.GetContext(),
				Option: response.Options{
					Revision:  int(b.GetRevision()),
					CreatedAt: b.GetCreatedAt(),
					UpdatedAt: b.GetUpdatedAt(),
				},
			}
			bl = append(bl, be)
		}
		user := &response.ActiveUserEntity{
			ID:       int(v.GetID()),
			Email:    v.GetEmail(),
			Name:     v.GetName(),
			Password: string(v.GetPassword()),
			Blogs:    bl,
			Option: response.Options{
				Revision:  int(v.GetRevision()),
				CreatedAt: v.GetCreatedAt(),
				UpdatedAt: v.GetUpdatedAt(),
			},
		}
		u = append(u, user)
	}
	return u
}
