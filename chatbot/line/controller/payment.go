package controller

import (
	"bytes"
	"encoding/base64"
	"strings"
	"work/wushu-backend/modules/connections"

	"github.com/polds/imgbase64"
)

func SaveProofOfPayment(url, filename string) int {
	bucket := connections.FirebaseStorage()
	img := imgbase64.FromRemote("http://somedomain.com/animage.jpg")
	idx := strings.Index(img, ";base64,")
	dec, err := base64.StdEncoding.DecodeString(img[idx+8:])
	if err != nil {
		return 1
	}
	res := bytes.NewReader(dec)
	if idx < 0 {
		return 1
	}

	if _, err := connections.PostFileFirebaseStorage(bucket, filename, res); err != nil {
		return 1
	}
	return 0
}
