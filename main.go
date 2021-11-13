package main

import (
	"log"
	"os"

	"github.com/yeqown/go-qrcode"
)

func main() {
	url := os.Args[1]
	qrCode, qrError := qrcode.New(url, qrcode.WithBorderWidth(1, 1, 1, 1), qrcode.WithQRWidth(1), qrcode.WithCustomImageEncoder(&Encoder{}))
	if qrError != nil {
		log.Fatalln("Oopsie: ", qrError)
	}

	qrCode.SaveTo(os.Stdout)
}
