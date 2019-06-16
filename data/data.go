package data

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var carriers map[string]string
var carrierMaxLenght = 0

type carrier struct {
	CountryCode string
	CarrierMNO  string
}

func LoadCarrierData() {
	carriers = make(map[string]string)
	root := "./data/carrier"
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatalf("failed reading a folder: %s", err)
	}

	for _, f := range files {
		file, err := os.Open(root + "/" + f.Name())

		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			var line = scanner.Text()
			if len(strings.TrimSpace(line)) == 0 || strings.HasPrefix(line, "#") {
				continue
			}
			var splitedLine = strings.Split(line, "|")
			carriers[splitedLine[0]] = f.Name()[:strings.IndexByte(f.Name(), '.')]

			if len(splitedLine[0]) > carrierMaxLenght {
				carrierMaxLenght = len(splitedLine[0])
			}
		}

		file.Close()
	}
}

func GetCarrier(msisdn string) *carrier {
	for i := carrierMaxLenght; i > 0; i-- {
		cc, exists := carriers[msisdn[:i]]
		if exists {
			return &carrier{CountryCode: cc, CarrierMNO: strings.TrimLeft(msisdn[:i], cc)}
		}
	}
	return nil
}
