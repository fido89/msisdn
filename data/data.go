package data

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var instance *Data
var once sync.Once

func GetInstance() *Data {
	once.Do(func() {
		instance = &Data{}
		instance.loadCarrierData()
		instance.loadCountryIsoIds()
	})
	return instance
}

type DataGetter interface {
	GetCarrier(string) *Carrier
	GetCountryIsoCode(string) *string
}

type Data struct {
	carriers         map[string]string
	carrierMaxLenght int
	countryIsoCodes  map[string]string
}

type Carrier struct {
	CountryCode string
	CarrierMNO  string
}

func (d *Data) loadCarrierData() {
	d.carriers = make(map[string]string)
	root := "./data/carrier"
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatalf("failed reading a folder: %s", err)
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) != ".txt" {
			continue
		}

		file, err := os.Open(root + "/" + f.Name())

		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			carrierData, countryCode := getCarrierAndCountryCode(scanner.Text(), f.Name())
			if carrierData != nil {
				d.carriers[*carrierData] = *countryCode

				if len(*carrierData) > d.carrierMaxLenght {
					d.carrierMaxLenght = len(*carrierData)
				}
			}
		}
		file.Close()
	}
}

func getCarrierAndCountryCode(line string, fileName string) (*string, *string) {
	if len(strings.TrimSpace(line)) == 0 || strings.HasPrefix(line, "#") {
		return nil, nil
	}
	var lineWithCarrierData = strings.Split(line, "|")
	var countryCode = fileName[:strings.IndexByte(fileName, '.')]
	return &lineWithCarrierData[0], &countryCode
}

func (d *Data) loadCountryIsoIds() {
	jsonFile, err := os.Open("./data/country-calling-codes.json")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	d.countryIsoCodes = make(map[string]string)
	var results []map[string]interface{}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &results)
	for _, result := range results {
		d.countryIsoCodes[fmt.Sprintf("%v", result["callingCode"])] = fmt.Sprintf("%v", result["code"])
	}
}

func (d *Data) GetCarrier(msisdn string) *Carrier {
	var carrierMaxLenghtLocal = d.carrierMaxLenght
	if len(msisdn) < carrierMaxLenghtLocal {
		carrierMaxLenghtLocal = len(msisdn)
	}

	for i := carrierMaxLenghtLocal; i > 0; i-- {
		cc, exists := d.carriers[msisdn[:i]]
		if exists {
			return &Carrier{CountryCode: cc, CarrierMNO: strings.TrimPrefix(msisdn[:i], cc)}
		}
	}
	return nil
}

func (d *Data) GetCountryIsoCode(callingCode string) *string {
	iso, exists := d.countryIsoCodes[callingCode]
	if exists {
		return &iso
	}
	return nil
}
