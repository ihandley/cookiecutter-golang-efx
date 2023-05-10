package router

import (
	"github.com/gin-gonic/gin"
	"midigator-portfolios/cookiecutter-golang/api/rest/controllers"
	"midigator-portfolios/cookiecutter-golang/app/initializer"
)

// Init sets router
func Init(dependencies initializer.Services) *gin.Engine {
	router := gin.Default()

	ping := controllers.NewPingController(dependencies)
	todo := controllers.NewTodoController(dependencies)
	router.GET("/ping", ping.Ping)
	router.GET("/todos", todo.List)
	router.POST("/todos", todo.Create)

	return router
}
