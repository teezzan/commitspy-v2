package config

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var firebaseAuth *auth.Client

func InitFirebase() {

	opt := option.WithCredentialsJSON([]byte(Cfg.Firebase.CredentialJSON))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	fbAUth, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Firebase Auth load error: %v", err)
	}
	firebaseAuth = fbAUth
}

func GetFirebaseAuthClient() *auth.Client {
	return firebaseAuth
}
