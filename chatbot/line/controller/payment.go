package controller

import (
	"bytes"
	"work/wushu-backend/modules/connections"
)

func SaveProofOfPayment(dec []byte, filename string) (string, int) {
	bucket := connections.FirebaseStorage()
	// img := imgbase64.FromRemote("http://somedomain.com/animage.jpg")
	// idx := strings.Index(img, ";base64,")
	// dec, err := base64.StdEncoding.DecodeString(img[idx+8:])
	// if err != nil {
	// 	return 1
	// }
	res := bytes.NewReader(dec)
	// if idx < 0 {
	// 	return 1
	// }

	f, err := connections.PostFileFirebaseStorage(bucket, filename, res)
	if err != nil {
		return "", 1
	}
	return f, 0
}

func DeleteProofOfPayment(filename string) error {
	bucket := connections.FirebaseStorage()
	err := connections.DeleteFileFirebaseStorage(bucket, filename)
	if err != nil {
		return err
	}
	return nil
}
