QR Tools
========
(c) 2022 Sam Caldwell.  See LICENSE.txt

## About this project

This project derives from the 2019 QRHacks project I wrote to create QR code mayhem for a pen test.
That code was later used for the Combat Diver Foundation membership card stuff I did.  

In this project we just take the same work and create two simple tools (QR Code reader and
QR Code generator).

QR codes are a great way to transmit a good deal of information over a visual medium.  Here we are only
sending a static field and static value along with a dynamic uuid encoded as JSON in each QR.

## Functionality

### Commandline tool
There is a `main.go` that provides some simple examples and tests.  Compile it and
run as a command-line tool for testing, experimenting.

### generator
The `generator` provides a--
- `QrCodeFile()` function to generate a QR code PNG file.
- `QrCodeImage()` function to generate a raw PNG image byte array.

### reader
The `reader` provides--
- `QrFile()` function to consume a given PNG or JPEG file and return the QR Code content.
- `QrImage()` function to consume a given PNG or JPEG image byte array and return the QR code content.

