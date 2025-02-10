package application

import (
	"fmt"
	authController "github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/auth/controller"
	authService "github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/auth/service"
	userController "github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/controller"
	userRepository "github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/repository"
	userService "github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/service"
	"github.com/Taielmolina01/gin-nextjs-template/src/internal/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/postgres"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type Router struct {
	Engine *gin.Engine
	Port   string
}

func CreateRouter() (*Router, error) {

	engine := gin.Default()
	addCorsConfiguration(engine)

	config := LoadConfig()

	db, err := ConnectDB(config)

	if err != nil {
		return nil, err
	}

	sqlDB, err2 := db.DB() // Get the *sql.db from gorm.db
	if err2 != nil {
		log.Fatalf("Error al obtener *sql.DB: %v", err2)
	}

	store, err3 := postgres.NewStore(sqlDB, []byte(config.JWTSecretKey))
	if err3 != nil {
		log.Fatalf("Error al crear el store: %v", err3)
	}

	if err2 != nil {
		return nil, err2
	}

	engine.Use(sessions.Sessions("session", store))

	authMiddleware := middlewares.NewAuthMiddleware(config.JWTSecretKey)

	createEndPoints(engine, db, authMiddleware)

	return &Router{
		Engine: engine,
		Port:   config.Port,
	}, nil
}

func (router *Router) Run() {
	fmt.Println("Server is running on", router.Port)
	if err := router.Engine.Run(":" + router.Port); err != nil {
		log.Fatalln("Error running server: ", err)
	}
}

func createEndPoints(engine *gin.Engine, db *gorm.DB, authMiddleware *middlewares.AuthMiddleware) {
	userController, userRepo := setUpUserLayers(db)
	authController := setUpAuthLayers(db, userRepo, config)

	setUpUserRoutes(engine, db, userController, authMiddleware)
	setUpAuthRoutes(engine, db, authController, authMiddleware)
}

func setUpUserLayers(db *gorm.DB) (*userController.UserController, userRepository.UserRepository) {
	userRepo := userRepository.CreateRepositoryImpl(db)

	userService := userService.NewUserServiceImpl(userRepo)

	userController := userController.NewUserController(userService)

	return userController, userRepo
}

func setUpAuthLayers(db *gorm.DB, userRepo userRepository.UserRepository, config *Configuration) *authController.AuthController {
	authService := authService.NewAuthService(userRepo, config.JWTAlgorithm, config.JWTSecretKey)

	authController := authController.NewAuthController(authService)

	return authController
}

func addCorsConfiguration(engine *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	engine.Use(cors.New(config))
}

func setUpUserRoutes(engine *gin.Engine, db *gorm.DB, userController *userController.UserController, authMiddleware *middlewares.AuthMiddleware) {
	usersGroup := engine.Group("/users")
	{
		usersGroup.POST("", userController.CreateUser)
		usersGroup.GET("/:email", userController.GetUser)
		usersGroup.PUT("/:email", authMiddleware.PublicAuthMiddleware(db), userController.UpdateUser)
		usersGroup.PUT("/:email/password", authMiddleware.PublicAuthMiddleware(db), userController.UpdateUserPassword)
		usersGroup.DELETE("/:email", authMiddleware.PublicAuthMiddleware(db), userController.DeleteUser)
	}
}

func setUpAuthRoutes(engine *gin.Engine, db *gorm.DB, authController *authController.AuthController, authMiddleware *middlewares.AuthMiddleware) {
	engine.POST("/login", authController.Login)
	engine.POST("/logout/:email", authMiddleware.PublicAuthMiddleware(db), authController.Logout)
}
