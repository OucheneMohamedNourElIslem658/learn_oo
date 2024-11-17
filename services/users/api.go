package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	routers "github.com/OucheneMohamedNourElIslem658/learn_oo/services/users/routers"
)

type Server struct {
	address string
}

func NewServer(address string) *Server {
	return &Server{
		address: address,
	}
}

func (server *Server) Run() {
	router := gin.Default()

	router.Use(gin.Logger())
	router.LoadHTMLGlob("./views/*")

	v1 := router.Group("/api/v1")

	v1.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome To Learn_oo API V1")
	})

	subRoute := v1.Group("/users/auth/")
	authRouter := routers.NewAuthRouter(subRoute)
	authRouter.RegisterRoutes()

	fmt.Printf("Listening and serving at %v\n", "http://"+server.address+"/api/v1/")
	if err := router.Run(server.address); err != nil {
		panic(err)
	}
}
