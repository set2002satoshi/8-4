package controllers

// type UsersController struct {
// 	Interactor usecase.UserInteractor
// }

// func NewUsersController(db database.DB) *UsersController {
// 	return &UsersController{
// 		Interactor: usecase.UserInteractor {
// 			// databaseからDB->(Begin, Connect)を配下に持つを代入
// 			DB: &database.DBRepository{ DB:db },
// 			User: &database.UserRepository{},
// 		},
// 	}
// }

// func (controller *UsersController) Get(c Context) {

// 	id, _ := strconv.Atoi(c.Param("id"));

// 	user, res := controller.Interactor.Get(id)
// 	if res != nil {
// 		// c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))
// 		c.JSON(404, NewH(res.Error(), nil))
// 		return
// 	}
// 	c.JSON(200, NewH("success", user))
// }

// func (controller *UsersController) Post(c Context) {
// 	// ここで行う処理 (多分userに一番近い処理)
// 	// userからデータを受け取る->usecaseにデータ渡す
// 	// 渡したデータの結果に応じてuserの対応をする
// 	var userForm domain.UsersForPost
// 	if err := c.BindJSON(&userForm); err != nil {
// 		c.JSON(400, NewH("400", nil))
// 		return
// 	}
// 	createdUser, res := controller.Interactor.Post(userForm)
// 	if res != nil {
// 		c.JSON(400, NewH("400", res))
// 		return
// 	}
// 	c.JSON(200, NewH("201", createdUser))
// }
