package user

import (
	// "time"

	"github.com/set2002satoshi/8-4/pkg/module/dto/response"
	"github.com/set2002satoshi/8-4/interfaces/database"
	DBuser "github.com/set2002satoshi/8-4/interfaces/database/user"
	"github.com/set2002satoshi/8-4/models"
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

func (uc *UsersController) convertActiveToDTO(au *models.ActiveUser) *response.UserEntity {
	u := response.UserEntity{}
	u.ID = int(au.GetID())
	u.Name = au.GetName()
	u.Password = string(au.GetPassword())
	o := response.Options{}
	o.Revision = int(au.GetRevision())
	o.CratedAt = au.GetCreatedAt()
	o.UpdatedAt = au.GetUpdatedAt()
	u.Option = o
	return &u
}

func (uc *UsersController) convertHistoryToDTO(au *models.HistoryUser) *response.UserEntity {
	u := response.UserEntity{}
	u.ID = int(au.GetID())
	u.Name = au.GetName()
	u.Password = string(au.GetPassword())
	o := response.Options{}
	o.Revision = int(au.GetRevision())
	o.CratedAt = au.GetCreatedAt()
	o.UpdatedAt = au.GetUpdatedAt()
	u.Option = o
	return &u
}

func (uc *UsersController) convertActivesToDTOs(au []*models.ActiveUser) []*response.UserEntity {
	u := []*response.UserEntity{}
	for _, v := range au {
		user := &response.UserEntity{}
		user.ID = int(v.GetID())
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
