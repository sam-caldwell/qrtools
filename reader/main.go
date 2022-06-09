package reader

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/liyue201/goqr"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"strings"
)

func QrFile(fileName string) (content string, err error) {
	imageData, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return QrImage(imageData)
}

func QrImage(imageData []byte) (content string, err error) {
	img, _, err := image.Decode(bytes.NewReader(imageData))
	if err != nil {
		return content, errors.New(fmt.Sprintf("image.Decode error: %v\n", err))
	}

	qrCodes, err := goqr.Recognize(img)

	if err != nil {
		return content, errors.New(fmt.Sprintf("Recognize failed: %v\n", err))
	}
	if len(qrCodes) > 1 {
		return content, errors.New("Multiple qr codes found.  Only one expected.")
	}
	return strings.TrimRight(string(qrCodes[0].Payload), "\n"), err
}
