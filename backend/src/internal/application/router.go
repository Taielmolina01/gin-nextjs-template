package application

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	authController "movie-reservation-system/controller"
	userController "movie-reservation-system/controller"
	authRepository "movie-reservation-system/repository/auth"
	userRepository "movie-reservation-system/repository/user"
	authService "movie-reservation-system/service/auth"
	userService "movie-reservation-system/service/user"
	"movie-reservation-system/infrastructure/middleware"
	"movie-reservation-system/configuration"
)

func Init(db *gorm.DB, config *configuration.Configuration) *gin.Engine {

	userRepo := userRepository.CreateRepositoryImpl(db)

	userController := setUpUserLayers(db, userRepo)
	authController := setUpAuthLayers(db, userRepo, config)

	router := gin.Default()

	addCorsConfiguration(router)

	setUpUserRoutes(router, db, userController)
	setUpAuthRoutes(router, db, authController)

	return router
}

func setUpUserLayers(db *gorm.DB, userRepo userRepository.UserRepository) *userController.UserController {
	userService := userService.NewUserServiceImpl(userRepo)

	userController := userController.NewUserController(userService)

	return userController
}

func setUpAuthLayers(db *gorm.DB, userRepo userRepository.UserRepository, config *configuration.Configuration) *authController.AuthController {
	authRepo := authRepository.NewAuthRepositoryImpl(db)

	authService := authService.NewAuthService(authRepo, userRepo, config.JwtAlgorithm, config.JwtSecretKey)

	authController := authController.NewAuthController(authService)

	return authController
}

func addCorsConfiguration(router *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	router.Use(cors.New(config))
}

func setUpUserRoutes(router *gin.Engine, db *gorm.DB, userController *userController.UserController) {
	usersGroup := router.Group("/users")
	{
		usersGroup.POST("", userController.CreateUser)
		usersGroup.GET("/:email", userController.GetUser)
		usersGroup.PUT("/:email", middleware.AuthMiddleware(db), userController.UpdateUser)
		usersGroup.PUT("/:email/password", middleware.AuthMiddleware(db),userController.UpdateUserPassword)
		usersGroup.DELETE("/:email", middleware.AuthMiddleware(db),userController.DeleteUser)
	}
}

func setUpAuthRoutes(router *gin.Engine, db *gorm.DB, authController *authController.AuthController) {
	router.POST("/login", authController.Login)
	router.POST("/logout/:email", middleware.AuthMiddleware(db), authController.Logout)
}
