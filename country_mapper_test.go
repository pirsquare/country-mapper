package country_mapper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockClient *CountryInfoClient

//===========================================
// Setup Tests
//===========================================
func Test_Init(t *testing.T) {
	client, err := LoadByUrl()
	assert.Nil(t, err)
	mockClient = client
}

func TestLoadByFile(t *testing.T) {
	client, err := LoadByFile("./files/country_info.csv")
	assert.Nil(t, err)
	mockClient = client
}

//===========================================
// CountryInfoClient MapByName
//===========================================
func Test_Client_MapByName(t *testing.T) {
	// should map by name
	ret := mockClient.MapByName("South Korea")
	assert.Equal(t, ret.Name, "South Korea")
	assert.Equal(t, ret.Alpha2, "KR")
	assert.Equal(t, ret.Alpha3, "KOR")
	assert.Equal(t, ret.Capital, "Seoul")
	assert.Equal(t, ret.Currency, []string{"KRW"})
	assert.Equal(t, ret.CallingCode, []string{"82"})
	assert.Equal(t, ret.Region, "Asia")
	assert.Equal(t, ret.Subregion, "Eastern Asia")

	// should be able to map different variations of name
	ret = mockClient.MapByName("south korea")
	assert.Equal(t, ret.Name, "South Korea")

	ret = mockClient.MapByName("대한민국")
	assert.Equal(t, ret.Name, "South Korea")

	// should return nil when you try to map names not commonly used
	ret = mockClient.MapByName("southkorea")
	assert.Nil(t, ret)
}

//===========================================
// CountryInfoClient MapByAlpha2
//===========================================
func Test_Client_MapByAlpha2(t *testing.T) {
	ret := mockClient.MapByAlpha2("SG")
	assert.Equal(t, ret.Name, "Singapore")
}

//===========================================
// CountryInfoClient MapByAlpha3
//===========================================
func Test_Client_MapByAlpha3(t *testing.T) {
	ret := mockClient.MapByAlpha3("SGP")
	assert.Equal(t, ret.Name, "Singapore")
}

//===========================================
// CountryInfoClient MapByCurrency
//===========================================
func Test_Client_MapByCurrency(t *testing.T) {
	ret := mockClient.MapByCurrency("SGD")
	assert.Equal(t, ret[0].Name, "Singapore")
}

//===========================================
// CountryInfoClient MapByCallingCode
//===========================================
func Test_Client_MapByCallingCode(t *testing.T) {
	ret := mockClient.MapByCallingCode("65")
	assert.Equal(t, ret[0].Name, "Singapore")
}

//===========================================
// CountryInfoClient MapByRegion
//===========================================
func Test_Client_MapByRegion(t *testing.T) {
	countriesInOceania := []string{
		"American Samoa",
		"Australia",
		"Christmas Island",
		"Cocos (Keeling) Islands",
		"Cook Islands",
		"Fiji",
		"French Polynesia",
		"Guam",
		"Kiribati",
		"Marshall Islands",
		"Micronesia",
		"Nauru",
		"New Caledonia",
		"New Zealand",
		"Niue",
		"Norfolk Island",
		"Northern Mariana Islands",
		"Palau",
		"Papua New Guinea",
		"Pitcairn Islands",
		"Samoa",
		"Solomon Islands",
		"Tokelau",
		"Tonga",
		"Tuvalu",
		"Vanuatu",
		"Wallis and Futuna",
	}
	ret := mockClient.MapByRegion("Oceania")
	for _, row := range ret {
		assert.Contains(t, countriesInOceania, row.Name)
	}
}

//===========================================
// CountryInfoClient MapBySubregion
//===========================================
func Test_Client_MapBySubregion(t *testing.T) {
	countriesInSEA := []string{
		"Brunei",
		"Cambodia",
		"Indonesia",
		"Laos",
		"Malaysia",
		"Myanmar",
		"Philippines",
		"Singapore",
		"Thailand",
		"Timor-Leste",
		"Vietnam",
	}
	ret := mockClient.MapBySubregion("South-Eastern Asia")
	for _, row := range ret {
		assert.Contains(t, countriesInSEA, row.Name)
	}
}
