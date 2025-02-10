package application

func CreateTestEngine(db *gorm.DB, config *Configuration) *gin.Engine {

	userRepo := userRepository.CreateRepositoryImpl(db)

	userController := setUpUserLayers(db, userRepo)
	authController := setUpAuthLayers(db, userRepo, config)

	router := gin.Default()

	addCorsConfiguration(router)

	setUpUserRoutes(router, db, userController)
	setUpAuthRoutes(router, db, authController)

	return router
}
