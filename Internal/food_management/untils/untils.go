package fooduntils

import (
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
)

func UploadImageToFirebaseStorage(imageData []byte, fileName string, folderName string) (string, error) {
	serviceAccountKey := "path/to/serviceAccountKey.json"

	ctx := context.Background()
	opt := option.WithCredentialsFile(serviceAccountKey)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Failed to create Firebase app: %v\n", err)
		return "", err
	}

	storageClient, err := app.Storage(ctx)
	if err != nil {
		log.Fatalf("Failed to create Firebase Storage client: %v\n", err)
		return "", err
	}

	bucket, err := storageClient.DefaultBucket()
	if err != nil {
		log.Fatalf("Failed to get default bucket: %v\n", err)
		return "", err
	}

	filePath := folderName + "/" + fileName

	obj := bucket.Object(filePath)

	wc := obj.NewWriter(ctx)

	if _, err := wc.Write(imageData); err != nil {
		log.Fatalf("Failed to upload image: %v\n", err)
		return "", err
	}
	if err := wc.Close(); err != nil {
		log.Fatalf("Failed to close writer: %v\n", err)
		return "", err
	}

	attrs, err := obj.Attrs(ctx)
	if err != nil {
		log.Fatalf("Failed to get object attributes: %v\n", err)
		return "", err
	}

	return attrs.MediaLink, nil
}
