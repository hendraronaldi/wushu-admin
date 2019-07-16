package connections

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func FirebaseConnection() *firestore.Client {
	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile("utilities/admin-wushu-firebase.json")
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

// func FirebaseConnection() *firego.Firebase {
// 	var configJSON []byte
// 	var err error
// 	if configJSONString := os.Getenv("firebase-auth"); configJSONString != "" {
// 		configJSON = []byte(configJSONString)
// 	} else {
// 		configJSON, err = ioutil.ReadFile("utilities/firebase-board-engine.json")
// 		if err != nil {
// 			fmt.Println("read file err", err)
// 			return nil
// 		}
// 	}

// 	conf, err := google.JWTConfigFromJSON(configJSON, "https://www.googleapis.com/auth/userinfo.email",
// 		"https://www.googleapis.com/auth/firebase.database")
// 	if err != nil {
// 		fmt.Println("jwt config error", err)
// 		return nil
// 	}

// 	fb := firego.New("https://board-a99ec.firebaseio.com/", conf.Client(oauth2.NoContext))
// 	return fb
// }
