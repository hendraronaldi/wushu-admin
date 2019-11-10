package connections

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"log"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func FirebaseStorage() *storage.BucketHandle {
	ctx := context.Background()
	config := &firebase.Config{
		StorageBucket: "admin-wushu.appspot.com",
	}
	opt := option.WithCredentialsFile("utilities/admin-wushu-firebase.json")
	app, err := firebase.NewApp(ctx, config, opt)
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

func PostFileFirebaseStorage(bucket *storage.BucketHandle, filename string, file *bytes.Reader) error {
	ctx := context.Background()
	// [START upload_file]
	wc := bucket.Object(filename).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}
	// [END upload_file]
	return nil
}
