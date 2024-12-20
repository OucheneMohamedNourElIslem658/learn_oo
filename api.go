package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	certaficatesRouters "github.com/OucheneMohamedNourElIslem658/learn_oo/services/certaficates/routers"
	coursesRouters "github.com/OucheneMohamedNourElIslem658/learn_oo/services/courses/routers"
	usersRouters "github.com/OucheneMohamedNourElIslem658/learn_oo/services/users/routers"
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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	router.Use(gin.Logger())
	// services\users\views
	router.MaxMultipartMemory = 10 << 20
	router.LoadHTMLGlob("services/users/views/*")

	v1 := router.Group("/api/v1")

	v1.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome To Learn_oo API V1")
	})

	usersRouter := v1.Group("/users")

	subRoute := usersRouter.Group("/auth")
	authRouter := usersRouters.NewAuthRouter(subRoute)
	authRouter.RegisterRoutes()

	subRoute = usersRouter.Group("/profiles")
	profilesRouter := usersRouters.NewProfilesRouter(subRoute)
	profilesRouter.RegisterRoutes()

	subRoute = v1.Group("/courses")
	coursesRouter := coursesRouters.NewCoursesRouter(subRoute)
	coursesRouter.RegisterRoutes()

	subRoute = v1.Group("/user-courses")
	userCoursesRouter := certaficatesRouters.NewUserCourseRouter(subRoute)
	userCoursesRouter.RegisterRoutes()

	log.Printf("Listening and serving at %v\n", "http://"+server.address+"/api/v1/")
	if err := router.Run(server.address); err != nil {
		panic(err)
	}
}
