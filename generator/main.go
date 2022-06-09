package generator

import (
	"fmt"
	qrCode "github.com/skip2/go-qrcode"
)

func QrCodeFile(fileName string, field1 string, value1 string,
	field2 string, value2 string) error {
	/*
		Generate a QR code with the JSON-encoded content
			{
				"<field1>":<value1>,
				"<field2>":<value2>,
			}
		...then stores it in a file with the given name.
	*/
	return qrCode.WriteFile(
		fmt.Sprintf("{\"%s\":\"%s\",\"%s\":\"%s\"}", field1, value1, field2, value2),
		qrCode.Medium, 256, fileName)
}

func QrCodeImage(field1 string, value1 string, field2 string, value2 string) ([]byte, error) {
	/*
		Generate a QR code with the JSON-encoded content
			{
				"<field1>":<value1>,
				"<field2>":<value2>,
			}
		and return the []byte array representing the raw image/png content.
	*/
	return qrCode.Encode(
		fmt.Sprintf("{\"%s\":\"%s\",\"%s\":\"%s\"}", field1, value1, field2, value2),
		qrCode.Medium, 256)
}
