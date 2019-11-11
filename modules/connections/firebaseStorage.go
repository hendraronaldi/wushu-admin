package connections

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func FirebaseStorage() *storage.BucketHandle {
	config := &firebase.Config{
		StorageBucket: "admin-wushu.appspot.com",
	}
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
	app, err := firebase.NewApp(ctx, config, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Storage(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}

	return bucket
}

func GetFileFirebaseStorage(bucket *storage.BucketHandle, filename string) ([]byte, error) {
	rc, err := bucket.Object(filename).NewReader(context.Background())
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func PostFileFirebaseStorage(bucket *storage.BucketHandle, filename string, file *bytes.Reader) (string, error) {
	ctx := context.Background()
	// [START upload_file]
	obj := bucket.Object(filename)
	wc := obj.NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}
	// [END upload_file]

	// [START public]
	acl := obj.ACL()
	if err := acl.Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return "", err
	}
	// [END public]

	objAttrs, err := obj.Attrs(ctx)
	if err != nil {
		// TODO: Handle error.
		return "", err
	}

	return objAttrs.MediaLink, nil
}
