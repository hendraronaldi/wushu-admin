package connections

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func FirebaseConnection() *firestore.Client {
	ctx := context.Background()
	var err error
	var configJSON []byte

	if configJSONString := os.Getenv("firebase-auth"); configJSONString != "" {
		configJSON = []byte(configJSONString)
	} else {
		configJSON, err = ioutil.ReadFile("utilities/admin-wushu-firebase.json")
		if err != nil {
			fmt.Println("read file err", err)
			return nil
		}
	}
	sa := option.WithCredentialsJSON(configJSON)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return client
}

func FirebaseTutorialConnection() *firestore.Client {
	ctx := context.Background()
	var err error
	var configJSON []byte

	if configJSONString := os.Getenv("firebase-tutorial-auth"); configJSONString != "" {
		configJSON = []byte(configJSONString)
	} else {
		configJSON, err = ioutil.ReadFile("utilities/wushu-tutorial-firebase.json")
		if err != nil {
			fmt.Println("read file err", err)
			return nil
		}
	}
	sa := option.WithCredentialsJSON(configJSON)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return client
}
