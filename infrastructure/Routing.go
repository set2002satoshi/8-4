package infrastructure

import (
	"github.com/gin-gonic/gin"

	"github.com/set2002satoshi/8-4/pkg/module/service/authentication/auth"

	bc "github.com/set2002satoshi/8-4/interfaces/controllers/blog"
	uc "github.com/set2002satoshi/8-4/interfaces/controllers/user"
)

type Routing struct {
	DB   *DB
	Gin  *gin.Engine
	Port string
}

func NewRouting(db *DB) *Routing {
	r := &Routing{
		DB:   db,
		Gin:  gin.Default(),
		Port: ":8080",
	}
	r.setRouting()
	return r
}

func (r *Routing) setRouting() {

	usersController := uc.NewUsersController(r.DB)
	blogsController := bc.NewBlogsController(r.DB)

	userNotLoggedIn := r.Gin.Group("/api")
	{
		// auth
		userNotLoggedIn.POST("/auth/login", func(c *gin.Context) { usersController.Login(c) })

		// user
		userNotLoggedIn.POST("/users", func(c *gin.Context) { usersController.Create(c) })
		userNotLoggedIn.GET("/users", func(c *gin.Context) { usersController.FindAll(c) })
		userNotLoggedIn.PUT("/users", func(c *gin.Context) { usersController.FindByID(c) })

		// blog
		userNotLoggedIn.GET("/blogs", func(c *gin.Context) { blogsController.Find(c) })
		userNotLoggedIn.PUT("/blog", func(c *gin.Context) { blogsController.FindByID(c) })
	}

	userLoggedIn := r.Gin.Group("/api")
	userLoggedIn.Use(auth.CheckLoggedIn())
	{

		// user
		userLoggedIn.POST("/users/update", func(c *gin.Context) { usersController.Update(c) })
		userLoggedIn.DELETE("/users", func(c *gin.Context) { usersController.Delete(c) })

		// blog
		userLoggedIn.POST("/blog", func(c *gin.Context) { blogsController.Create(c) })
		userLoggedIn.POST("/blog/update", func(c *gin.Context) { blogsController.Update(c) })
		userLoggedIn.DELETE("/blog", func(c *gin.Context) { blogsController.Delete(c) })
	}
}

func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}
