package main

import (
	"BE-REPO/config"
	bu "BE-REPO/features/comment/repositorycomment"
	pu "BE-REPO/features/posting/repositoryposting"
	uh "BE-REPO/features/user/hendleruser"
	ur "BE-REPO/features/user/repositoryuser"
	us "BE-REPO/features/user/serviceuser"

	"BE-REPO/routes"
	"BE-REPO/utils/database"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	if cfg == nil {
		e.Logger.Fatal("Tidak bisa start server kesalahan database")
	}

	db, err := database.InitMysql(*cfg)
	if err != nil {
		e.Logger.Fatal("tidak bisa start bro", err.Error())
	}

	db.AutoMigrate(&ur.UserModel{}, &pu.PostingModel{}, &bu.CommentModel{})

	// hash := enkrip.New()
	userRepo := ur.New(db)
	userService := us.New(userRepo)
	userHandler := uh.New(userService)

	routes.InitRoute(e, userHandler)

	e.Logger.Fatal(e.Start(":8000"))
}
