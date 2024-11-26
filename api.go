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
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(gin.Logger())
	// services\users\views
	router.LoadHTMLGlob("services/users/views/*")

	v1 := router.Group("/api/v1")

	v1.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome To Learn_oo API V1")
	})

	subRoute := v1.Group("/users/auth")
	authRouter := routers.NewAuthRouter(subRoute)
	authRouter.RegisterRoutes()

	subRoute = v1.Group("/users/profiles")
	profilesRouter := routers.NewProfilesRouter(subRoute)
	profilesRouter.RegisterRoutes()

	fmt.Printf("Listening and serving at %v\n", "http://"+server.address+"/api/v1/")
	if err := router.Run(server.address); err != nil {
		panic(err)
	}
}
