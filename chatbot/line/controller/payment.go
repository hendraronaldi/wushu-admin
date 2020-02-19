package controller

import (
	"bytes"
	"work/wushu-backend/modules/connections"
)

func SaveProofOfPayment(dec []byte, filename string) (string, int) {
	bucket := connections.FirebaseStorage()
	res := bytes.NewReader(dec)

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
