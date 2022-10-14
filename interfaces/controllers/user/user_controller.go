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
	u := response.ActiveUserEntity{}
	u.ID = int(au.GetID())
	u.Email = au.GetEmail()
	u.Name = au.GetName()
	u.Password = string(au.GetPassword())
	o := response.Options{}
	o.Revision = int(au.GetRevision())
	o.CratedAt = au.GetCreatedAt()
	o.UpdatedAt = au.GetUpdatedAt()
	u.Option = o
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
	o.CratedAt = au.GetCreatedAt()
	o.UpdatedAt = au.GetUpdatedAt()
	u.Option = o
	return &u
}

func (uc *UsersController) convertActivesToDTOs(au []*models.ActiveUser) []*response.ActiveUserEntity {
	u := []*response.ActiveUserEntity{}
	for _, v := range au {
		user := &response.ActiveUserEntity{}
		user.ID = int(v.GetID())
		user.Email = v.GetEmail()
		user.Name = v.GetName()
		user.Password = string(v.GetPassword())
		o := response.Options{}
		o.Revision = int(v.GetRevision())
		o.CratedAt = v.GetCreatedAt()
		o.UpdatedAt = v.GetUpdatedAt()
		user.Option = o
		u = append(u, user)
	}
	return u
}
