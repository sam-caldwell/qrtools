package main

/*
	(c) 2022 Sam Caldwell.  See LICENSE.txt
*/

import (
	"flag"
	"fmt"
	"github.com/google/uuid"
	"os"
	"qrtools/generator"
	"qrtools/reader"
	"strings"
)

const (
	argCommand         = "command"
	argStaticFieldName = "staticFieldName"
	argStaticValue     = "staticValue"

	cmdBadCommand   = "badCommand"
	cmdGenerateFile = "generateFile"
	cmdReadFile     = "readFile"
	cmdTestCycle    = "test"
)

func errCheck(e error) {
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1)
	}
}

func generateUuid() string {
	id, _ := uuid.NewUUID()
	return strings.ToLower(id.String())
}

func generateQrCodeAsFile(fileName *string, fieldName *string, fieldValue *string) {
	errCheck(generator.QrCodeFile(*fileName, *fieldName, *fieldValue, "uuid", generateUuid()))
}

func readQrCodeFile(fileName *string) (content string) {
	//Todo: implement function
	content, err := reader.QrFile(*fileName)
	errCheck(err)
	return content
}

func fullTestCycle(fieldName *string, fieldValue *string) (bool, string, string) {
	/*
		generate a random uuid and encode with the static inputs to create a signal.
		Then use that signal to generate a qr code raw image []byte array.
		Turn around and decode it and expect the inputs to match the outputs.
	*/
	thisId := generateUuid()
	expectedContent := fmt.Sprintf("{\"%s\":\"%s\",\"%s\":\"%s\"}", *fieldName, *fieldValue, "uuid", thisId)

	img, err := generator.QrCodeImage(*fieldName, *fieldValue, "uuid", thisId)
	errCheck(err)
	actualContent, err := reader.QrImage(img)
	return actualContent == expectedContent, actualContent, expectedContent
}

func main() {
	commandPtr := flag.String(argCommand, cmdBadCommand, "Command (or mode)")
	filePtr := flag.String("file", "", "QR Code filename (empty if not specified)")
	staticFieldName := flag.String(argStaticFieldName, "", "The static field name we will render in the QR code")
	staticFieldValue := flag.String(argStaticValue, "", "the static value")
	flag.Parse()

	fmt.Printf("\ninputs:\n\tcommand:%s\n\tfilename:%s\n\tstaticField:%s\n\tstaticValue:%s\n\n",
		*commandPtr, *filePtr, *staticFieldName, *staticFieldValue)

	switch *commandPtr {
	case cmdGenerateFile:
		if *filePtr == "" {
			fmt.Println("You must specify -file=<filename>")
		}
		if *staticFieldName == "" {
			fmt.Printf("You must specify %s with %s\n", argStaticFieldName, *commandPtr)
		}
		if *staticFieldValue == "" {
			fmt.Printf("You must specify %s with %s\n", argStaticValue, *commandPtr)
		}
		generateQrCodeAsFile(filePtr, staticFieldName, staticFieldValue)
		fmt.Printf("generated file: %s\n", *filePtr)
		break
		/**/
	case cmdReadFile:
		if *filePtr == "" {
			fmt.Println("You must specify -file=<filename>")
		}
		if *staticFieldName != "" {
			fmt.Printf("Do not specify %s with %s\n", argStaticFieldName, *commandPtr)
		}
		if *staticFieldValue != "" {
			fmt.Printf("Do not specify %s with %s\n", argStaticValue, *commandPtr)
		}
		content := readQrCodeFile(filePtr)
		fmt.Printf("content:%s\n", content)
		break
		/**/
	case cmdTestCycle:
		if *filePtr != "" {
			fmt.Printf("Do not specify -file=<filename> with %s\n", cmdTestCycle)
		}
		if *staticFieldName == "" {
			fmt.Printf("You must specify %s with %s\n", argStaticFieldName, *commandPtr)
		}
		if *staticFieldValue == "" {
			fmt.Printf("You must specify %s with %s\n", argStaticValue, *commandPtr)
		}
		result, a, e := fullTestCycle(staticFieldName, staticFieldValue)
		if result {
			fmt.Println("test pass")
		} else {
			fmt.Println("test failed")
			fmt.Printf("actual:   |%s|\n", a)
			fmt.Printf("expected: |%s|\n", e)
			os.Exit(2)
		}
		/**/
	default:
		fmt.Println("Error: unknown command")
		os.Exit(1)
	}
	os.Exit(0)
}
