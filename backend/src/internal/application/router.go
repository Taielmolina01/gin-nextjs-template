package application

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"fmt"
	"log"
	authController "github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/auth/controller"
	userController "github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/controller"
	userRepository "github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/repository"
	authService "github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/auth/service"
	userService "github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/service"
	"github.com/Taielmolina01/gin-nextjs-template/src/internal/middlewares"
	"github.com/gin-contrib/sessions"
  	"github.com/gin-contrib/sessions/postgres"
)

type Router struct {
	Engine *gin.Engine
	Port	string
}


func CreateRouter(port string) *Router {

	engine := gin.Default()
	addCorsConfiguration(engine)

	config := LoadConfig()

	db := ConnectDB(config)

	store, err := postgres.NewStore(db, []byte(config.secretKey))

	engine.Use(sessions.Sessions("session", store))

	middlewares.NewAuthMiddleware(config.JWTSecretKey)

	createEndPoints(db)

	return &Router{
		engine: engine,
		port: config.Port,
	}
}

func (router *Router) Run() {
	fmt.Println("Server is running on", router.port)
	if err := router.engine.Run(r.port); err != nil {
		log.Fatalln("Error running server: ", err)
	}
}

func createEndPoints(db *gorm.DB) {
	userController, userRepo := setUpUserLayers(db)
	authController := setUpAuthLayers(db, userRepo, config)

	setUpUserRoutes(engine, db, userController)
	setUpAuthRoutes(engine, db, authController)
}

func setUpUserLayers(db *gorm.DB) (*userController.UserController, userRepository.UserRepository) {
	userRepo := userRepository.CreateRepositoryImpl(db)

	userService := userService.NewUserServiceImpl(userRepo)

	userController := userController.NewUserController(userService)

	return userController, userRepo
}

func setUpAuthLayers(db *gorm.DB, userRepo userRepository.UserRepository, config *Configuration) *authController.AuthController {
	authService := authService.NewAuthService(userRepo, config.JwtAlgorithm, config.JwtSecretKey)

	authController := authController.NewAuthController(authService)

	return authController
}

func addCorsConfiguration(engine *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	router.Use(cors.New(config))
}

func setUpUserRoutes(engine *gin.Engine, db *gorm.DB, userController *userController.UserController) {
	usersGroup := router.Group("/users")
	{
		usersGroup.POST("", userController.CreateUser)
		usersGroup.GET("/:email", userController.GetUser)
		usersGroup.PUT("/:email", middlewares.PublicAuthMiddleware(db), userController.UpdateUser)
		usersGroup.PUT("/:email/password", middlewares.PublicAuthMiddleware(db),userController.UpdateUserPassword)
		usersGroup.DELETE("/:email", middlewares.PublicAuthMiddleware(db),userController.DeleteUser)
	}
}

func setUpAuthRoutes(engine *gin.Engine, db *gorm.DB, authController *authController.AuthController) {
	router.POST("/login", authController.Login)
	router.POST("/logout/:email", middlewares.PublicAuthMiddleware(db), authController.Logout)
}
