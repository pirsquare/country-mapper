package main

import (
	"fmt"
	"github.com/pirsquare/country-mapper"
)

var countryClient *country_mapper.CountryInfoClient

func init() {
	client, err := country_mapper.Load()
	if err != nil {
		panic(err)
	}

	countryClient = client
}

func main() {

	fmt.Println("========================================")
	fmt.Println("MapByName - South Korea")
	fmt.Println("========================================")
	data := countryClient.MapByName("South Korea")
	fmt.Println(data.Name)           // Will Print: South Korea
	fmt.Println(data.Alpha2)         // Will Print: KR
	fmt.Println(data.Alpha3)         // Will Print: KOR
	fmt.Println(data.Currency[0])    // Will Print: KRW
	fmt.Println(data.CallingCode[0]) // Will Print: 82
	fmt.Println(data.Region)         // Will Print: Asia
	fmt.Println(data.Subregion)      // Will Print: Eastern Asia

	// you can try different variations of name, but be careful though, only commonly used names is supported
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("MapByName - south korea")
	fmt.Println("========================================")
	data = countryClient.MapByName("south korea")
	fmt.Println(data.Name) // Will Print: South Korea

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("MapByName - 대한민국")
	fmt.Println("========================================")
	data = countryClient.MapByName("대한민국")
	fmt.Println(data.Name) // Will Print: South Korea

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("MapByName - southkorea")
	fmt.Println("========================================")
	data = countryClient.MapByName("southkorea")
	fmt.Println(data == nil) // Will Print: true

	// You can also use different variations of mappings:

	// Map by Country Code Alpha-2
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("MapByAlpha2 - SG")
	fmt.Println("========================================")
	data = countryClient.MapByAlpha2("SG")
	fmt.Println(data.Name) // Will Print: Singapore

	// Map by Country Code Alpha-3
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("MapByAlpha3 - SGP")
	fmt.Println("========================================")
	data = countryClient.MapByAlpha3("SGP")
	fmt.Println(data.Name) // Will Print: Singapore

	// Get all countries that uses Aussie dollar as currency
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("MapByCurrency - AUD")
	fmt.Println("========================================")
	dataList := countryClient.MapByCurrency("AUD")
	for _, row := range dataList {
		fmt.Println(row.Name)
		// Will Print:
		// Australia
		// Christmas Island
		// Cocos (Keeling) Islands
		// Heard Island and McDonald Islands
		// Kiribati
		// Nauru
		// Norfolk Island
		// Tuvalu
	}

	// Get all countries that uses "61" as calling code
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("MapByCallingCode - 61")
	fmt.Println("========================================")
	dataList = countryClient.MapByCallingCode("61")
	for _, row := range dataList {
		fmt.Println(row.Name)
		// Will Print:
		// Australia
		// Christmas Island
		// Cocos (Keeling) Islands
	}

	// Get all countries from Oceania Region
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("MapByRegion - Oceania")
	fmt.Println("========================================")
	dataList = countryClient.MapByRegion("Oceania")
	for _, row := range dataList {
		fmt.Println(row.Name)
		// Will Print:
		// American Samoa
		// Australia
		// Christmas Island
		// Cocos (Keeling) Islands
		// Cook Islands
		// Fiji
		// French Polynesia
		// Guam
		// Kiribati
		// Marshall Islands
		// Micronesia
		// Nauru
		// New Caledonia
		// New Zealand
		// Niue
		// Norfolk Island
		// Northern Mariana Islands
		// Palau
		// Papua New Guinea
		// Pitcairn Islands
		// Samoa
		// Solomon Islands
		// Tokelau
		// Tonga
		// Tuvalu
		// Vanuatu
		// Wallis and Futuna
	}

	// Get all countries from South Eastern Asia
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("MapBySubregion - South-Eastern Asia")
	fmt.Println("========================================")
	dataList = countryClient.MapBySubregion("South-Eastern Asia")
	for _, row := range dataList {
		fmt.Println(row.Name)
		// Will Print:
		// Brunei
		// Cambodia
		// Indonesia
		// Laos
		// Malaysia
		// Myanmar
		// Philippines
		// Singapore
		// Thailand
		// Timor-Leste
		// Vietnam
	}
}
