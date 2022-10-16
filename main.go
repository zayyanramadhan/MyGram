package main

import (
	"final_MyGram/controllers"
	"final_MyGram/db"
	_ "final_MyGram/docs"
	"final_MyGram/repositories"
	"final_MyGram/routes"
	"log"

	"github.com/joho/godotenv"
)

// @title MyGram API
// @description Sample API Spec for MyGram
// @version v2.0
// @termsOfService http://swagger.io/terms/
// @BasePath /
// @host localhost:7000
// @contact.name zayyan
// @contact.email zayyanramadhan@gmail.com

func main() {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatalf("Error read env file with err: %s", errEnv)
	}

	db.Connect()
	mdb := db.DbManager()
	repoAuth := repositories.NewauthRepository(mdb)
	repoUser := repositories.NewUserRepository(mdb)
	repoPhoto := repositories.NewPhotoRepository(mdb)
	repoComment := repositories.NewCommentRepository(mdb)
	repoSocialmedia := repositories.NewSocialMediaRepository(mdb)
	authController := controllers.NewAuthController(&repoAuth)
	userController := controllers.NewUserController(&repoUser, &repoAuth)
	photoController := controllers.NewPhotoController(&repoPhoto, &repoAuth)
	commentController := controllers.NewCommentController(&repoComment, &repoAuth)
	socialmediaController := controllers.NewSocialMediaController(&repoSocialmedia, &repoAuth)

	app := routes.NewRouter(authController, userController, photoController, commentController, socialmediaController)
	app.Start(":7000")
}
