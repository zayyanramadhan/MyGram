package main

import (
	"final_MyGram/controllers"
	"final_MyGram/db"
	_ "final_MyGram/docs"
	"final_MyGram/repositories"
	"final_MyGram/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// @title MyGram API
// @description Sample API Spec for MyGram
// @version v2.0
// @termsOfService http://swagger.io/terms/
// @BasePath /
// @contact.name zayyan
// @contact.email zayyanramadhan@gmail.com

func main() {
	domain := ""
	port := os.Getenv("APP_PORT")
	if port == "" {
		errEnv := godotenv.Load(".env")
		if errEnv != nil {
			log.Fatalf("Error read env file with err: %s", errEnv)
		}
		port = os.Getenv("APP_PORT")
	} else {

	}
	runport := domain + ":" + port

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
	app.Start(runport)
}
