QR Tools
========
(c) 2022 Sam Caldwell.  See LICENSE.txt

A simple set of QR code tools for testing, pen testing, observability or other
interesting use cases.

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

