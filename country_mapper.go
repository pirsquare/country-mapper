package country_mapper

import (
	"bytes"
	"encoding/csv"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultFile = "https://raw.githubusercontent.com/pirsquare/country-mapper/master/files/country_info.csv"
)

type CountryInfoClient struct {
	Data []*CountryInfo
}

func (c *CountryInfoClient) MapByName(name string) *CountryInfo {
	for _, row := range c.Data {
		// check Name field
		if strings.ToLower(row.Name) == strings.ToLower(name) {
			return row
		}

		// check AlternateNames field
		if stringInSlice(strings.ToLower(name), row.AlternateNamesLower()) {
			return row
		}
	}
	return nil
}

func (c *CountryInfoClient) MapByAlpha2(alpha2 string) *CountryInfo {
	for _, row := range c.Data {
		if strings.ToLower(row.Alpha2) == strings.ToLower(alpha2) {
			return row
		}
	}
	return nil
}

func (c *CountryInfoClient) MapByAlpha3(alpha3 string) *CountryInfo {
	for _, row := range c.Data {
		if strings.ToLower(row.Alpha3) == strings.ToLower(alpha3) {
			return row
		}
	}
	return nil
}

func (c *CountryInfoClient) MapByCurrency(currency string) []*CountryInfo {
	rowList := []*CountryInfo{}
	for _, row := range c.Data {
		if stringInSlice(strings.ToLower(currency), row.CurrencyLower()) {
			rowList = append(rowList, row)
		}
	}
	return rowList
}

func (c *CountryInfoClient) MapByCallingCode(callingCode string) []*CountryInfo {
	rowList := []*CountryInfo{}
	for _, row := range c.Data {
		if stringInSlice(strings.ToLower(callingCode), row.CallingCodeLower()) {
			rowList = append(rowList, row)
		}
	}
	return rowList
}

func (c *CountryInfoClient) MapByRegion(region string) []*CountryInfo {
	rowList := []*CountryInfo{}
	for _, row := range c.Data {
		if strings.ToLower(row.Region) == strings.ToLower(region) {
			rowList = append(rowList, row)
		}
	}
	return rowList
}

func (c *CountryInfoClient) MapBySubregion(subregion string) []*CountryInfo {
	rowList := []*CountryInfo{}
	for _, row := range c.Data {
		if strings.ToLower(row.Subregion) == strings.ToLower(subregion) {
			rowList = append(rowList, row)
		}
	}
	return rowList
}

type CountryInfo struct {
	Name           string
	AlternateNames []string
	Alpha2         string
	Alpha3         string
	Capital        string
	Currency       []string
	CallingCode    []string
	Region         string
	Subregion      string
}

func (c *CountryInfo) AlternateNamesLower() []string {
	updated := []string{}
	for _, alternateName := range c.AlternateNames {
		updated = append(updated, strings.ToLower(alternateName))
	}
	return updated
}

func (c *CountryInfo) CurrencyLower() []string {
	updated := []string{}
	for _, currency := range c.Currency {
		updated = append(updated, strings.ToLower(currency))
	}
	return updated
}

func (c *CountryInfo) CallingCodeLower() []string {
	updated := []string{}
	for _, callingCode := range c.CallingCode {
		updated = append(updated, strings.ToLower(callingCode))
	}
	return updated
}

func readCSV(fileURL string) ([][]string, error) {
	var body io.Reader
	if isUrl(fileURL) {
		resp, err := http.Get(fileURL)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		body = resp.Body
	} else {
		data, err := ioutil.ReadFile(fileURL)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(data)
	}

	reader := csv.NewReader(body)
	reader.Comma = ';'
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Pass in an optional url or file path if you would like to use your own csv file for country's data.
// This is useful if you prefer to host the data file yourself or if you have modified some of the fields
// for your specific use case.
func Load(specifiedPath ...string) (*CountryInfoClient, error) {
	var filePath string

	// use user specified path for csv file if provided, else use default file URL
	if len(specifiedPath) > 0 {
		filePath = specifiedPath[0]
	} else {
		filePath = defaultFile
	}

	data, err := readCSV(filePath)
	if err != nil {
		return nil, err
	}

	recordList := []*CountryInfo{}
	for idx, row := range data {
		// skip header
		if idx == 0 {
			continue
		}

		// get name
		name := strings.Split(row[0], ",")[:1][0]

		// use commonly used & altSpellings names as AlternateNames
		alternateNames := strings.Split(row[0], ",")[1:]
		alternateNames = append(alternateNames, strings.Split(row[8], ",")...)

		record := &CountryInfo{
			Name:           name,
			AlternateNames: alternateNames,
			Alpha2:         row[2],
			Alpha3:         row[4],
			Capital:        row[7],
			Currency:       strings.Split(row[5], ","),
			CallingCode:    strings.Split(row[6], ","),
			Region:         row[10],
			Subregion:      row[11],
		}

		recordList = append(recordList, record)
	}

	return &CountryInfoClient{Data: recordList}, nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func isUrl(str string) bool {
	_, err := url.ParseRequestURI(str)
	return err == nil
}
