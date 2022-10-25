package infrastructure

import (
	"github.com/gin-gonic/gin"

	uc "github.com/set2002satoshi/8-4/interfaces/controllers/user"
	bc "github.com/set2002satoshi/8-4/interfaces/controllers/blog"
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
	user := r.Gin.Group("/api")
	{
		usersController := uc.NewUsersController(r.DB)
		user.GET("/users", func(c *gin.Context) { usersController.FindAll(c) })
		user.PUT("/users", func(c *gin.Context) { usersController.FindByID(c) })
		user.POST("/users", func(c *gin.Context) { usersController.Create(c) })
		user.POST("/users/update", func(c *gin.Context) { usersController.Update(c)} )
		user.DELETE("/users", func(c *gin.Context) { usersController.Delete(c) })
	}
	blog := r.Gin.Group("/api")
	{	
		blogsController := bc.NewBlogsController(r.DB)
		blog.GET("/blogs", func(c *gin.Context) {blogsController.Find(c)})
		blog.PUT("/blog", func(c *gin.Context) { blogsController.FindByID(c)})
		blog.POST("/blog", func (c *gin.Context) { blogsController.Create(c)})
		blog.POST("/blog/update", func (c *gin.Context) { blogsController.Update(c)})
		blog.DELETE("/blog", func(c *gin.Context) { blogsController.Delete(c)})
	}
}

func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}
