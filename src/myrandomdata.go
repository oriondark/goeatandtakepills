package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

//GetRandomMedication pulls from the formatted pieces of the NDC
func GetRandomMedication() NDCItem {
	return randomFrom(getItems())
}

//GetRandomPharmacy pulls random from a file
func GetRandomPharmacy() PharmaItem {
	return randomFromPharm(getrandompharmaitems())

}
func randomFromPharm(source []PharmaItem) PharmaItem {
	return source[seedAndReturnRandom(len(source))]
}
func getrandompharmaitems() []PharmaItem {
	raw, err := ioutil.ReadFile("pharam.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []PharmaItem
	json.Unmarshal(raw, &c)
	return c
}

//PharmaItem is a holder for randomness
type PharmaItem struct {
	PharmacyID string  `json:"pharmacyid"`
	Name       string  `json:"name"`
	Street     string  `json:"street"`
	City       string  `json:"city"`
	State      string  `json:"state"`
	PostCode   int     `json:"postcode"`
	Image      string  `json:"image"`
	Phone      string  `json:"phone"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

func seedAndReturnRandom(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

// Returns a random part of a slice
func randomFrom(source []NDCItem) NDCItem {
	return source[seedAndReturnRandom(len(source))]
}

func getItems() []NDCItem {
	raw, err := ioutil.ReadFile("ndcdata.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []NDCItem
	json.Unmarshal(raw, &c)
	return c
}

//NDCItem is useless
type NDCItem struct {
	ProductID      string `json:"productid"`
	NDC            string `json:"ndc"`
	Proprietary    string `json:"proprietary"`
	NonProprietary string `json:"nonproprietary"`
	RouteName      string `json:"routename"`
	Classes        string `json:"classes"`
	DosageFormName string `json:"dosageformname"`
	IngredientUnit string `json:"ingredientunit"`
	Strength       string `json:"strength"`
	Substances     string `json:"substances"`
}

func contains(NDCItems []NDCItem, e string) bool {
	for _, a := range NDCItems {
		if a.Proprietary == e {
			return true
		}
	}
	return false
}

//CreateRandom uses the NDC codes and pushes them into JSon
func CreateRandom() {
	// read data from CSV file

	csvFile, err := os.Open("drugdata.csv")

	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	reader.Comma = '\t' // Use tab-delimited instead of comma <---- here!

	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var oneRecord NDCItem
	var allRecords []NDCItem

	for _, each := range csvData {
		oneRecord.ProductID = each[0]
		oneRecord.NDC = each[1]
		oneRecord.Proprietary = each[3]
		oneRecord.NonProprietary = each[5]
		oneRecord.DosageFormName = each[6]
		oneRecord.RouteName = each[7]
		oneRecord.IngredientUnit = each[15]
		oneRecord.Strength = each[14]
		oneRecord.Substances = each[13]
		oneRecord.Classes = each[16]

		if oneRecord.RouteName == "ORAL" && oneRecord.DosageFormName != "" {
			if !contains(allRecords, oneRecord.Proprietary) {
				allRecords = append(allRecords, oneRecord)
			}
		}

	}

	jsondata, err := json.Marshal(allRecords) // convert to JSON

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// sanity check
	// NOTE : You can stream the JSON data to http service as well instead of saving to file
	fmt.Println(string(jsondata))

	// now write to JSON file

	jsonFile, err := os.Create("ndcdata.json")

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsondata)
	jsonFile.Close()
}
