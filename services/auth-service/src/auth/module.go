package auth

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type AuthModule struct {
	App *echo.Echo
	Db  *gorm.DB
}

func (module AuthModule) RegisterModule() {
	ctx := context.Background()

	fbApp, err := firebase.NewApp(ctx, &firebase.Config{ProjectID: "ehealth-local-project"})
	if err != nil {
		log.Fatalf("Firebase App error: %v", err)
	}

	fbAuth, err := fbApp.Auth(ctx)
	if err != nil {
		log.Fatalf("Firebase Auth error: %v", err)
	}

	service := AuthService{
		Db:           module.Db,
		FirebaseAuth: fbAuth,
	}

	controller := &AuthController{Service: &service}

	controller.RegisterEndpoints(module.App)
}
