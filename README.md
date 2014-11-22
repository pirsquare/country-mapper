# Country Info Mapper in Go

## Installation

    go get github.com/pirsquare/country-mapper


## Examples
<pre>
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
	data := countryClient.MapByName("South Korea")
	fmt.Println(data.Name)           // Will Print: South Korea
	fmt.Println(data.Alpha2)         // Will Print: KR
	fmt.Println(data.Alpha3)         // Will Print: KOR
	fmt.Println(data.Currency[0])    // Will Print: KRW
	fmt.Println(data.CallingCode[0]) // Will Print: 82
	fmt.Println(data.Region)         // Will Print: Asia
	fmt.Println(data.Subregion)      // Will Print: Eastern Asia

	// you can try different variations of name, but be careful though, 
	// only commonly used names is supported
	data = countryClient.MapByName("south korea")
	fmt.Println(data.Name) // Will Print: South Korea

	data = countryClient.MapByName("대한민국")
	fmt.Println(data.Name) // Will Print: South Korea

	data = countryClient.MapByName("southkorea")
	fmt.Println(data == nil) // Will Print: true

	// You can also use different variations of mappings:

	// Map by Country Code Alpha-2
	data = countryClient.MapByAlpha2("SG")
	fmt.Println(data.Name) // Will Print: Singapore

	// Map by Country Code Alpha-3
	data = countryClient.MapByAlpha3("SGP")
	fmt.Println(data.Name) // Will Print: Singapore

	// Get all countries that uses Aussie dollar as currency
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
	dataList = countryClient.MapByCallingCode("61")
	for _, row := range dataList {
		fmt.Println(row.Name)
		// Will Print:
		// Australia
		// Christmas Island
		// Cocos (Keeling) Islands
	}

	// Get all countries from Oceania Region
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

</pre>


## How To?
### How can I use my own csv file for country's data
You can pass in an optional url to `country_mapper.Load("http"//example.com/file.csv")`. This is useful if you prefer to host the data file yourself or if you have modified some of the fields for your specific use case. Do keep in mind that the schema and order of columns should still be kept the same.


## Credits
- Thanks to @mledoze for country's data ([https://github.com/mledoze/countries](https://github.com/mledoze/countries))