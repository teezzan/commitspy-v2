package config

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var FirebaseAuth *auth.Client

func init() {

	opt := option.WithCredentialsJSON([]byte(Cfg.Firebase.CredentialJSON))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	} else {
		log.Println("Successfully initiated Firebase App!")
	}

	auth, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Firebase Auth load error: %v", err)
	}
	FirebaseAuth = auth
}